package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	// "os"
)

func main() {
	// numerator := 20
	// denominator := 0
	// remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(remainder, mod)

	// num, err := doubleEven1(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(num)

	// num, err := doubleEven2(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(num)

	// data := []byte("This is not a zip file")
	// notAZipFile := bytes.NewReader(data)
	// _, err := zip.NewReader(notAZipFile, int64(len(data)))
	// if err == zip.ErrFormat {
	// 	fmt.Println("Told you so")
	// }

	// err := GenerateError(true)
	// fmt.Println(err != nil)
	// err = GenerateError(false)
	// fmt.Println(err != nil)

	// err := fileChecker("not_here.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
	// 		fmt.Println(wrappedErr)
	// 	}
	// }

	// err := fileChecker("not_here.txt")
	// if err != nil {
	// 	if errors.Is(err, os.ErrNotExist) {
	// 		fmt.Println("That file doesn't exist")
	// 	}
	// }

	// doPanic(os.Args[0])

	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}

func div60(i int) {
	// defer func() {
	// 	if v := recover(); v != nil {
	// 		fmt.Println(v)
	// 	}
	// }()
	fmt.Println(60 / i)
}

func doPanic(msg string) {
	panic(msg)
}

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}
func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		return matchResource && matchCode || matchResource && ignoreCode || ignoreResource && matchCode
	}
	return false
}

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, me2)
	}
	return false
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker %w", err)
	}
	f.Close()
	return nil
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func GenerateError(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func doubleEven1(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func doubleEven2(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}
