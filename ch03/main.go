package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var x = [3]int{1, 2, 3}
	// var y = [12]int{1, 5: 4, 6, 10: 100, 15}
	// var z = [...]int{1, 2, 3}
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(x == z)
	// var multi [2][3]int
	// multi[1][1] = 7
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("%d ", multi[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println(len(multi[1]))

	// var x = []int{10, 20, 30}
	// var y = []int{1, 5: 4, 6, 10: 100, 15}
	// var z []int
	// z = append(z, 10)
	// z = append(z, 50)
	// x = append(x, 40, 50)
	// x = append(x, y...)
	// x = append(x, z...)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(cap(x))
	// fmt.Println("x:", len(x), "y:", len(y), "z:", len(z))

	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 30)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 50)
	// fmt.Println(x, len(x), cap(x))

	// x := make([]int, 5, 10)
	// x = append(x, 10)
	// x[2] = 3
	// fmt.Println(x)
	// y := make([]int, 0, 10)
	// y = append(y, 5, 6, 7, 8)
	// fmt.Println(y)

	// x := []int{1, 2, 3, 4}
	// a := x[:2]
	// fmt.Println(a)
	// b := x[1:]
	// fmt.Println(b)
	// c := x[1:3]
	// fmt.Println(c)
	// d := x[:]
	// fmt.Println(d)

	// x := []int{1, 2, 3, 4}
	// y := x[:2] // 1, 2
	// z := x[1:] // 2, 3, 4
	// fmt.Println(y)
	// fmt.Println(z)
	// x[1] = 20
	// fmt.Println(x) // 1, 20, 3, 4
	// y[0] = 10
	// fmt.Println(y) // 10, 20
	// z[1] = 30
	// fmt.Println(z) // 20, 30, 4

	// x := []int{1, 2, 3, 4}
	// y := x[:2]
	// fmt.Println(cap(x), cap(y))
	// y = append(y, 30)
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)

	// x := make([]int, 0, 5)    // cap:5
	// x = append(x, 1, 2, 3, 4) // cap:5 x = 1, 2, 3, 4, X
	// y := x[:2]                // cap:5 1, 2, X, X, X
	// z := x[2:]                // cap:3 3, 4, X
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, 30, 40, 50) // cap:5 1, 2, 30, 40, 70
	// x = append(x, 60)         // cap:5 1, 2, 30, 40, 70
	// z = append(z, 70)         // cap:3 30, 40, 70
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// x := make([]int, 0, 5)    // cap:5
	// x = append(x, 1, 2, 3, 4) // cap:5 x = 1, 2, 3, 4, X
	// y := x[:2:2]              // cap:2 1, 2
	// z := x[2:4:4]             // cap:2 3, 4
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, 30, 40, 50) // cap:8 1, 2, 30, 40, 50
	// x = append(x, 60)         // cap:5 1, 2, 3, 4, 60
	// z = append(z, 70)         // cap:4 3, 4, 70
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// x := [4]int{5, 6, 7, 8} // 10, 6, 7, 8cap 4
	// y := x[:2]              // 10, 6, X, X - cap 4
	// z := x[2:]              // 7, 8 - cap 2
	// fmt.Println(cap(x), cap(y), cap(z))
	// x[0] = 10
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// x := []int{1, 2, 3, 4}
	// y := make([]int, 4)
	// numCopy := copy(y, x)
	// fmt.Println(y, numCopy)

	// x := []int{1, 2, 3, 4}
	// y := make([]int, 2)
	// numCopy := copy(y, x)
	// fmt.Println(y, numCopy)

	// x := []int{1, 2, 3, 4}
	// y := make([]int, 2)
	// numCopy := copy(y, x[2:])
	// fmt.Println(y, numCopy)
	// add(3, 4)

	// x := []int{1, 2, 3, 4}
	// num := copy(x[:3], x[1:]) // 1,2,3 - 2,3,4 - 2,3,4,4
	// fmt.Println(x, num)

	// x := []int{1, 2, 3, 4}
	// d := [4]int{5, 6, 7, 8}
	// y := make([]int, 2)
	// copy(y, d[:])
	// fmt.Println(y)
	// copy(d[:], x)
	// fmt.Println(d)

	// var s string = "Hello there"
	// var b byte = s[6]
	// fmt.Println(s, b)

	// //							01234567890
	// var s string = "Hello there"
	// var s2 string = s[4:7] // 0 t
	// var s3 string = s[:5]  // Hello
	// var s4 string = s[6:]  // there
	// fmt.Println(s2)
	// fmt.Println(s3)
	// fmt.Println(s4)

	// //							01234567890
	// var s string = "Hello 🌞"
	// var s2 string = s[4:7]
	// var s3 string = s[:5]
	// var s4 string = s[6:]
	// fmt.Println(s2)
	// fmt.Println(s3)
	// fmt.Println(s4)
	// fmt.Println(len(s))

	// var a rune = 'x'
	// var s string = string(a)
	// var b byte = 'y'
	// var s2 string = string(b)
	// fmt.Println(s, reflect.TypeOf(s2))

	// var x int = 654
	// var y = string(x)
	// fmt.Println(y)

	// var s string = "Hello, 🌞"
	// var bs []byte = []byte(s)
	// var rs []rune = []rune(s)
	// var x []int32
	// fmt.Println(bs)
	// fmt.Println(rs)
	// fmt.Println(reflect.TypeOf(rs))
	// fmt.Println(reflect.TypeOf(x))

	// var nilMap map[string]int     // nil map
	// totalWins := map[string]int{} // empty map
	// fmt.Println(len(nilMap), len(totalWins))

	// teams := map[string][]string{
	// 	"Orcas":   {"Fred", "Ralph", "Bijou"},
	// 	"Lions":   {"Sarah", "Peter", "Billie"},
	// 	"Kittens": {"Waldo", "Raul", "Ze"},
	// }
	// fmt.Println(teams["Orcas"])

	// ages := make(map[int][]string, 10)
	// fmt.Println(ages)

	// totalWins := map[string]int{}
	// totalWins["Orcas"] = 1
	// totalWins["Lions"] = 2
	// fmt.Println(totalWins["Orcas"])
	// fmt.Println(totalWins["Kittens"])
	// totalWins["Kittens"]++
	// fmt.Println(totalWins["Kittens"])
	// totalWins["Lions"] = 3
	// fmt.Println(totalWins["Lions"])

	// m := map[string]int{
	// 	"hello": 5,
	// 	"world": 0,
	// }
	// v, ok := m["hello"]
	// fmt.Println(v, ok)
	// v, ok = m["world"]
	// fmt.Println(v, ok)
	// v, ok = m["goodbye"]
	// fmt.Println(v, ok)
	// fmt.Println(m["world"], m["goodbye"])

	// m := map[string]int{
	// 	"hello": 5,
	// 	"world": 10,
	// }
	// fmt.Println(m)
	// delete(m, "hello")
	// fmt.Println(m)

	// intSet := map[int]bool{}
	// vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	// for _, v := range vals {
	// 	intSet[v] = true
	// }
	// fmt.Println(len(vals), len(intSet))
	// fmt.Println(intSet[5])
	// fmt.Println(intSet[500])
	// if intSet[100] {
	// 	fmt.Println("100 is in the set")
	// }

	// intSet := map[int]bool{}
	// vals := []int{5, 1, 2, 3, 4, 5, 6, 6, 7, 8}
	// for _, v := range vals {
	// 	intSet[v] = true
	// }
	// fmt.Println(len(vals), len(intSet))

	// intSet := map[int]struct{}{}
	// vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	// for _, v := range vals {
	// 	intSet[v] = struct{}{}
	// }
	// if _, ok := intSet[4]; ok {
	// 	fmt.Println("5 is in the set")
	// }

	// type person struct {
	// 	name string
	// 	age  int
	// 	pet  string
	// }
	// var fred person
	// bob := person{}
	// julia := person{
	// 	"Julia",
	// 	40,
	// 	"Cat",
	// }
	// beth := person{
	// 	pet:  "Dog",
	// 	name: "Beth",
	// }
	// fmt.Println(fred, bob, julia, beth, julia.age)

	// var person struct {
	// 	name string
	// 	age  int
	// 	pet  string
	// }
	// person.name = "Bob"
	// person.age = 50
	// person.pet = "dog"
	// pet := struct {
	// 	name string
	// 	kind string
	// }{
	// 	name: "Fido",
	// 	kind: "Dog",
	// }
	// fmt.Println(pet, person)

	// type firstPerson struct {
	// 	name string
	// 	age  int
	// }
	// type secondPerson struct {
	// 	name string
	// 	age  int
	// }
	// type thirdPerson struct {
	// 	age  int
	// 	name string
	// }
	// type fourthPerson struct {
	// 	firstName string
	// 	age       int
	// }
	// type fifthPerson struct {
	// 	name          string
	// 	age           int
	// 	favoriteColor string
	// }

	// var fp firstPerson
	// sp := secondPerson{
	// 	name: "joy",
	// 	age:  26,
	// }
	// fmt.Println(reflect.TypeOf(sp))
	// fp = firstPerson(sp)
	// fmt.Println(reflect.TypeOf(fp))
	// fmt.Println(fp)
	// // tp := thirdPerson{
	// // 	age:  26,
	// // 	name: "adil",
	// // }
	// // fp = firstPerson(tp)

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	} = f
	fmt.Println(f == g)
	fmt.Println(reflect.TypeOf(g))
}
