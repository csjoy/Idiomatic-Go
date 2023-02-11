package solver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

type MathSolverStub struct{}

func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
	switch expr {
	case "2+2*10":
		return 22, nil
	case "(2+2)*10":
		return 40, nil
	case "(2+2*10":
		return 0, errors.New("invalid expression: (2+2*10)")
	}
	return 0, nil
}

func TestProcessorProcessExpression(t *testing.T) {
	p := Processor{MathSolverStub{}}
	in := strings.NewReader(`2+2*10`)
	in1 := strings.NewReader(`(2+2)*10`)
	in2 := strings.NewReader(`(2+2*10`)
	sf := []io.Reader{in, in1, in2}
	var bb []byte
	in.Read(bb)
	fmt.Println(string(bb))
	data := []float64{22, 40, 0}
	hasErr := []bool{false, false, true}
	for i, d := range data {
		result, err := p.ProcessorExpression(context.Background(), sf[i])
		if err != nil && !hasErr[i] {
			t.Error(err)
		}
		if result != d {
			t.Errorf("Expected result %f, got %f", d, result)
		}
	}
}
