package formula

import (
	"fmt"
	"github.com/ghjan/formula/utils"
	"go/importer"
	"math"
	"testing"

	"github.com/ghjan/formula/internal/fs"
	"github.com/ghjan/formula/opt"
)

func TestNewExpression(t *testing.T) {
	const expString = "1+2"

	exp := NewExpression(expString, opt.IgnoreCase)

	result, err := exp.Evaluate()
	if err != nil {
		t.Fatal(err)
	}

	v, err := result.Int64()
	if err != nil {
		t.Fatal(err)
	}

	if v != 3 {
		t.Fatal()
	}
}

func TestErrorExpression(t *testing.T) {
	exp := NewExpression("1++1")
	result, err := exp.Evaluate()
	if err == nil {
		t.Fatal("error should not be nil")
	}

	if result != nil {
		t.Fatal("result should be nil")
	}
}

func TestCompute(t *testing.T) {
	exp := NewExpression("(1+2)*3/6+3*2")
	result, err := exp.Evaluate()
	if err != nil {
		t.Fatal(err)
	}

	v, err := result.Float64()
	if err != nil {
		t.Fatal(err)
	}

	if v != 7.5 {
		t.Fatal()
	}
}

func TestCommonExpression(t *testing.T) {
	testCases := []struct {
		exp    string
		result float64
	}{
		{"1+2", 3},
		{"1+2-3", 0},
		{"1+2-3*4", -9},
		{"(1+2)*3/6", 1.5},
		{"(1+2)*3/(2+4)", 1.5},
		{"(1+2)*3/(2+4-3+4-1)", 1.5},
		{"(1+2)*3/(2+4-3+4-1*0-1)", 1.5},
		{"(1+2)*3/(2+4-3+4-1*0-1)*0", 0},
	}

	for _, tt := range testCases {
		expression := NewExpression(tt.exp)
		result, err := expression.Evaluate()

		if err != nil {
			t.Fatal(err)
		}

		v, err := result.Float64()
		if err != nil {
			t.Fatal(err)
		}

		if v != tt.result {
			t.Fatal()
		}
	}
}

func TestFunctionExpression(t *testing.T) {
	testCases := []struct {
		exp    string
		result float64
	}{
		{"sin(π/2)", math.Sin(math.Pi / 2)},
		{"cos(π/2)", math.Cos(math.Pi / 2)},
		{"tan(π/2)", math.Tan(math.Pi / 2)},
		{"asin(sin(π/2))", math.Pi / 2},
	}

	for _, tt := range testCases {
		expression := NewExpression(tt.exp)
		result, err := expression.Evaluate()

		if err != nil {
			t.Fatal(err)
		}

		v, err := result.Float64()
		if err != nil {
			t.Fatal(err)
		}

		if v != tt.result {
			t.Fatal()
		}
	}
}

type MyTestCase struct {
	I int
	J int
	K int
}

var testCases = []MyTestCase{
	{1, 1, 1},
	{1, 2, 2},
	{2, 3, 3},
	{3, 4, 4},
	{4, 5, 5},
}

func TestExpressionWithParameter(t *testing.T) {
	const expressionString = "[I]+[J]"
	results := expressTt(expressionString, t, testCases)
	for index, tt := range testCases {
		result := results[index]
		if result != int64(tt.I+tt.J) {
			t.Fatal()
		}
	}
}

func TestExpressionWithParameter2(t *testing.T) {
	const expressionString = "[I]+[J]*[K]"

	results := expressTt(expressionString, t, testCases)
	for index, tt := range testCases {
		result := results[index]
		if result != int64(tt.I+tt.J*tt.K) {
			t.Fatal()
		}
	}
}

func expressTt(expressionString string, t *testing.T, testCases []MyTestCase) (results []int64) {
	expression := NewExpression(expressionString)
	params := expression.GetParameters()
	t.Logf("expression.GetParameters():%#v\n", params)
	for _, tt := range testCases {
		expression.ResetParameters()
		m := utils.StructToMapViaReflect(&tt)
		for _, param := range params {
			err := expression.AddParameter(param, m[param])
			if err != nil {
				t.Fatal(err)
			}
		}
		result, err := expression.Evaluate()
		if err != nil {
			t.Fatal(err)
		}

		v, err := result.Int64()
		if err != nil {
			t.Fatal(err)
		}
		results = append(results, v)
	}
	return
}

func BenchmarkOnePlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expression := NewExpression("1+1")
		result, err := expression.Evaluate()
		if err != nil {
			b.Fatal(err)
		}

		v, err := result.Int64()
		if err != nil {
			b.Fatal(err)
		}
		if v != 2 {
			b.Fatal()
		}
	}
}

func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expression := NewExpression("1")
		result, err := expression.Evaluate()
		if err != nil {
			b.Fatal(err)
		}

		v, err := result.Int64()
		if err != nil {
			b.Fatal(err)
		}
		if v != 1 {
			b.Fatal()
		}
	}
}

func BenchmarkComplexOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expression := NewExpression("(1>0)?(1+2)*4/2+(3+3)/2-1+1*4+99*10:-1")
		result, err := expression.Evaluate()
		if err != nil {
			b.Fatal(err)
		}

		v, err := result.Int64()
		if err != nil {
			b.Fatal(err)
		}
		if v != 1002 {
			b.Fatal()
		}
	}
}

