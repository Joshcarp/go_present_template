package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"testing"

	anz_optimised "github.com/Joshcarp/decimalOptimised"
	"github.com/anz-bank/decimal"
	ericlagergren "github.com/ericlagergren/decimal"
	shopspring "github.com/shopspring/decimal"
)

type testCaseStrings struct {
	testName       string
	testFunc       string
	val1           string
	val2           string
	val3           string
	expectedResult string
}

const IgnorePanics bool = true

var numloops = 10000
var testPaths = []string{
	"ddAdd.decTest",
	"ddMultiply.decTest",
	"ddAbs.decTest",
	"ddDivide.decTest",
}
var testPathdir = "dectest/"

type testcase struct {
	op     string
	v1, v2 interface{}
}

var prettyNames = map[string]string{
	"float64":           "float64",
	"decimal.Decimal64": "anzDecimal",
	"big.Float":         "bigFloat",
}
var fl64 float64
var typelist = []interface{}{decimal.Decimal64{}, fl64, big.Float{}}
var typeNamelist = []string{"decimal.Decimal64", "float64", "big.Float"}

func BenchmarkDecimal(b *testing.B) {
	// map a type (decimal.Decimal64 eg) to a list of testcases
	typeMap := make(map[string][]testcase)

	// For every arithmetic test
	for _, dectest := range testPaths {
		file, _ := os.Open(testPathdir + dectest)
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			testVal := getInput(scanner.Text())
			if testVal.testName != "" {
				// for every type
				for i, t := range typelist {

					// Convert string to type t
					typeMap[typeNamelist[i]] = append(typeMap[typeNamelist[i]], ParseDecimal(testVal.val1, testVal.val2, t, testVal.testFunc))

					// Add to map
				}
			}
		}

		// Run the arithmetic test of the seperate types
		for i, t := range typelist {

			b.Run(dectest+"_"+prettyNames[typeNamelist[i]], func(b *testing.B) {
				// Run tests
				for j := 0; j < b.N; j++ {
					for _, test := range typeMap[typeNamelist[i]] {
						runtests(test, t)

					}
				}
			})
		}

	}

}

// Parse the vals as type of interface v
func ParseDecimal(val1, val2 string, v interface{}, op string) (test testcase) {
	switch v.(type) {
	case float64:
		c, _ := strconv.ParseFloat(val1, 64)
		d, _ := strconv.ParseFloat(val2, 64)

		test.v1 = c
		test.v2 = d
	case big.Float:
		c := big.Float{}
		d := big.Float{}
		fmt.Sscan(val1, c)
		fmt.Sscan(val2, d)
		test.v1 = c
		test.v2 = d
	case ericlagergren.Big:
		c := ericlagergren.Big{}
		d := ericlagergren.Big{}
		test.v1, _ = c.SetString(val1)
		test.v2, _ = d.SetString(val2)
	case decimal.Decimal64:
		c, _ := decimal.ParseDecimal64(val1)
		d, _ := decimal.ParseDecimal64(val2)
		test.v1 = c
		test.v2 = d
	case shopspring.Decimal:
		c, _ := shopspring.NewFromString(val1)
		d, _ := shopspring.NewFromString(val2)
		test.v1 = &c
		test.v2 = &d
	case anz_optimised.DecParts:
		test.v1, _ = anz_optimised.ParseDecParts(val1)
		test.v2, _ = anz_optimised.ParseDecParts(val1)
		// a = c
		// b = d

	}
	test.op = op
	return
}

// // Run the testPaths
func runtests(test testcase, t interface{}) {
	if IgnorePanics {
		defer func() {
			if r := recover(); r != nil {
				// fmt.Println("ERROR: PANIC IN", op, a, b) // There are some issues here that i'm still debugging
			}
		}()
	}
	switch (t).(type) {
	case float64:
		a := test.v1.(float64)
		b := test.v2.(float64)
		execOpFloat(a, b, test.op)
	case big.Float:
		a := test.v1.(big.Float)
		b := test.v2.(big.Float)
		execOpBig(a, b, test.op)
	case ericlagergren.Big:
		a := test.v1.(*ericlagergren.Big)
		b := test.v2.(*ericlagergren.Big)
		execOpEric(a, b, test.op)
	case decimal.Decimal64:
		a := test.v1.(decimal.Decimal64)
		b := test.v1.(decimal.Decimal64)
		execOp(a, b, test.op)
	case shopspring.Decimal:
		a := test.v1.(*shopspring.Decimal)
		b := test.v1.(*shopspring.Decimal)
		execOpShop(a, b, test.op)
	case anz_optimised.DecParts:
		a := test.v1.(*anz_optimised.DecParts)
		b := test.v1.(*anz_optimised.DecParts)
		execOpDec(a, b, test.op)
	default:
	}

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
func execOpEric(a, b *ericlagergren.Big, op string) {
	if a == nil || b == nil {
		return
	}
	var g ericlagergren.Big
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			g.Add(a, b)
		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			g.Mul(a, b)
		}

	case "abs":
		for i := 0; i < numloops; i++ {
			g.Abs(a)
		}
	case "divide":
		for i := 0; i < numloops; i++ {
			g.Quo(a, b)
		}
	default:

	}
}
func execOp(a, b decimal.Decimal64, op string) {
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			a.Add(b)

		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			a.Mul(b)
		}
	case "abs":
		for i := 0; i < numloops; i++ {
			a.Abs()
		}
	case "divide":
		for i := 0; i < numloops; i++ {
			a.Quo(b)
		}
	default:
	}
	// return decimal.Zero64
}
func execOpBig(a, b big.Float, op string) {

	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			a.Add(&a, &b)

		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			a.Mul(&a, &b)
		}
	case "abs":
		for i := 0; i < numloops; i++ {
			a.Abs(&a)
		}
	case "divide":
		for i := 0; i < numloops; i++ {
			a.Quo(&a, &b)
		}
	default:
	}
	// return decimal.Zero64
}
func execOpShop(a, b *shopspring.Decimal, op string) {
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			a.Add(*b)
		}
		// return
	case "multiply":
		for i := 0; i < numloops; i++ {
			a.Mul(*b)
		}
	case "abs":
		for i := 0; i < numloops; i++ {
			a.Abs()
		}
	case "divide":
		// shopspring.Zero
		for i := 0; i < numloops; i++ {
			a.Div(*b)
		}
	}
}
func execOpFloat(a, b float64, op string) {
	var e float64
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			e = a + b
		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			e = a * b
		}
	case "abs":
		// 		e = math.Abs((float64)a)
	case "divide":
		for i := 0; i < numloops; i++ {
			e = a / b
		}
	default:

	}
	if false {
		println(e)
	}

}

func execOpDec(a, b *anz_optimised.DecParts, op string) {
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			a.Add(b)
		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			a.Mul(b)
		}
	case "abs":
		for i := 0; i < numloops; i++ {
			a.Abs()
		}
	case "divide":
		for i := 0; i < numloops; i++ {
			a.Quo(b)
		}
	default:
	}
}
