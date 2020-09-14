package fs

import (
	"github.com/ghjan/formula/opt"
	"github.com/ghjan/formula/utils"
	"reflect"
	"strings"
)

type RoundUpFunction struct {
}

func (*RoundUpFunction) Name() string {
	return "roundup"
}

func (f *RoundUpFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil && strings.Contains(err.Error(), "required only two arguments") {
		err = opt.MatchOneArgument(f.Name(), args...)
	} else if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}
	v1 := 0
	if len(args) > 1 {
		arg1, err := (*args[1]).Evaluate(context)
		if err != nil {
			return nil, err
		}
		v1, err = arg1.Int()
		if err != nil {
			return nil, err
		}
	}

	v0, err := arg0.Float64()
	if err != nil {
		return nil, err
	}

	v := utils.RoundUp(v0, v1)
	return opt.NewArgumentWithType(v, reflect.Float64), nil
}

func NewRoundUpFunction() *RoundUpFunction {
	return &RoundUpFunction{}
}
