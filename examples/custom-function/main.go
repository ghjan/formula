package main

import (
	"github.com/ghjan/formula"
	"github.com/ghjan/formula/opt"
	"log"
)

func init() {
	var f opt.Function = new(CustomFunction)
	err := formula.Register(&f)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	expressionFormula := "CustomFunction(1,2)"
	expression := formula.NewExpression(expressionFormula)
	result, err := expression.Evaluate()
	if err != nil {
		log.Fatal(err)
	}

	v, err := result.Int64()
	if err != nil {
		log.Fatal(err)
	}

	if v != 4 { //CustomFunction: i+j+1
		log.Fatal("error")
	}

	log.Printf("custom function succeed,expressionFormula:%s, result:=%s\n", expressionFormula, result)
}
