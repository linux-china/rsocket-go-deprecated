package test

import (
	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
	"testing"
)

func TestSubscribe(t *testing.T) {
	onNext := handlers.NextFunc(func(item interface{}) {
		println(item)
	})

	onDone := handlers.DoneFunc(func() {
		println("Done")
	})
	observable.Just(1).Subscribe(observer.New(onNext, onDone))
}
