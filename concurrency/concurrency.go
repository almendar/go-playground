package concurrency

type Person struct {
	Name string
	Age  int
}

var People = map[string]Person{
	"Jimmy": {Name: "Jimmy", Age: 33},
}
