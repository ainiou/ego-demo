// Code generated by struct2interface; DO NOT EDIT.

package example

import (
	"context"
)

// ExampleInterface ...
type ExampleInterface interface {
	// SayHello ...
	SayHello(ctx context.Context) (record string, err error)
}