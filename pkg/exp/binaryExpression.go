package exp

import (
	"github.com/ghjan/formula/internal/cache"
	"github.com/ghjan/formula/opt"
)

type BinaryExpression struct {
	Name            string
	LeftExpression  *opt.LogicalExpression
	RightExpression *opt.LogicalExpression
}

func (expression *BinaryExpression) Parameters() []string {
	p1 := (*expression.LeftExpression).Parameters()
	p2 := (*expression.RightExpression).Parameters()
	if p1 != nil {
		if p2!=nil{
			p1 = append(p1, p2...)
		}
		return p1
	} else {
		return p2
	}

}

func NewBinaryExpression(name string, leftExpression, rightExpression *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &BinaryExpression{
		Name:            name,
		LeftExpression:  leftExpression,
		RightExpression: rightExpression,
	}

	return &result
}

func (expression *BinaryExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	f, err := cache.FindFunction(expression.Name)
	if err != nil {
		return nil, err
	}

	return (*f).Evaluate(context, expression.LeftExpression, expression.RightExpression)
}
