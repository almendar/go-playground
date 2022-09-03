package generics

import (
	"errors"
	"fmt"
)

type stringer interface {
	Stringify() string
}

type IntCover struct {
	val int
}

func (it IntCover) Stringify() string {
	return fmt.Sprintf("%d", it.val)
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.Stringify()
	}
	return result
}

func concatSingle[T stringer](val T) string {
	return val.Stringify()
}

func concatSingle1[T stringer](val chan T) string {
	return (<-val).Stringify()
}

func concatSingle123(val chan stringer) string {
	return (<-val).Stringify()
}

func indexOf[T comparable](s []T, x T) (int, error) {
	for i, v := range s {
		if v == x {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}
