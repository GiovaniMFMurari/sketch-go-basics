package main

import (
	"fmt"
	"reflect"
)

type Pet struct {
	Name     string // follow the same naming convetions
	owner    string
	age      int
	vaccines []string
}

type TamableBeast struct {
	Pet   // composition by embedding the Pet struct type in TamableBeast
	tamed bool
}

type User struct {
	Name  string `required max: 100`
	Email string `required`
}

func test() int {
	fmt.Println("Start")
	defer fmt.Println("Deferred")
	fmt.Println("End")
	recover()
	panic("oh! Shit")
	return 25
}

type myStruct struct {
	foo int
}

func pointers() {
	a := 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	// var ms *myStruct = &myStruct{foo: 42}
	ms := &myStruct{foo: 42}
	fmt.Println(ms)
	fmt.Println(ms.foo)
	ms.foo = 7
	fmt.Println(ms.foo)

}

func sum(param1, param2 int) int {
	fmt.Println(param1, param2)
	return (param1 + param2)
}

func sumMany(values ...int) int {
	var result int
	for _, v := range values {
		result += v
	}

	return result
}

func substractMany(values ...int) (result int) {
	for _, v := range values {
		result -= v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

//method
func (p Pet) sound() {
	fmt.Println("Quiiiin!")
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func interfaces() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

func misc() {
	// var (
	// 	name string = "Giovani"
	// 	age  int64  = 42
	// )

	// const myConst string = "I AM A CONST"
	// const myConst2 = "I AM ANOTHER CONST"
	// const myConstInt = 10

	// fmt.Println("Hello Go!")
	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(float32(age))
	// fmt.Println(string(age))
	//fmt.Println(strconv.Itoa(age))
	// fmt.Println(myConst)
	// fmt.Println(myConst2)
	// fmt.Println(age + myConstInt) // same as: fmt.Println(age + 42)

	const i = iota
	const (
		o = iota
		t = iota
		a = iota
	)
	const p = iota

	fmt.Println(i)
	fmt.Println(o)
	fmt.Println(t)
	fmt.Println(a)
	fmt.Println(p)

	const (
		q = iota
		w
		e
	)

	// fmt.Println(q)
	// fmt.Println(w)
	// fmt.Println(e)
	// iota is scoped to constant declaration block

	// iota use case
	// const (
	// 	catSpecialist = iota
	// 	dogSpecialist
	// 	ratSpecialist
	// )
	// specialistType := 0
	// fmt.Println(catSpecialist == specialistType)

	// const (
	// 	// errorSpecialist2 = iota
	// 	_ = iota // write only value, throw away imediatelly
	// 	catSpecialist2
	// 	dogSpecialist2
	// 	ratSpecialist2
	// )
	// specialistType2 := 0
	// fmt.Println(catSpecialist2 == specialistType2)

	// arrays

	arr := [3]int{1, 2, 3}
	arr2 := [...]int{3, 2, 1}
	var arr3 [3]int

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)

	arr3[1] = 4
	fmt.Println(arr3)
	fmt.Println(len(arr3))

	arrC := arr3
	arrC[0] = 1 // create a copy with a new pointer
	fmt.Println(arr3)
	fmt.Println(arrC)

	// slices
	fmt.Println("slices")
	slc := []int{1, 2, 3}
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))

	slcC := slc
	slcC[0] = 9
	fmt.Println(slc)
	fmt.Println(slcC)

	bigArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // also works with arrays [...]
	bASlc := bigArr[:]
	bASlc2 := bigArr[3:]
	bASlc3 := bigArr[:6]
	bASlc4 := bigArr[3:6]

	fmt.Println(bASlc)
	fmt.Println(bASlc2)
	fmt.Println(bASlc3)
	fmt.Println(bASlc4)

	// maps
	fmt.Println("MAPS")
	countryCodeNames := map[string]string{
		"BR": "Brazil",
		"US": "United States",
	}

	lateMap := make(map[int]string)
	lateMap = map[int]string{
		1: "banana",
		2: "potato",
	}

	fmt.Println(countryCodeNames)
	fmt.Println(lateMap)

	lateMap[3] = "apple"
	fmt.Println(lateMap)
	fmt.Println(lateMap[4])
	pop, ok := countryCodeNames["AR"]
	fmt.Println(pop, ok)

	kitty := Pet{
		Name:     "Snowball",
		owner:    "Gio",
		age:      3,
		vaccines: []string{"fiv", "felvp"},
	}

	fmt.Println(kitty)
	fmt.Println(kitty.owner, kitty.Name)

	doggo := Pet{ // positional syntax, pls don't
		"Rex",
		"Natty",
		5,
		[]string{"sinomosis", "flu"},
	}

	fmt.Println(doggo)
	fmt.Println(doggo.Name)

	creature := struct {
		name string
		age  int
	}{name: "Wormo", age: 500} // anonymous struct
	fmt.Println(creature)
	fmt.Printf("%T\n", creature)
	fmt.Printf("%T\n", doggo)

	beast := TamableBeast{
		Pet: Pet{
			Name: "Boar",
			age:  11,
		},
		tamed: false,
	}

	fmt.Println(beast)

	u := reflect.TypeOf(User{})
	field, _ := u.FieldByName("Name")
	fmt.Println(u.Name)
	fmt.Println(field.Tag)
	fmt.Println(test())
}

func methods() {
	p := Pet{}
	p.sound()
}
func main() {
	//pointers()

	// methods()
	interfaces()

}
