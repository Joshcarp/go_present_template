package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"testing"

	anz_optimised "github.com/Joshcarp/decimalOptimised"
	"github.com/anz-bank/decimal"
	apd "github.com/cockroachdb/apd"
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
	"decimal.Decimal64":      "anz",
	"float64":                "float64",
	"decimal.Decimal":        "shopspringDecimal",
	"decimal.Big":            "ericlagergrenDecimal",
	"apd.Decimal":            "apdDecimal",
	"anz_optimised.DecParts": "anz"}

var typelist = []interface{}{decimal.Decimal64{}, 0.0, shopspring.Decimal{}, ericlagergren.Big{}, apd.Decimal{}, anz_optimised.DecParts{}}
var typeNamelist = []string{"decimal.Decimal64", "float64", "decimal.Decimal", "decimal.Big", "apd.Decimal", "anz_optimised.DecParts"}

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
					a, b := ParseDecimal(testVal.val1, testVal.val2, t)

					// Add to map
					typeMap[typeNamelist[i]] = append(typeMap[typeNamelist[i]], testcase{testVal.testFunc, a, b})
				}
			}
		}

		// Run the arithmetic test of the seperate types
		for i, t := range typelist {

			b.Run(dectest+"_"+prettyNames[typeNamelist[i]], func(b *testing.B) {
				// Run tests
				for j := 0; j < b.N; j++ {
					for _, test := range typeMap[typeNamelist[i]] {
						runtests(test.v1, test.v2, t, test.op)
					}
				}
			})
		}

	}

}

// Parse the vals as type of interface v
func ParseDecimal(val1, val2 string, v interface{}) (a, b interface{}) {
	switch v.(type) {
	case decimal.Decimal64:
		a, _ = decimal.ParseDecimal64(val1)
		b, _ = decimal.ParseDecimal64(val2)
	case float64:
		b, _ = strconv.ParseFloat(val2, 64)
		a, _ = strconv.ParseFloat(val2, 64)
	case shopspring.Decimal:
		a, _ = shopspring.NewFromString(val1)
		b, _ = shopspring.NewFromString(val2)
	case ericlagergren.Big:
		c := ericlagergren.Big{}
		d := ericlagergren.Big{}
		a, _ = c.SetString(val1)
		b, _ = d.SetString(val2)
	case apd.Decimal:
		a, _, _ = apd.NewFromString(val1)
		b, _, _ = apd.NewFromString(val2)
	case anz_optimised.DecParts:
		a, _ = anz_optimised.ParseDecParts(val1)
		b, _ = anz_optimised.ParseDecParts(val1)
	default:

	}
	return
}

// Run the testPaths
func runtests(a, b, c interface{}, op string) {
	if IgnorePanics {
		defer func() {
			if r := recover(); r != nil {
				// fmt.Println("ERROR: PANIC IN", op, a, b) // There are some issues here that i'm still debugging
			}
		}()
	}
	switch a.(type) {
	case decimal.Decimal64:
		execOp(a.(decimal.Decimal64), b.(decimal.Decimal64), c.(decimal.Decimal64), op)
	case float64:
		execOpFloat(a.(float64), b.(float64), c.(float64), op)
	case shopspring.Decimal:
		execOpShop(a.(shopspring.Decimal), b.(shopspring.Decimal), c.(shopspring.Decimal), op)
	case ericlagergren.Big:
		execOpEric(a.(ericlagergren.Big), b.(ericlagergren.Big), c.(ericlagergren.Big), op)
	case apd.Decimal:
		execOpAdp(a.(apd.Decimal), b.(apd.Decimal), c.(apd.Decimal), op)
	case anz_optimised.DecParts:
		execOpDec(a.(anz_optimised.DecParts), b.(anz_optimised.DecParts), c.(anz_optimised.DecParts), op)
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
func execOpEric(a, b, c ericlagergren.Big, op string) {
	switch op {
	case "add":
		a.Add(&a, &b)
	case "multiply":
		a.Mul(&a, &b)
	case "abs":
		a.Abs(&a)
	case "divide":
		a.Quo(&a, &b)
	default:
	}
	// return ericlagergren.Big{}
}
func execOpFloat(a, b, c float64, op string) {
	var e float64
	switch op {
	case "add":
		e = a + b
	case "multiply":
		e = a * b
	case "abs":
		e = math.Abs(a)
	case "divide":
		e = a / b
	default:

	}
	if false {
		println(e)
	}

}

func execOpShop(a, b, c shopspring.Decimal, op string) {
	switch op {
	case "add":
		a.Add(b)
		// return
	case "multiply":
		a.Mul(b)
	case "abs":
		a.Abs()
	case "divide":
		// shopspring.Zero
		// return a.Div(b)
	}
	// shopspring.Zero
}
func execOp(a, b, c decimal.Decimal64, op string) {
	switch op {
	case "add":
		a.Add(b)
	case "multiply":
		a.Mul(b)
	case "abs":
		a.Abs()
	case "divide":
		a.Quo(b)
	default:
	}
	// return decimal.Zero64
}

func execOpDec(a, b, c anz_optimised.DecParts, op string) {
	switch op {
	case "add":
		a.Add(&b)
	case "multiply":
		a.Mul(&b)
	case "abs":
		a.Abs()
	case "divide":
		// return anz_optimised.DecZero
		a.Quo(&b)
	default:
	}
	// return anz_optimised.DecZero
}

func execOpAdp(a, b, c apd.Decimal, op string) {
	switch op {
	case "add":
		r := apd.ErrDecimal{}
		r.Add(&a, &a, &b)
	case "multiply":
		r := apd.ErrDecimal{}
		r.Mul(&a, &a, &b)
	case "abs":
		r := apd.ErrDecimal{}
		r.Abs(&a, &b)
	case "divide":
		r := apd.ErrDecimal{}
		r.Quo(&a, &a, &b)
	default:
	}
	// apd.Decimal{}
}
