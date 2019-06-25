package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"testing"

	"github.com/anz-bank/decimal"
	shop "github.com/shopspring/decimal"
)

type testCaseStrings struct {
	testName       string
	testFunc       string
	val1           string
	val2           string
	val3           string
	expectedResult string
	rounding       string
}

const PrintFiles bool = true
const PrintTests bool = false
const RunTests bool = true
const IgnorePanics bool = true
const IgnoreRounding bool = false

var tests = []string{
	"dectest/ddAdd.decTest",
	"dectest/ddMultiply.decTest",
	// "dectest/ddFMA.decTest",
	// "dectest/ddClass.decTest",
	// TODO: Implement following tests
	// "dectest/ddCompare.decTest",
	"dectest/ddAbs.decTest",
	// 	"dectest/ddCopysign.decTest",
	"dectest/ddDivide.decTest",
	// 	"dectest/ddLogB.decTest",
	// 	"dectest/ddMin.decTest",
	// 	"dectest/ddMinMag.decTest",
	// 	"dectest/ddMinus.decTest",
}

func (testVal testCaseStrings) String() string {
	return fmt.Sprintf("%s %s %v %v %v -> %v\n", testVal.testName, testVal.testFunc, testVal.val1, testVal.val2, testVal.val3, testVal.expectedResult)
}

var supportedRounding = []string{"half_up", "half_even"}
var ignoredFunctions = []string{"apply"}

func BenchmarkSuiteANZ(b *testing.B) {
	var testList [][2]interface{}
	var typelist = []interface{}{decimal.Decimal64{}, 0.1, shop.Decimal{}}
	type testcase struct {
		op     string
		v1, v2 interface{}
	}
	var typeMap map[reflect.Type][]testcase

	var testfunc []string
	for _, file := range tests {
		f, _ := os.Open(file)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var twoVals [2]interface{}
			testVal := getInput(scanner.Text())
			for _, t := range typelist {
				a, b := ParseDecimal(testVal.val1, testVal.val2, t)
				testfunc = append(testfunc, testVal.testFunc)
				testList = append(testList, twoVals)
				typeMap[reflect.TypeOf(t)] = testcase{testVal.testFunc, a, b}
			}
		}
		for _, t := range typelist {
			b.Run(file+reflect.TypeOf(t).String(), func(b *testing.B) {
				for j := 0; j < 500; j++ {
					va1, typeMap[reflect.TypeOf(t)]
					for i, val := range typeMap {
						runtests(val[0], val[1], val[1], testfunc[i])
					}
				}
			})
		}

	}

}
func ParseDecimal(val1, val2 string, v interface{}) (a, b interface{}) {
	switch v.(type) {
	case decimal.Decimal64:
		a, _ = decimal.ParseDecimal64(val1)
		b, _ = decimal.ParseDecimal64(val2)
	case float64:
		b, _ = strconv.ParseFloat(val2, 64)
		a, _ = strconv.ParseFloat(val2, 64)
	case shop.Decimal:
		a, _ = shop.NewFromString(val1)
		b, _ = shop.NewFromString(val2)
	default:

	}
	return
}

func runtests(a, b, c interface{}, op string) {
	switch a.(type) {
	case decimal.Decimal64:
		execOp(a.(decimal.Decimal64), b.(decimal.Decimal64), c.(decimal.Decimal64), op)
	case float64:
		execOpFloat(a.(float64), b.(float64), c.(float64), op)
	case shop.Decimal:
		execOpShop(a.(shop.Decimal), b.(shop.Decimal), c.(shop.Decimal), op)

	}
}

// TODO: get runTest to run more functions ssuch as FMA.
func execOp(a, b, c decimal.Decimal64, op string) decimal.Decimal64 {
	if IgnorePanics {
		defer func() {
			if r := recover(); r != nil {
			}
		}()
	}
	switch op {
	case "add":

		return a.Add(b)
	case "multiply":
		return a.Mul(b)
	case "abs":
		return a.Abs()
	case "divide":
		return a.Quo(b)
	// case "fma":
	// 	return a.FMA(b, c)
	// case "compare":
	// 	a.Cmp(b)
	// 	return decimal.Zero64
	default:
	}
	return decimal.Zero64
}

// getInput gets the test file and extracts test using regex, then returns a map object and a list of test names.
func getInput(line string) testCaseStrings {
	testRegex := regexp.MustCompile(
		`(?P<testName>dd[\w]*)` + // first capturing group: testfunc made of anything that isn't a whitespace
			`(?:\s*)` + // match any whitespace (?: non capturing group)
			`(?P<testFunc>[\S]*)` + // testfunc made of anything that isn't a whitespace
			`(?:\s*\'?)` + // after can be any number of spaces and quotations if they exist (?: non capturing group)
			`(?P<val1>\+?-?[^\t\f\v\' ]*)` + // first test val is anything that isnt a whitespace or a quoteation mark
			`(?:'?\s*'?)` + // match any quotation marks and any space (?: non capturing group)
			`(?P<val2>\+?-?[^\t\f\v\' ]*)` + // second test val is anything that isnt a whitespace or a quoteation mark
			`(?:'?\s*'?)` +
			`(?P<val3>\+?-?[^->]?[^\t\f\v\' ]*)` + //testvals3 same as 1 but specifically dont match with '->'
			`(?:'?\s*->\s*'?)` + // matches the indicator to answer and surrounding whitespaces (?: non capturing group)
			`(?P<expectedResult>\+?-?[^\r\n\t\f\v\' ]*)`) // matches the answer that's anything that is plus minus but not quotations

	// Add regex to match to  rounding: rounding mode her

	// capturing gorups are testName, testFunc, val1,  val2, and expectedResult)
	ans := testRegex.FindStringSubmatch(line)
	if len(ans) < 6 {
		return testCaseStrings{}
	}
	data := testCaseStrings{
		testName:       ans[1],
		testFunc:       ans[2],
		val1:           ans[3],
		val2:           ans[4],
		val3:           ans[5],
		expectedResult: ans[6],
	}
	return data
}

func execOpFloat(a, b, c float64, op string) float64 {
	if IgnorePanics {
		defer func() {
			if r := recover(); r != nil {
			}
		}()
	}
	switch op {
	case "add":
		return a + b
	case "multiply":
		return a * b
	case "abs":
		return math.Abs(a)
	case "divide":
		return a / b
	default:
	}
	return 0
}

// TODO: get runTest to run more functions ssuch as FMA.
// execOp returns the calculated answer to the operation as Decimal64.
func execOpShop(a, b, c shop.Decimal, op string) shop.Decimal {
	if IgnorePanics {
		defer func() {
			if r := recover(); r != nil {
			}
		}()
	}
	switch op {
	case "add":
		return a.Add(b)
	case "multiply":
		return a.Mul(b)
	case "abs":
		return a.Abs()
	case "divide":
		return a.Div(b)
	}
	return shop.Zero
}
