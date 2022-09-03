package generics

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

type car interface{}
type truck interface{}

type vehicleUpgrader[C car, T truck] interface {
	Upgrade(C) T
}

type vehivleUpgraderSimple interface {
	Upgrade(car) truck
}

type LinkedList[T any] struct {
	val      T
	nextNode *LinkedList[T]
	prevNode *LinkedList[T]
}

func fun() {
	// maps.Keys[]()
	// slices.Sort[]()
	k := Map([]int{1, 2, 3, 4}, strconv.Itoa)
	fmt.Printf("k: %v\n", k)
}

func smallest[T constraints.Ordered](list LinkedList[T]) T {
	var t T = list.val
	iter := &list
	for iter.nextNode != nil {
		iter = iter.nextNode
		if iter.val < t {
			t = iter.val
		}
	}
	return t
}

func Map[F, T any](s []F, f func(F) T) []T {
	r := make([]T, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// type Comparator[T comparable] func(t1, t2 T) int

// func sort[T comparable](cc Comparator[T], ll LinkedList[T]) {
// 	for {
// 		cc(ll.val, ll.nextNode.val)
// 	}
// }
