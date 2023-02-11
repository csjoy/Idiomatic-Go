package main

import (
	"fmt"
	"math/rand"
	// "math/rand"
)

func main() {
	// x := 10
	// if x > 5 {
	// 	fmt.Println(x)
	// 	x := 5
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// x := 10
	// if x > 5 {
	// 	x, y := 5, 20
	// 	fmt.Println(x, y)
	// }
	// fmt.Println(x)

	// x := 10
	// fmt.Println(x)
	// fmt := "oops"
	// fmt.Println(fmt)

	// fmt.Println(true)
	// true := 10
	// fmt.Println(true)

	// n := rand.Intn(10)
	// if n == 0 {
	// 	fmt.Println("This is too low")
	// } else if n > 5 {
	// 	fmt.Println("That's too big")
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }

	// if n := rand.Intn(10); n == 0 {
	// 	fmt.Println("That's too low")
	// } else if n > 5 {
	// 	fmt.Println("That's too big")
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }
	// // fmt.Println(n)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 1
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i *= 2
	// }

	// for {
	// 	fmt.Println("Hello")
	// }

	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		if i%5 == 0 {
	// 			fmt.Println("FizzBuzz")
	// 		} else {
	// 			fmt.Println("Fizz")
	// 		}
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 		continue
	// 	}
	// 	if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 		continue
	// 	}
	// 	if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for i, v := range evenVals {
	// 	fmt.Println(i, v)
	// }

	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	fmt.Println(v)
	// }

	// uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	// for k := range uniqueNames {
	// 	fmt.Println(k)
	// }

	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Loop", i)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }
	// fmt.Println(m)

	// samples := []string{"hello", "apple_π!"}
	// for _, sample := range samples {
	// 	for i, r := range sample {
	// 		fmt.Println(i, r, string(r))
	// 	}
	// 	fmt.Println()
	// }

	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	v *= 2
	// }
	// fmt.Println(evenVals)

	// 	samples := []string{"hello", "apple_π!"}
	// outer:
	// 	for _, sample := range samples {
	// 		for i, r := range sample {
	// 			fmt.Println(i, r, string(r))
	// 			if r == 'l' {
	// 				continue outer
	// 			}
	// 		}
	// 	}

	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for i, v := range evenVals {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	if i == len(evenVals)-1 {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for i := 1; i < len(evenVals)-1; i++ {
	// 	fmt.Println(evenVals[i])
	// }

	// words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	// for _, word := range words {
	// 	switch size := len(word); size {
	// 	case 1, 2, 3, 4:
	// 		fmt.Println(word, "is a short word!")
	// 	case 5:
	// 		fmt.Println(word, "is exactly the right length:", size)
	// 	case 6, 7, 8, 9:
	// 	default:
	// 		fmt.Println(word, "is a long word!")
	// 	}
	// }

	// loop:
	//
	//	for i := 0; i < 10; i++ {
	//		switch {
	//		case i%2 == 0:
	//			fmt.Println(i, "is even")
	//		case i%3 == 0:
	//			fmt.Println(i, "is divisible by 3 but not 2")
	//		case i%7 == 0:
	//			fmt.Println("exit the loop!")
	//			break loop
	//		default:
	//			fmt.Println(i, "is boring")
	//		}
	//	}

	// words := []string{"hi", "salutations", "hello"}
	// for _, word := range words {
	// 	switch wordLen := len(word); {
	// 	case wordLen < 5:
	// 		fmt.Println(word, "is a short word!")
	// 	case wordLen > 10:
	// 		fmt.Println(word, "is a long word!")
	// 	default:
	// 		fmt.Println(word, "is exactly the right length.")
	// 	}
	// }

	// var a int = 3
	// switch {
	// case a == 2:
	// 	fmt.Println("a is 2")
	// case a == 3:
	// 	fmt.Println("a is 3")
	// case a == 4:
	// 	fmt.Println("a is 4")
	// default:
	// 	fmt.Println("a is", a)
	// }
	// switch a {
	// case 2:
	// 	fmt.Println("a is 2")
	// case 3:
	// 	fmt.Println("a is 3")
	// case 4:
	// 	fmt.Println("a is 4")
	// default:
	// 	fmt.Println("a is", a)
	// }

	// switch n := rand.Intn(10); {
	// case n == 0:
	// 	fmt.Println("That's too low")
	// case n > 5:
	// 	fmt.Println("That's too big:", n)
	// default:
	// 	fmt.Println("That's a good number:", n)
	// }

	//	a := 10
	//	goto skip
	//	b := 20
	//
	// skip:
	//
	//	c := 30
	//	fmt.Println(a, b, c)
	//	if c > a {
	//		goto inner
	//	}
	//	if a < b {
	//	inner:
	//		fmt.Println("a is less than b")
	//	}

	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
