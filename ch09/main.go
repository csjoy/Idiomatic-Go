// Package money provides various utilities to make it easy to manage money
package main

import (
	"encoding/binary"
	"fmt"
	// "reflect"

	// "github.com/csjoy/package-demo/formatter"
	// "github.com/csjoy/package-demo/math"
	crand "crypto/rand"
	"math/rand"
)

func main() {
	// num := math.Double(2)
	// str := print.Format(num)
	// fmt.Println(str)

	// seedRand()

	// s := Bar{
	// 	X: 27,
	// 	S: "hello",
	// }
	// fmt.Println(reflect.TypeOf(s))
}

func MakeBar() Bar {
	bar := Bar{
		X: 20,
		S: "Hello",
	}
	var f Foo = bar
	fmt.Println(f.Hello())
	return bar
}

type Foo struct {
	X int
	S string
}

func (f Foo) Hello() string {
	return "hello"
}
func (f Foo) Goodbye() string {
	return "goodbye"
}

type Bar = Foo

// Decimal type declairation
type Decimal int

// Money represents the combination of an amount of money
// and the currency the money is in
type Money struct {
	Value    Decimal
	Currency string
}

// Convert converts the value of one currency to another
// It has two parameters: a Money instance with the value to convert,
// and a string that represents the currency to convert to. Convert returns
// the converted currency and any errors encountered from unknown or unconvertible
// currencies.
// If an error is retured, the Money instance is set to the sero value.
//
// Supported currencies are:
//
//	USD - US Dollar
//	CAD - Canadian Dollar
//	EUR - Euro
//	INR - Indian Rupee
//
// More information on exchange rates can be found
// at https://www.investopedia.com/terms/e/exchanerate.asp
func Convert(from Money, to string) (Money, error) {
	return from, nil
}

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed with cryptographic random number generator")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	fmt.Println(b, r)
	return r
}
