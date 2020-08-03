package exp

import (
	"github.com/ghjan/formula/opt"
	"math"
	"reflect"
	"strconv"
	"time"
)

type StringValueExpression struct {
	Value string
}

func (expression *StringValueExpression) Parameters() ([]string) {
	return nil
}

func NewStringValueExpression(value string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &StringValueExpression{
		Value: value,
	}

	return &result
}

func (expression *StringValueExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType(expression.Value, reflect.String), nil
}

func (*StringValueExpression) ToString() string {
	return ""
}

type IntegerValueExpression struct {
	Value int64
}

func (expression *IntegerValueExpression) Parameters() ([]string) {
	return nil
}

func NewIntegerValueExpression(value string) *opt.LogicalExpression {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}

	var result opt.LogicalExpression = &IntegerValueExpression{
		Value: i,
	}

	return &result
}

func (expression *IntegerValueExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType(expression.Value, reflect.Int64), nil
}

type FloatExpression struct {
	Value float64
}

func (expression *FloatExpression) Parameters() ([]string) {
	return nil
}

func NewFloatExpression(value string) *opt.LogicalExpression {
	v, err := strconv.ParseFloat(value, 10)
	if err != nil {
		panic(err)
	}

	var result opt.LogicalExpression = &FloatExpression{
		Value: v,
	}

	return &result
}

func (expression *FloatExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType(expression.Value, reflect.Float64), nil
}

type PiExpression struct {
}

func (expression *PiExpression) Parameters() ([]string) {
	return nil
}

func (*PiExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType(math.Pi, reflect.Float64), nil
}

func NewPiExpression() *opt.LogicalExpression {
	var result opt.LogicalExpression = &PiExpression{}

	return &result
}

//todo:时间表达式如何处理？
type DateTimeExpression struct {
	Value time.Time
}

func (expression *DateTimeExpression) Parameters() []string {
	return nil
}

func NewDateTimeExpression(value string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &DateTimeExpression{}
	return &result
}

func (expression *DateTimeExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType(expression.Value.Second(), reflect.Int64), nil
}

type BooleanValueExpression struct {
	Value bool
}

func (expression *BooleanValueExpression) Parameters() ([]string) {
	return nil
}

func NewBooleanValueExpression(value bool) *opt.LogicalExpression {
	var result opt.LogicalExpression = &BooleanValueExpression{
		Value: value,
	}

	return &result
}

func (expression *BooleanValueExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType(expression.Value, reflect.Bool), nil
}
