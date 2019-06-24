package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/anz-bank/decimal"
	shop "github.com/shopspring/decimal"
)

type decValContainer struct {
	val1, val2, val3, expected, calculated decimal.Decimal64
	calculatedString                       string
	parseError                             error
}

type decValContainerFloat struct {
	val1, val2, val3, expected, calculated float64
	calculatedString                       string
	parseError                             error
}
type decValContainerShop struct {
	val1, val2, val3, expected, calculated shop.Decimal
	calculatedString                       string
	parseError                             error
}

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
	// 	"dectest/ddAbs.decTest",
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

// TODO(joshcarp): This test cannot fail. Proper assertions will be added once the whole suite passes
// TestFromSuite is the master tester for the dectest suite.
func BenchmarkSuiteANZ(b *testing.B) {
	var testList []decValContainer
	var testStrings []testCaseStrings
	for _, file := range tests {
		f, _ := os.Open(file)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			testVal := getInput(scanner.Text())
			if testVal.testFunc != "" {
				dec64vals := convertToDec64(testVal)
				testList = append(testList, dec64vals)
				testStrings = append(testStrings, testVal)
			}
		}
	}
	b.ResetTimer()
	for j := 0; j < 500; j++ {
		for i, val := range testList {
			runTest(val, testStrings[i])
		}
	}
}

func BenchmarkSuiteFloat(b *testing.B) {
	var testList []decValContainerFloat
	var testStrings []testCaseStrings
	for _, file := range tests {
		f, _ := os.Open(file)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			testVal := getInput(scanner.Text())

			if testVal.testFunc != "" {
				dec64vals := convertToDec64Float(testVal)
				testList = append(testList, dec64vals)
				testStrings = append(testStrings, testVal)
			}
		}
	}
	b.ResetTimer()
	for j := 0; j < 500; j++ {
		for i, val := range testList {
			runTestFloat(val, testStrings[i])
		}
	}
}

