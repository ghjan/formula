package fs

import (
	"fmt"
	"github.com/ghjan/formula/pkg/exp"
	"reflect"
	"testing"

	"github.com/ghjan/formula/opt"
)

func TestInFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args []*opt.LogicalExpression
		want bool
	}{
		{args: []*opt.LogicalExpression{
			exp.NewFloatExpression("1.1"),
			exp.NewFloatExpression("1.1"),
			exp.NewStringValueExpression("1"),
		}, want: true},
		{args: []*opt.LogicalExpression{
			exp.NewFloatExpression("1.1"),
			exp.NewStringValueExpression("1.1"),
			exp.NewStringValueExpression("1"),
		}, want: false},
		{args: []*opt.LogicalExpression{
			exp.NewStringValueExpression("1.1"),
			exp.NewFloatExpression("1.1"),
			exp.NewStringValueExpression("1"),
		}, want: false},
		{args: []*opt.LogicalExpression{
			exp.NewStringValueExpression("楼梯房"),
			exp.NewStringValueExpression("电梯房"),
			exp.NewStringValueExpression("楼梯房"),
		}, want: true},
		{args: []*opt.LogicalExpression{
			exp.NewStringValueExpression("电梯房"),
			exp.NewStringValueExpression("电梯房"),
			exp.NewStringValueExpression("楼梯房"),
		}, want: true},
		{args: []*opt.LogicalExpression{
			exp.NewStringValueExpression(""),
			exp.NewStringValueExpression("电梯房"),
			exp.NewStringValueExpression("楼梯房"),
		}, want: false},
	}
	i := NewInFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {

			got, err := i.Evaluate(nil, tt.args...)
			if err != nil {
				t.Fatal(err)
			}

			if got.Type != reflect.Bool {
				t.Fatal("error type")
			}

			if got.Value.(bool) != tt.want {
				t.Fatal(got.Value)
			}
		})
	}
}
