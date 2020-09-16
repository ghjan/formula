package fs

import (
	"github.com/ghjan/formula/opt"
	"math"
	"reflect"
	"strings"
)

type RoundFunction struct {
}

func (*RoundFunction) Name() string {
	return "round"
}

func (f *RoundFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil && strings.Contains(err.Error(), "required only") {
		err = opt.MatchOneArgument(f.Name(), args...)
	} else if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}
	v1 := 0.0
	if len(args) > 1 {
		arg1, err := (*args[1]).Evaluate(context)
		if err != nil {
			return nil, err
		}
		v1, err = arg1.Float64()
		if err != nil {
			return nil, err
		}
	}

	v0, err := arg0.Float64()
	if err != nil {
		return nil, err
	}

	v := v0
	if math.Abs(v1) > 0.01 && math.Abs(v1) < 10 {
		v = v0 * math.Pow(10, v1)
		return opt.NewArgumentWithType(math.Round(v)/math.Pow(10, v1), reflect.Float64), nil
	} else {
		return opt.NewArgumentWithType(math.Round(v), reflect.Float64), nil
	}
}

func NewRoundFunction() *RoundFunction {
	return &RoundFunction{}
}
