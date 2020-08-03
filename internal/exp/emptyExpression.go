package exp

import (
	"github.com/ghjan/formula/opt"
	"reflect"
)

type EmptyExpression struct {
}

func (expression *EmptyExpression) Parameters() []string {
	return nil
}

func NewEmptyExpression() *opt.LogicalExpression {
	var result opt.LogicalExpression = &EmptyExpression{}
	return &result
}

func (expression *EmptyExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType("", reflect.String), nil
}