func BenchmarkSin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expression := NewExpression("2*asin(sin(π/2))")
		result, err := expression.Evaluate()
		if err != nil {
			b.Fatal(err)
		}

		v, err := result.Float64()
		if err != nil {
			b.Fatal(err)
		}
		if v != math.Pi {
			b.Fatal()
		}
	}
}

func TestImplements(t *testing.T) {
	//通过断言判断类型是否实现接口或组合了其他结构
	var ii interface{} = fs.NewPlusFunction()
	f, ok := ii.(opt.Function)
	fmt.Println(ok)
	fmt.Println(f.Evaluate(nil))

	d := importer.For("source", nil)
	pkg, err := d.Import("github.com/ghjan/formula/internal/fs")
	if err != nil {
		t.Fatal(err)
	}

	for _, declName := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(declName)
		fmt.Printf("%s	%s\n", obj.Name(), obj.Type())
	}
}

func TestForREADME(t *testing.T) {
	testCases := []string{
		"abs(-1)",
		"acos(sqrt(3)/2)",
		"asin(1/2)",
		"atan(1)",
		"ceil(34)",
		"ceil(3.4)",
		"ceil(3.43)",
		"concat(1,23,hello)",
		"cos(π/3)",
		"3/4",
		"exp(3.3)",
		"floor(2.2)",
		"3 > 2",
		"iif(3 > 2,π,10)",
		"iif(abs(-3) > 2,π,10)",
		"iif(abs(-2) == 2,'等于2','不等于2')",
		"iif('旧房翻新' == '旧房翻新','等于旧房翻新','不等于旧房翻新')",
		"iif('毛坯房' == '旧房翻新','等于毛坯房','不等于毛坯房')",
		"iif('true' == 'true','等于true','不等于true')",
		"iif('true' == 'TRUE','等于true','不等于true')",
		"iif('true' == '1','等于true','不等于true')",
		"iif('true' == '0','等于true','不等于true')",
		"iif(abs(-2) == 2,'等于2','不等于2')",
		"iif('毛坯房' != '旧房翻新','不等于旧房翻新','等于旧房翻新')",
		"iif('旧房翻新' != '毛坯房','不等于毛坯房','等于毛坯房')",
		"iif('true' != 'true','不等于true','等于true')",
		"iif('true' != 'TRUE','不等于true','等于true')",
		"iif('true' != '1','不等于true','等于true')",
		"iif('true' != '0','不等于true','等于true')",
		"iif(abs(-3) != 2,'不等于2','等于2')",
		"iif(abs(-2) != 2,'不等于2','等于2')",
		"iif(aand(in('电梯房','电梯房','楼梯房'),'框架'=='框架'),true,false)",
		"iif(aand(in('电梯房','电梯房','楼梯房'),'砖混'=='框架'),true,false)",
		"oor('true' == '0',3==2,2==2)",
		"oor('true' == '0',3==2,1==2)",
		"aand('true' == '0',3==2,2==2)",
		"aand('true' == '0',3==2,1==2)",
		"aand('true' == '1',3==3,2==2)",
		"aand('true' == 'true',3==3,2==2)",
		"aand('true' == 'true',max(1,2,3)==3,max(0,1,2)==2,min(0,1,2)==0)",
		"aand('true' == 'true',max(1,2,3)<3,max(0,1,2)==2,min(0,1,2)==0)",
		"aand('true' == 'true',max(1,2,3)<=3,max(0,1,2)>=2,min(0,1,2)<=0)",
		"aand('true' == 'true',max(1,2,3)>=3,max(0,1,2)>=2,min(0,1,2)>=0)",
		"in(3,3,4,5)",
		"3 < 2",
		"3 == 2",
		"3 == 3",
		"3 <= 3",
		"3 <= 2",
		"2 <= 3",
		"3.0 >= 3",
		"4.2 >= 4.12",
		"4.12 >= 4.121",
		"3 >= 2",
		"2 >= 3",
		"3 >= 3",
		"3.0 <= 3",
		"4.2 <= 4.12",
		"4.12 <= 4.123",
		"log2(16)",
		"log10(100000)",
		"log(100,10)",
		"ln(2.718281828)",
		"max(-1,2,3.1)",
		"min(-1,2,3.1)",
		"mod(5,2)",
		"3*3.4",
		"5+10",
		"pow(10,2)",
		"round(123.113)",
		"round(123.113,2)",
		"round(123.113,-1)",
		"round(123.113,-2)",
		"round(123.113,3)",
		"roundup(123.113)",
		"roundup(123.113,2)",
		"roundup(123.113,3)",
		"roundup(123.113,-1)",
		"roundup(123.113,-2)",
		"rounddown(123.113)",
		"rounddown(123.113,2)",
		"rounddown(123.113,3)",
		"rounddown(123.113,-1)",
		"rounddown(123.113,-2)",
		"sign(100)",
		"sin(π/6)",
		"3-6",
		"tan(π/4)",
		"truncate(12.3)",
		"2>>1",
		"1<<1",
	}

	for i := 0; i < len(testCases); i++ {
		expression := NewExpression(testCases[i])
		result, err := expression.Evaluate()
		if err != nil {
			t.Log(err)
		}

		if result != nil {
			t.Log(testCases[i], "|", result)
		}
	}
}
