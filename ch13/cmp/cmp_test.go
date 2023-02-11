package cmp_test

import (
	"github.com/google/go-cmp/cmp"
	main "main/cmp"
	"testing"
)

func TestCreatePerson(t *testing.T) {
	expected := main.Person{
		Name: "Dennis",
		Age:  37,
	}
	result := main.CreatePerson("Dennis", 37)
	comparer := cmp.Comparer(func(x, y main.Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}
