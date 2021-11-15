package phonebook

import "strings"

type Person struct {
	Name string
	Age  int
}

type AdressBook struct {
	mappings map[string]Person
}

func (ab *AdressBook) findPrefix(prefix string) []Person {
	acc := make([]Person, 0, 10)
	for k, v := range ab.mappings {
		if strings.HasPrefix(k, prefix) {
			acc = append(acc, v)
		}
	}
	return acc
}
