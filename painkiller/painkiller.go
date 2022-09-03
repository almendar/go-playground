package painkiller

//go:generate jsonenums -type=Pill

type Pill int

const (
	Unknown      = 0
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen
)
