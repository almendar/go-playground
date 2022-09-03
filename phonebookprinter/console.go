package phonebookprinter

import (
	"fmt"

	"github.com/almendar/go-playground/phonebook"
)

type PersonShow struct {
	phonebook.Person
}

type Shower[A fmt.Stringer] interface {
	Show()
}
