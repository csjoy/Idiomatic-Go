package circular

type Pet struct {
	Name      string
	Type      string
	OwnerName string
}

var owners = map[string]Person{
	"Bob":   {"Bob", 30, "Fluffy"},
	"Julia": {"Julia", 40, "Rex"},
}

func (p Pet) Owner() Person {
	return owners[p.OwnerName]
}
