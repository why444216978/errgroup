package errgroup_test

import (
	"context"
	"errors"
	"testing"
	"time"

	myErrgroup "github.com/why444216978/errgroup"
	"golang.org/x/sync/errgroup"
)

func MyErrgroup() {
	g, _ := myErrgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		time.Sleep(time.Second * 1)
		return
	})
	g.Go(func() (err error) {
		return errors.New("error")
	})
	//模拟等待调度时间差
	time.Sleep(time.Second * 1)
	g.Go(func() (err error) {
		time.Sleep(time.Second * 5)
		return
	})
	_ = g.Wait()
}

func GolangErrgroup() {
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		time.Sleep(time.Second * 1)
		return
	})
	g.Go(func() (err error) {
		return errors.New("error")
	})
	//模拟等待调度时间差
	time.Sleep(time.Second * 1)
	g.Go(func() (err error) {
		time.Sleep(time.Second * 5)
		return
	})
	_ = g.Wait()
}

func BenchmarkMyErrgroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyErrgroup()
	}
}

func BenchmarkGolangErrgroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GolangErrgroup()
	}
}
