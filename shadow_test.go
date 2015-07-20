package shadow

import (
	"errors"
	"testing"

	"github.com/tamird/shadow/other"
)

func foo() (int, error) {
	return 0, nil
}

func TestVetShadow(t *testing.T) {
	// Local varibles: this passes.
	// a, err := "a", errors.New("Foo")

	// Function from same package: this passes.
	// a, err := foo()

	// Function from standard library: this passes.
	// r := &bytes.Buffer{}
	// a, err := r.Read(nil)

	// Function from different package: this triggers shadowing warning.
	a, err := other.Foo()

	if err != nil {

	}

	if _, err := other.Foo(); err != nil {
	}

	b, err := "b", errors.New("Foo")
	if err != nil {
	}

	_, _ = a, b
}
