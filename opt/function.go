package opt

import (
	"fmt"
)

type Function interface {
	Name() string
	Evaluate(context *FormulaContext, args ...*LogicalExpression) (*Argument, error)
}

func MatchArgument(name string, args ...*LogicalExpression) error {
	if args == nil || len(args) == 0 {
		return fmt.Errorf(name)
	}

	return nil
}

func MatchSeveralArgument(argsCount int, name string, args ...*LogicalExpression) error {
	if len(args) != argsCount {
		return fmt.Errorf("function %s required only %d argument", name, argsCount)
	}

	return nil
}

func AtLeastSeveralArgument(argsCount int, name string, args ...*LogicalExpression) error {
	if len(args) < argsCount {
		return fmt.Errorf("function %s required at least 1 arguments", name)
	}
	return nil
}

func MatchOneArgument(name string, args ...*LogicalExpression) error {
	return MatchSeveralArgument(1, name, args...)
}

func MatchTwoArgument(name string, args ...*LogicalExpression) error {
	return MatchSeveralArgument(2, name, args...)
}

func MatchTwelveArgument(name string, args ...*LogicalExpression) error {
	return MatchSeveralArgument(12, name, args...)
}

func AtLeastOneArgument(name string, args ...*LogicalExpression) error {
	return AtLeastSeveralArgument(1, name, args...)
}
