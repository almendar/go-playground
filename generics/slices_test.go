package generics

import "testing"

func Test_concat(t *testing.T) {
	input := []IntCover{{2}, {3}}
	concatSingle(input[0])
	funnyChan := make(chan stringer)
	concatSingle123(funnyChan)
	funnyChan <- input[0]
	concat(input)
}
