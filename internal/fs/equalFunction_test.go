package fs

import (
	"fmt"
	"github.com/ghjan/formula/opt"
	"github.com/ghjan/formula/pkg/exp"
	"github.com/ghjan/formula/utils"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestEqualFunction_Evaluate(t *testing.T) {
	tests := []struct {
		arg0 string
		arg1 string
		want bool
	}{
		{"2", "1", false},
		{"2", "2", true},
		{"-2", "-2", true},
		{"1.1", "0.999", false},
		{"旧房翻新", "旧房翻新", true},
		{"毛坯房", "毛坯房", true},
		{"旧房翻新", "毛坯房", false},
		{"毛坯房", "旧房翻新", false},
		{"true", "TRUE", true},
		{"FALSE", "false", true},
		{"true", "FALSE", false},
		{"true", "1", true},
		{"true", "0", false},
		{"FALSE", "0", true},
	}
	g := NewEqualFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.arg0, tt.arg1), func(t *testing.T) {
			var got *opt.Argument
			var err error
			if utils.IsNum(tt.arg0) && utils.IsNum(tt.arg1) {
				got, err = g.Evaluate(nil, []*opt.LogicalExpression{exp.NewFloatExpression(tt.arg0), exp.NewFloatExpression(tt.arg1)}...)
			} else {
				val1 := strings.TrimSpace(strings.ToLower(tt.arg0))
				val2 := strings.TrimSpace(strings.ToLower(tt.arg1))
				if (val1 == "true" || val1 == "false") || (val2 == "true" || val2 == "false") {
					val1Bool, _ := strconv.ParseBool(tt.arg0)
					val2Bool, _ := strconv.ParseBool(tt.arg1)
					got, err = g.Evaluate(nil, []*opt.LogicalExpression{exp.NewBooleanValueExpression(val1Bool), exp.NewBooleanValueExpression(val2Bool)}...)
				} else {
					got, err = g.Evaluate(nil, []*opt.LogicalExpression{exp.NewStringValueExpression(tt.arg0), exp.NewStringValueExpression(tt.arg1)}...)
				}
			}
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
