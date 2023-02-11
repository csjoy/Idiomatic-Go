package main

import (
	"errors"
	"fmt"
	"os"
	// "io"
	// "log"
	// "os"
	// "reflect"
	// "sort"
	// "strconv"
)

type person struct {
	age  int
	name string
}

// Person type
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// opFuncType type
type opFuncType func(int, int) int

// MyFuncOpts type
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// result := division(5, 2)
	// fmt.Println(result)

	// MyFunc(MyFuncOpts{
	// 	LastName: "Patel",
	// 	Age:      50,
	// })
	// MyFunc(MyFuncOpts{
	// 	FirstName: "Joe",
	// 	LastName:  "Smith",
	// })

	// vals := []int{1, 2, 3, 4, 5}
	// result := addTo(2, vals...)
	// fmt.Println(result)

	// result, remainder, err := divAndRemainder(5, 2)
	// divAndRemainder(2, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result, remainder)

	// x, y, z := divAndRemainder1(5, 2)
	// fmt.Println(x, y, z)

	// x, y, z := divAndRemainder3(5, 0)
	// fmt.Println(x, y, z)

	// var optionMap = map[string]opFuncType{
	// 	"+": add,
	// 	"-": sub,
	// 	"*": mul,
	// 	"/": div,
	// }
	// expressions := [][]string{
	// 	{"2", "+", "3"},
	// 	{"2", "-", "3"},
	// 	{"2", "*", "3"},
	// 	{"2", "/", "3"},
	// 	{"2", "%", "3"},
	// 	{"two", "+", "three"},
	// 	{"5"},
	// }
	// for _, expression := range expressions {
	// 	if len(expression) != 3 {
	// 		fmt.Println("invalid expression:", expression)
	// 		continue
	// 	}
	// 	v1, err := strconv.Atoi(expression[0])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	op := expression[1]
	// 	opFunc, ok := optionMap[op]
	// 	if !ok {
	// 		fmt.Println("unsupported operator:", op)
	// 		continue
	// 	}
	// 	v2, err := strconv.Atoi(expression[2])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	result := opFunc(v1, v2)
	// 	fmt.Println(result)
	// }

	// for i := 0; i < 5; i++ {
	// 	func(j int) {
	// 		fmt.Println("Printing", j, "from inside of an anonymous function")
	// 	}(i)
	// }

	// people := []Person{
	// 	{"Pat", "Patterson", 37},
	// 	{"Tracy", "Bobbert", 23},
	// 	{"Fred", "Fredson", 18},
	// }
	// fmt.Println(people)
	// sort.Slice(people, func(i, j int) bool {
	// 	return people[i].LastName < people[j].LastName
	// })
	// fmt.Println(people)
	// sort.Slice(people, func(i, j int) bool {
	// 	return people[i].Age < people[j].Age
	// })
	// fmt.Println(people)

	// twoBase := makeMult(2)
	// threeBase := makeMult(3)
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(twoBase(i), threeBase(i))
	// }

	// if len(os.Args) < 2 {
	// 	log.Fatal("no file specified")
	// }
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// data := make([]byte, 2048)
	// for {
	// 	count, err := f.Read(data)
	// 	os.Stdout.Write(data[:count])
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}
	// }
	// var a byte
	// fmt.Println(reflect.TypeOf(a))

	// f, closer, err := getFile(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer closer()
	// data := make([]byte, 2048)
	// for {
	// 	count, err := f.Read(data)
	// 	os.Stdout.Write(data[:count])
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}
	// }

	// p := person{}
	// i := 2
	// s := "Hello"
	// modifyFails(i, s, p)
	// fmt.Println(i, s, p)

	// m := map[int]string{
	// 	1: "first",
	// 	2: "second",
	// }
	// modMap(m)
	// fmt.Println(m)
	// s := []int{1, 2, 3}
	// modSlice(s)
	// fmt.Println(s)
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

func example() {
	defer func() int {
		return 2
	}()
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }

func divAndRemainder3(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}

func divAndRemainder2(numerator int, denominator int) (result int, remainder int, err error) {
	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func divAndRemainder1(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// MyFunc function
func MyFunc(opts MyFuncOpts) {
	fmt.Println(opts)
}

func division(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
