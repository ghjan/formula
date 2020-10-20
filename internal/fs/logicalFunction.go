package fs

import (
	"fmt"
	"github.com/ghjan/formula/opt"
	"reflect"
)

type ORFunction struct {
}

func (*ORFunction) Name() string {
	return "oor"
}

func (f *ORFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("function %s required  at least 2 arguments", f.Name())
	}

	for _, a := range args {
		argTemp, err := (*a).Evaluate(context)
		if err != nil {
			return nil, err
		}

		if argTemp.Type != reflect.Bool {
			return nil, fmt.Errorf("the argument of function %s should be bool", f.Name())
		}
		if argTemp.Value.(bool) {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	}
	return opt.NewArgumentWithType(false, reflect.Bool), nil

}

func NewORFunction() *ORFunction {
	return &ORFunction{}
}

type ANDFunction struct {
}

func (*ANDFunction) Name() string {
	return "aand"
}

func (f *ANDFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("function %s required  at least 2 arguments", f.Name())
	}

	for _, a := range args {
		argTemp, err := (*a).Evaluate(context)
		if err != nil {
			return nil, err
		}

		if argTemp.Type != reflect.Bool {
			return nil, fmt.Errorf("the argument of function %s should be bool", f.Name())
		}
		if !argTemp.Value.(bool) {
			return opt.NewArgumentWithType(false, reflect.Bool), nil
		}
	}
	return opt.NewArgumentWithType(true, reflect.Bool), nil

}

func NewANDFunction() *ANDFunction {
	return &ANDFunction{}
}