func convertToDec64Float(testvals testCaseStrings) (dec64vals decValContainerFloat) {
	var err1, err2, err3, expectedErr error
	dec64vals.val1, err1 = strconv.ParseFloat(testvals.val1, 64)
	dec64vals.val2, err2 = strconv.ParseFloat(testvals.val2, 64)
	dec64vals.val3, err3 = strconv.ParseFloat(testvals.val3, 64)
	dec64vals.expected, expectedErr = strconv.ParseFloat(testvals.expectedResult, 64)

	if err1 != nil || err2 != nil || expectedErr != nil {
		dec64vals.parseError = fmt.Errorf("error parsing in test: %s: \nval 1:%s: \nval 2: %s  \nval 3: %s\nexpected: %s ",
			testvals.testName,
			err1,
			err2,
			err3,
			expectedErr)
	}
	return
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

	if len(ans) == 0 {
		roundingRegex := regexp.MustCompile(`(?:rounding:[\s]*)(?P<rounding>[\S]*)`)
		ans = roundingRegex.FindStringSubmatch(line)
		if len(ans) == 0 {
			return testCaseStrings{}
		}
		return testCaseStrings{rounding: ans[1]}
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

// convertToDec64 converts the map object strings to decimal64s.
func convertToDec64(testvals testCaseStrings) (dec64vals decValContainer) {
	var err1, err2, err3, expectedErr error
	dec64vals.val1, err1 = decimal.ParseDecimal64(testvals.val1)
	dec64vals.val2, err2 = decimal.ParseDecimal64(testvals.val2)
	dec64vals.val3, err3 = decimal.ParseDecimal64(testvals.val3)
	dec64vals.expected, expectedErr = decimal.ParseDecimal64(testvals.expectedResult)

	if err1 != nil || err2 != nil || expectedErr != nil {
		dec64vals.parseError = fmt.Errorf("error parsing in test: %s: \nval 1:%s: \nval 2: %s  \nval 3: %s\nexpected: %s ",
			testvals.testName,
			err1,
			err2,
			err3,
			expectedErr)
	}
	return
}

// runTest completes the tests and returns a boolean and string on if the test passes.
func runTest(testVals decValContainer, testValStrings testCaseStrings) {
	execOp(testVals.val1, testVals.val2, testVals.val3, testValStrings.testFunc)
}

// runTest completes the tests and returns a boolean and string on if the test passes.
func runTestFloat(testVals decValContainerFloat, testValStrings testCaseStrings) {
	execOpFloat(testVals.val1, testVals.val2, testVals.val3, testValStrings.testFunc)
}

func execOpFloat(a, b, c float64, op string) decValContainerFloat {
	if IgnorePanics {
		defer func() {
			if r := recover(); r != nil {
			}
		}()
	}
	switch op {
	case "add":
		return decValContainerFloat{calculated: a + b}
	case "multiply":
		return decValContainerFloat{calculated: a * b}
	case "abs":
		return decValContainerFloat{calculated: math.Abs(a)}
	case "divide":
		return decValContainerFloat{calculated: a / b}
	// case "fma":
	// return decValContainer{calculated: a.FMA(b, c)}
	// case "compare":
	// 	return decValContainerFloat{calculatedString: fmt.Sprintf("%d", int64(a.Cmp(b)))}
	default:
	}
	return decValContainerFloat{calculated: 0}
}

// TODO: get runTest to run more functions ssuch as FMA.
func execOp(a, b, c decimal.Decimal64, op string) decValContainer {
	if IgnorePanics {
		defer func() {
			if r := recover(); r != nil {
			}
		}()
	}
	switch op {
	case "add":
		return decValContainer{calculated: a.Add(b)}
	case "multiply":
		return decValContainer{calculated: a.Mul(b)}
	case "abs":
		return decValContainer{calculated: a.Abs()}
	case "divide":
		return decValContainer{calculated: a.Quo(b)}
	case "fma":
		return decValContainer{calculated: a.FMA(b, c)}
	case "compare":
		return decValContainer{calculatedString: fmt.Sprintf("%d", int64(a.Cmp(b)))}
	default:
	}
	return decValContainer{calculated: decimal.Zero64}
}

func BenchmarkSuiteShop(b *testing.B) {
	var testList []decValContainerShop
	var testStrings []testCaseStrings
	for _, file := range tests {
		f, _ := os.Open(file)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			testVal := getInput(scanner.Text())

			if testVal.testFunc != "" {
				dec64vals := convertToShopDec64(testVal)
				testList = append(testList, dec64vals)
				testStrings = append(testStrings, testVal)
			}
		}
	}
	b.ResetTimer()
	for j := 0; j < 500; j++ {
		for i, val := range testList {
			runTestShop(val, testStrings[i])
		}
	}
}
func convertToShopDec64(testvals testCaseStrings) (dec64vals decValContainerShop) {
	var err1, err2, err3, expectedErr error
	dec64vals.val1, err1 = shop.NewFromString(testvals.val1)
	dec64vals.val2, err2 = shop.NewFromString(testvals.val2)
	dec64vals.val3, err3 = shop.NewFromString(testvals.val3)
	dec64vals.expected, expectedErr = shop.NewFromString(testvals.expectedResult)

	if err1 != nil || err2 != nil || expectedErr != nil {
		dec64vals.parseError = fmt.Errorf("error parsing in test: %s: \nval 1:%s: \nval 2: %s  \nval 3: %s\nexpected: %s ",
			testvals.testName,
			err1,
			err2,
			err3,
			expectedErr)
	}
	return
}

// runTest completes the tests and returns a boolean and string on if the test passes.
func runTestShop(testVals decValContainerShop, testValStrings testCaseStrings) {
	execOpShop(testVals.val1, testVals.val2, testVals.val3, testValStrings.testFunc)
}

// TODO: get runTest to run more functions ssuch as FMA.
// execOp returns the calculated answer to the operation as Decimal64.
func execOpShop(a, b, c shop.Decimal, op string) decValContainerShop {
	if IgnorePanics {
		defer func() {
			if r := recover(); r != nil {
			}
		}()
	}
	switch op {
	case "add":
		return decValContainerShop{calculated: a.Add(b)}
	case "multiply":
		return decValContainerShop{calculated: a.Mul(b)}
	case "abs":
		return decValContainerShop{calculated: a.Abs()}
	case "divide":
		return decValContainerShop{calculated: a.Div(b)}
		// case "fma":
		// return decValContainerShop{calculated: a.FMA(b, c)}
	case "compare":
		return decValContainerShop{calculatedString: fmt.Sprintf("%d", int64(a.Cmp(b)))}
	}
	return decValContainerShop{}
}
