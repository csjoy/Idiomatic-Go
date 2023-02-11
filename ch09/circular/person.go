package circular

type Person struct {
	Name    string
	Age     int
	PetName string
}

var pets = map[string]Pet{
	"Fluffy": {"Fluffy", "Cat", "Bob"},
	"Rex":    {"Rex", "Dog", "Julia"},
}

func (p Person) Pet() Pet {
	return pets[p.PetName]
}
