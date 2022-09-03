package adt

import (
	"fmt"
	"testing"
)

func TestAdt(t *testing.T) {

	var acc Account = Premium{}
	pc := PaymentCalcualtor{}
	toPay := Fold[int](acc, pc)
	fmt.Printf("toPay: %v\n", toPay)

}
