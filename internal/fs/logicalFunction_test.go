package fs

import (
	"fmt"
	"github.com/ghjan/formula/opt"
	"github.com/ghjan/formula/pkg/exp"
	"testing"
)

func TestORFunction_Evaluate(t *testing.T) {
	type TestArgs struct {
		args     []bool
		expected bool
	}
	tests := []TestArgs{
		{[]bool{true, true}, true},
		{[]bool{true, false}, true},
		{[]bool{false, false}, false},
		{[]bool{false, false, false}, false},
		{[]bool{false, false, true}, true},
		{[]bool{true, false, false}, true},
		{[]bool{false, true, false}, true},
	}

	i := NewORFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			expressions := make([]*opt.LogicalExpression, 0)
			for _, arg := range tt.args {
				expTemp := exp.NewBooleanValueExpression(arg)
				expressions = append(expressions, expTemp)
			}
			result, err := i.Evaluate(nil, expressions...)
			if err != nil {
				t.Fatal(err)
			}

			if result.Bool() != tt.expected {
				t.Fail()
			}
		})
	}
}

func TestANDFunction_Evaluate(t *testing.T) {
	type TestArgs struct {
		args     []bool
		expected bool
	}
	tests := []TestArgs{
		{[]bool{true, true}, true},
		{[]bool{true, false}, false},
		{[]bool{true, true, true}, true},
		{[]bool{true, true, false}, false},
		{[]bool{false, false, false}, false},
		{[]bool{false, false, true}, false},
		{[]bool{true, false, false}, false},
		{[]bool{false, true, false}, false},
	}

	i := NewANDFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			expressions := make([]*opt.LogicalExpression, 0)
			for _, arg := range tt.args {
				expTemp := exp.NewBooleanValueExpression(arg)
				expressions = append(expressions, expTemp)
			}
			result, err := i.Evaluate(nil, expressions...)
			if err != nil {
				t.Fatal(err)
			}

			if result.Bool() != tt.expected {
				t.Fail()
			}
		})
	}
}
