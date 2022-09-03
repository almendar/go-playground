package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContex(t *testing.T) {

	rootCtx := context.Background()
	ctxWithValue := context.WithValue(rootCtx, "some_key", "847")
	ctxTimeout, cancel := context.WithTimeout(ctxWithValue, 1*time.Second)
	defer cancel()
	waitUnilDoneAndPrint(ctxTimeout)

}

func waitUnilDoneAndPrint(ctx context.Context) {
	<-ctx.Done()
	fmt.Printf("ctx.Value(\"some_key\"): %v\n", ctx.Value("some_key"))
}
