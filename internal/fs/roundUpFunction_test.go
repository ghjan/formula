package fs

import (
	"fmt"
	"github.com/ghjan/formula/opt"
	"github.com/ghjan/formula/pkg/exp"
	"strconv"
	"testing"
)

var testRoundUpArgs = []struct {
	args []string
	want string
}{
	{[]string{"123.4567", "3"}, "123.457"},
	{[]string{"123.4567", "2"}, "123.46"},
	{[]string{"123.4567", "0"}, "124"},
	{[]string{"123.4567", "-1"}, "130"},
	{[]string{"123.4567", "-2"}, "200"},
	{[]string{"1123.4567", "-3"}, "2000"},
}

func TestRoundUpFunction_Evaluate(t *testing.T) {

	f := NewRoundUpFunction()
	for _, tt := range testRoundUpArgs {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			express := make([]*opt.LogicalExpression, 0)
			express = append(express, exp.NewFloatExpression(tt.args[0]))
			express = append(express, exp.NewFloatExpression(tt.args[1]))

			result, err := f.Evaluate(nil, express...)
			if err != nil {
				t.Fatal(err)
			}

			v, err := result.Float64()
			if err != nil {
				t.Fatal(err)
			}
			want, _ := strconv.ParseFloat(tt.want, 64)
			if v != want {
				t.Fatal()
			}
		})
	}
}
