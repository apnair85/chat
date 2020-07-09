package trace

import (
	"fmt"
	"io"
)

//Tracer is the interface for an object capable of
// tracing events throughout code
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type niltracer struct{}

func (t *niltracer) Trace(a ...interface{}) {}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New is a test function
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Off creates a tracer that will ignore calls to Trace
func Off() Tracer {
	return &niltracer{}
}
