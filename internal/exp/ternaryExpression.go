package exp

import (
	"fmt"
	"github.com/ghjan/formula/opt"
	"reflect"
)

type TernaryExpression struct {
	Left   *opt.LogicalExpression
	Middle *opt.LogicalExpression
	Right  *opt.LogicalExpression
}

func (expression *TernaryExpression) Parameters() []string {
	p1 := (*expression.Left).Parameters()
	p2 := (*expression.Middle).Parameters()
	p3 := (*expression.Right).Parameters()
	if p1 != nil {
		if p2 != nil {
			p1 = append(p1, p2...)
		}
		if p3 != nil {
			p1 = append(p1, p3...)
		}
		return p1
	} else {
		if p2 != nil {
			if p3 != nil {
				p2 = append(p2, p3...)
			}
			return p2
		} else {
			return p3
		}
	}

}

func NewTernaryExpression(left, mid, right *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &TernaryExpression{
		Left:   left,
		Middle: mid,
		Right:  right,
	}

	return &result
}

func (expression *TernaryExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	left, err := (*expression.Left).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if left.Type != reflect.Bool {
		return nil, fmt.Errorf("ternary need bool first")
	}

	if left.Value.(bool) {
		return (*expression.Middle).Evaluate(context)
	}

	return (*expression.Right).Evaluate(context)
}
