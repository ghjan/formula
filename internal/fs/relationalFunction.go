package fs

import (
	"github.com/ghjan/formula/opt"
	"math"
	"reflect"
)

type EqualFunction struct {
}

func (*EqualFunction) Name() string {
	return "=="
}

func (f *EqualFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if arg0.IsNumber() && arg1.IsNumber() {
		v0, _ := arg0.Float64()
		v1, _ := arg1.Float64()
		if math.Abs(v0-v1) <= 0.000001 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	} else if arg0.IsBool() || arg1.IsBool() {
		v0 := arg0.Bool()
		v1 := arg1.Bool()
		if v0 == v1 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	} else {
		v0 := arg0.String()
		v1 := arg1.String()
		if v0 == v1 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	}

	return opt.NewArgumentWithType(false, reflect.Bool), nil
}

func NewEqualFunction() *EqualFunction {
	return &EqualFunction{}
}

type NotGreaterFunction struct {
}

func (*NotGreaterFunction) Name() string {
	return "<="
}

func (f *NotGreaterFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if arg0.IsNumber() && arg1.IsNumber() {
		v0, _ := arg0.Float64()
		v1, _ := arg1.Float64()
		if v0 <= v1 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	} else if arg0.IsBool() || arg1.IsBool() {
		v0 := arg0.Bool()
		v1 := arg1.Bool()
		if !v0 || v1 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	} else {
		v0 := arg0.String()
		v1 := arg1.String()
		if v0 <= v1 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	}

	return opt.NewArgumentWithType(false, reflect.Bool), nil
}

func NewNotGreaterFunction() *NotGreaterFunction {
	return &NotGreaterFunction{}
}

type NotLessFunction struct {
}

func (*NotLessFunction) Name() string {
	return ">="
}

func (f *NotLessFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if arg0.IsNumber() && arg1.IsNumber() {
		v0, _ := arg0.Float64()
		v1, _ := arg1.Float64()
		if v0 >= v1 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	} else if arg0.IsBool() || arg1.IsBool() {
		v0 := arg0.Bool()
		v1 := arg1.Bool()
		if !v1 || v0 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	} else {
		v0 := arg0.String()
		v1 := arg1.String()
		if v0 >= v1 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	}

	return opt.NewArgumentWithType(false, reflect.Bool), nil
}

func NewNotLessFunction() *NotLessFunction {
	return &NotLessFunction{}
}


type NotEqualFunction struct {
}

func (*NotEqualFunction) Name() string {
	return "!="
}

func (f *NotEqualFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if arg0.IsNumber() && arg1.IsNumber() {
		v0, _ := arg0.Float64()
		v1, _ := arg1.Float64()
		if math.Abs(v0-v1) > 0.000001 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	} else if arg0.IsBool() || arg1.IsBool() {
		v0 := arg0.Bool()
		v1 := arg1.Bool()
		if v0 != v1 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	} else {
		v0 := arg0.String()
		v1 := arg1.String()
		if v0 != v1 {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	}

	return opt.NewArgumentWithType(false, reflect.Bool), nil
}

func NewNotEqualFunction() *NotEqualFunction {
	return &NotEqualFunction{}
}
