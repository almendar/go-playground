package shop

import "time"

type Owner struct {
	Address   string
	City      string
	Telephone string
	Pets      []Pet
}

type Pet struct {
	Name      string
	BirthDate time.Time
	Owner     *Owner
}
