package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"testing"

	ericlagergren "github.com/ericlagergren/decimal"
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
	"float64":     "float64",
	"decimal.Big": "ericlagergrenDecimal"}
var fl64 float64
var typelist = []interface{}{fl64, ericlagergren.Big{}}
var typeNamelist = []string{"float64", "decimal.Big"}

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
					typeMap[typeNamelist[i]] = append(typeMap[typeNamelist[i]], ParseDecimal(testVal.val1, testVal.val2, t, testVal.testName))

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
		c, _ := strconv.ParseFloat(val2, 64)
		d, _ := strconv.ParseFloat(val2, 64)

		test.v1 = &c
		test.v2 = &d
	case ericlagergren.Big:
		c := ericlagergren.Big{}
		d := ericlagergren.Big{}
		test.v1, _ = c.SetString(val1)
		test.v2, _ = d.SetString(val2)
		// a = c
		// b = d

	}
	test.op = op
	return
}

// // Run the testPaths
func runtests(test testcase, t interface{}) {
	// if IgnorePanics {
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			// fmt.Println("ERROR: PANIC IN", op, a, b) // There are some issues here that i'm still debugging
	// 		}
	// 	}()
	// }

	switch (t).(type) {
	case float64:
		a := test.v1.(*float64)
		b := test.v2.(*float64)
		execOpFloat(a, b, test.op)
	case ericlagergren.Big:
		a := test.v1.(*ericlagergren.Big)
		b := test.v2.(*ericlagergren.Big)
		execOpEric(a, b, test.op)
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
	var f *ericlagergren.Big
	switch op {
	case "add":

		f = a.Add(a, b)
	case "multiply":
		f = a.Mul(a, b)
	case "abs":
		f = a.Abs(a)
	case "divide":
		f = a.Quo(a, b)
	default:

	}
	if false {
		println(f)
	}
	// return ericlagergren.Big{}
}
func execOpFloat(a, b *float64, op string) {
	var e float64
	switch op {
	case "add":
		e = *a + *b
	case "multiply":
		e = *a * *b
	case "abs":
		// e = math.Abs((float64)a)
	case "divide":
		e = *a / *b
	default:

	}
	if false {
		println(e)
	}

}
