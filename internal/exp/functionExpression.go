package exp

import (
	"github.com/ghjan/formula/internal/cache"
	"github.com/ghjan/formula/opt"
)

type FunctionExpression struct {
	Identifier *opt.LogicalExpression
	Arguments  []*opt.LogicalExpression
}

func (expression *FunctionExpression) Parameters() []string {
	var results []string
	for _, arg := range expression.Arguments {
		if pa :=(*arg).Parameters(); pa != nil {
			results = append(results, pa...)
		}
	}
	return results
}

func NewFunctionExpression(id *opt.LogicalExpression, args []*opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &FunctionExpression{
		Identifier: id,
		Arguments:  args,
	}

	return &result
}

func (expression *FunctionExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	name, _ := (*expression.Identifier).Evaluate(context)
	f, err := cache.FindFunction(name.String())
	if err != nil {
		return nil, err
	}

	return (*f).Evaluate(context, expression.Arguments...)
}
