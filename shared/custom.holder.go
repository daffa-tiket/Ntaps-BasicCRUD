package shared

import (
	"go.uber.org/dig"
)

type (
	CustomHolder struct {
		dig.In
	}
)

func (ch *CustomHolder) Close() {
	// - define closer here
	// - example
	/*

	   var i interface{} = nil

	   i = ch.SomeField
	   if closer, ok := i.(io.Closer); ok {
	   	_ = closer.Close()
	   }

	*/
}

func RegisterCustomHolder(container *dig.Container) error {
	// - register your custom holder here

	return nil
}
