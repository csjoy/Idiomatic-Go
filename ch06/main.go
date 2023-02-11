package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	// "reflect"
	// "unsafe"
)

func main() {
	// var x int = 10
	// var y bool = true
	// pointerX := &x
	// pointerY := &y
	// var pointerZ *string
	// fmt.Println(unsafe.Sizeof(x))
	// fmt.Println(unsafe.Sizeof(y))
	// fmt.Println(unsafe.Sizeof(pointerX))
	// fmt.Println(unsafe.Sizeof(pointerY))
	// fmt.Println(unsafe.Sizeof(pointerZ))

	// x := 10
	// pointerX := &x
	// fmt.Println(pointerX)
	// fmt.Println(*pointerX)
	// z := 5 + *pointerX
	// fmt.Println(z)

	// var x *int
	// fmt.Println(x == nil)
	// fmt.Println(reflect.TypeOf(x))
	// // fmt.Println(*x)

	// var x = new(int)
	// fmt.Println(x == nil)
	// fmt.Println(reflect.TypeOf(x))
	// fmt.Println(*x)

	// type Foo struct {
	// 	name string
	// 	age  int
	// }
	// x := &Foo{}
	// var y string
	// z := &y
	// fmt.Println(x.age)
	// fmt.Println(*z)

	// type person struct {
	// 	FirstName  string
	// 	MiddleName *string
	// 	LastName   string
	// }
	// p := person{
	// 	FirstName:  "Prosenjit",
	// 	MiddleName: stringp("Majumder"),
	// 	LastName:   "Joy",
	// }
	// fmt.Println(p)

	// var f *int
	// failedUpdate(f)
	// fmt.Println(f)

	// x := 10
	// failedUpdate1(&x)
	// fmt.Println(x)
	// update(&x)
	// fmt.Println(x)

	// f := struct {
	// 	Name string `json:"name"`
	// 	Age  int    `json:"age"`
	// }{}
	// err := json.Unmarshal([]byte(`{"name": "bob", "age": 30}`), &f)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(f)

	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make([]byte, 2048)
	for {
		count, err := file.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		} else if count == 0 {
			break
		}
		// process(data)
		os.Stdout.Write(data[:count])
	}
}

// func MakeFoo() (Foo, error) {
// 	f := Foo{
// 		Field1: "val",
// 		Field2: 20,
// 	}
// 	return f, nil
// }

func update(px *int) {
	*px = 20
}

func failedUpdate1(px *int) {
	x2 := 20
	px = &x2
}

func failedUpdate(g *int) {
	x := 10
	g = &x
	fmt.Println("g:", g)
}

func stringp(s string) *string {
	return &s
}
