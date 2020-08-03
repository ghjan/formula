package formula

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ghjan/formula/internal/cache"
	"github.com/ghjan/formula/internal/exp"
	//register system functions
	_ "github.com/ghjan/formula/internal/fs"
	"github.com/ghjan/formula/internal/parser"
	"github.com/ghjan/formula/opt"
)

//Expression for build user's input
type Expression struct {
	context            *opt.FormulaContext
	originalExpression string
	parsedExpression   *opt.LogicalExpression
}

//NewExpression create new Expression
func NewExpression(expression string, options ...opt.Option) *Expression {
	return &Expression{
		originalExpression: strings.TrimSpace(expression),
		context:            opt.NewFormulaContext(options...),
	}
}

func (expression *Expression) compile() error {
	//handle empty expression
	if expression.originalExpression == "" {
		expression.parsedExpression = exp.NewEmptyExpression()
		return nil
	}

	//restore expression from cache
	logicExpression := cache.Restore(expression.context.Option, expression.originalExpression)
	if logicExpression != nil {
		expression.parsedExpression = logicExpression
		return nil
	}

	//compile expression
	lexer := parser.NewFormulaLexer(antlr.NewInputStream(expression.originalExpression))
	formulaParser := parser.NewFormulaParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))

	//handle compile error
	errListener := newFormulaErrorListener()
	formulaParser.AddErrorListener(errListener)
	calcContext := formulaParser.Calc()

	if errListener.HasError() {
		return errListener.Error()
	}

	expression.parsedExpression = calcContext.GetRetValue()
	return nil
}

//OriginalString return user's input text
func (expression *Expression) OriginalString() string {
	return expression.originalExpression
}

//AddParameter add user's parameter which is required in the expression
func (expression *Expression) AddParameter(name string, value interface{}) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("argument name can not be empty")
	}

	if _, ok := expression.context.Parameters[name]; ok {
		return fmt.Errorf("argument %s dureplate", name)
	}

	expression.context.Parameters[name] = value
	return nil
}

//ResetParameters clear all parameter
func (expression *Expression) ResetParameters() {
	expression.context.Parameters = make(map[string]interface{})
}

//GetParameters get all parameter
func (expression *Expression) GetParameters() []string{
	if	pe,err :=expression.GetParsedExpression();err==nil && pe!=nil{
		return (*pe).Parameters()
	}
	return nil

}

func (expression *Expression) GetParsedExpression() (*opt.LogicalExpression, error) {
	var err error
	if expression.parsedExpression == nil {
		err = expression.compile()
	}
	return expression.parsedExpression, err
}

//Evaluate return result of expression
func (expression *Expression) Evaluate() (*opt.Argument, error) {
	//err := expression.compile()

	parsedExpression, err := expression.GetParsedExpression()
	if err != nil {
		return nil, err
	}
	result, err := (*parsedExpression).Evaluate(expression.context)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type formulaErrorListener struct {
	buf bytes.Buffer
}

func newFormulaErrorListener() *formulaErrorListener {
	return new(formulaErrorListener)
}

func (l *formulaErrorListener) HasError() bool {
	return l.buf.Len() > 0
}

func (l *formulaErrorListener) Error() error {
	return errors.New(l.buf.String())
}

func (l *formulaErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.buf.WriteString(msg)
}

func (*formulaErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (*formulaErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (*formulaErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}
