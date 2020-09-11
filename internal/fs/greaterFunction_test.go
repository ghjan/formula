package fs

import (
	"fmt"
	"github.com/ghjan/formula/pkg/exp"
	"reflect"
	"testing"

	"github.com/ghjan/formula/opt"
)

func TestGreaterFunction_Evaluate(t *testing.T) {
	tests := []struct {
		arg0 string
		arg1 string
		want bool
	}{
		{"2", "1", true},
		{"-2", "1", false},
		{"1.1", "0.999", true},
	}
	g := NewGreaterFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.arg0, tt.arg1), func(t *testing.T) {
			got, err := g.Evaluate(nil, []*opt.LogicalExpression{exp.NewFloatExpression(tt.arg0), exp.NewFloatExpression(tt.arg1)}...)
			if err != nil {
				t.Fatal(err)
			}

			if got.Type != reflect.Bool {
				t.Fatal("error type")
			}

			if got.Value.(bool) != tt.want {
				t.Fatalf("%v!=%v", tt.want, !tt.want)
			}
		})
	}
}
