package concurrency

import (
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestErrGroupFun(t *testing.T) {

	var g errgroup.Group

	g.Go(func() error {

		_, err := http.Get("www.dasdasd.asfpl")
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {

		_, err := http.Get("www.dasdasd.asfpl")
		if err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		t.Error(err)
	}
}
