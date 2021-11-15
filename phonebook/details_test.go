package phonebook

import "testing"

var phoneBook = AdressBook{
	mappings: map[string]Person{
		"Tomek":     {Name: "Tomek", Age: 35},
		"Tymoteusz": {Name: "Tymoteusz", Age: 45},
	},
}

func TestPrefix(t *testing.T) {
	found := phoneBook.findPrefix("T")
	if len(found) != 2 {
		t.Logf("Found Persons: %v", found)
		t.Error("What?")
	}
}
