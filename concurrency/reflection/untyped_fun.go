package reflection

import "fmt"

type Nat interface {
	Next() Nat
	Plus(other Nat) Nat
}

type Succ struct {
	Predecessor Nat
}

type Zero struct{}

func (it Zero) Next() Nat {
	return Succ{it}
}

func (it Zero) Plus(other Nat) Nat {
	return other
}

func (it Succ) Next() Nat {
	return Succ{it}
}

func (it Succ) Plus(other Nat) Nat {
	return Succ{it.Predecessor.Plus(other)}
}

var (
	zero  = Zero{}
	one   = zero.Next()
	two   = one.Next()
	three = two.Next()
	four  = three.Next()
	five  = four.Next()
)

func FunnyStuff() {
	fmt.Printf("(two.Plus(two) == four): %v\n", (two.Plus(two) == four))
	fmt.Printf("(three.Plus(two) == four): %v\n", (three.Plus(two) == four))
	fmt.Printf("(three.Plus(two) == five): %v\n", (three.Plus(two) == five))
}
