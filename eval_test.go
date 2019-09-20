package evalfilter

import (
	"fmt"
	"strings"
	"testing"
)

// TestLess tests uses `>` and `>=`.
func TestLess(t *testing.T) {

	// Dummy structure to test field-access.
	type Object struct {
		Count int
	}

	// Instance of object
	var object Object
	object.Count = 3

	type Test struct {
		Input  string
		Result bool
	}

	tests := []Test{
		{Input: `if ( 1 < 4 ) { return true; }`, Result: true},
		{Input: `if ( 1 < 4 ) { return true; } else { print "help!"; return false;}`, Result: true},
		{Input: `if ( 1.4 < 2 ) { return false; }`, Result: false},
		{Input: `if ( 3 <= 3 ) { return true; }`, Result: true},
		{Input: `if ( 1 <= 3 ) { return false; }`, Result: false},
		{Input: `if ( Count <= 3 ) { print "", ""; return false; }`, Result: false},
		{Input: `if ( len("steve") <= 3 ) { return false; } else { return true; }`, Result: true},
	}

	for _, tst := range tests {

		obj := New(tst.Input)

		ret, err := obj.Run(object)
		if err != nil {
			t.Fatalf("Found unexpected error running test %s\n", err.Error())
		}

		if ret != tst.Result {
			t.Fatalf("Found unexpected result running script")
		}
	}
}

// TestMore tests uses `>` and `>=`.
func TestMore(t *testing.T) {

	// Dummy structure to test field-access.
	type Object struct {
		Count int
	}

	// Instance of object
	object := &Object{Count: 32}

	type Test struct {
		Input  string
		Result bool
	}

	tests := []Test{
		{Input: `if ( 1 > 4 ) { return true; } return false;`, Result: false},
		{Input: `if ( 1.1 > 1 ) { return true; }`, Result: true},
		{Input: `if ( Count > 1 ) { return true; }`, Result: true},
		{Input: `if ( 3 >= 3 ) { return true; }`, Result: true},
		{Input: `if ( 1 >= 3 ) { return false; } return true;`, Result: true},
	}

	for _, tst := range tests {

		obj := New(tst.Input)

		ret, err := obj.Run(object)
		if err != nil {
			t.Fatalf("Found unexpected error running test %s\n", err.Error())
		}

		if ret != tst.Result {
			t.Fatalf("Found unexpected result running script")
		}
	}
}

// TestEq tests uses `==` and `!=`.
func TestEq(t *testing.T) {

	// Dummy structure to test field-access.
	type Object struct {
		Count float64
	}

	// Instance of object
	object := &Object{Count: 12.4}

	type Test struct {
		Input  string
		Result bool
	}

	tests := []Test{
		{Input: `if ( Count == 12.4 ) { return true; } return false;`, Result: true},
		{Input: `if ( len(trim(" steve " ) ) == 5 ) { return true; } return false;`, Result: true},
		{Input: `if ( len(trim() ) == 0 ) { return true; } return false;`, Result: true},
		{Input: `if ( Count == 3 ) { return true; } return false;`, Result: false},
		{Input: `if ( Count != 1 ) { return true; }`, Result: true},
		{Input: `if ( Count != 12.4 ) { return false; } return true;`, Result: true},
		{Input: `if ( "Steve" == "Steve" ) { return true; }`, Result: true},
		{Input: `if ( "Steve" == "Kemp" ) { return true; } return false;`, Result: false},
	}

	for _, tst := range tests {

		obj := New(tst.Input)

		ret, err := obj.Run(object)
		if err != nil {
			t.Fatalf("Found unexpected error running test %s\n", err.Error())
		}

		if ret != tst.Result {
			t.Fatalf("Found unexpected result running script")
		}
	}
}

// TestContains tests uses `~=` and `!~`.
func TestContains(t *testing.T) {

	// Dummy structure to test field-access.
	type Object struct {
		Greeting string
	}

	// Instance of object
	object := &Object{Greeting: "Hello World"}

	type Test struct {
		Input  string
		Result bool
	}

	tests := []Test{
		{Input: `if ( Greeting ~= "World" ) { return true; } return false;`, Result: true},
		{Input: `if ( Greeting ~= "Moi" ) { return true; } return false;`, Result: false},
		{Input: `if ( Greeting !~ "Cake" ) { return true; } return false;`, Result: true},
		{Input: `if ( Greeting ~= "Cake" ) { return true; } return false;`, Result: false},
	}

	for _, tst := range tests {

		obj := New(tst.Input)

		ret, err := obj.Run(object)
		if err != nil {
			t.Fatalf("Found unexpected error running test %s\n", err.Error())
		}

		if ret != tst.Result {
			t.Fatalf("Found unexpected result running script")
		}
	}
}

// TestFunction calls a function
func TestFunction(t *testing.T) {

	// Dummy structure to test field-access.
	type Object struct {
		Count int
	}

	// Instance of object
	var object Object
	object.Count = 3

	type Test struct {
		Input  string
		Result bool
	}

	tests := []Test{
		{Input: `if ( True() ) { return true; } return false;`, Result: true},
		{Input: `if ( True() == false ) { return true; } return false;`, Result: false},
		{Input: `if ( True() ~= "true" ) { return true; } return false;`, Result: true},
		{Input: `True(); return false;`, Result: false},
	}

	for _, tst := range tests {

		obj := New(tst.Input)
		obj.AddFunction("True",
			func(eval *Evaluator, obj interface{}, args []Argument) interface{} {
				return true
			})

		ret, err := obj.Run(object)
		if err != nil {
			t.Fatalf("Found unexpected error running test %s\n", err.Error())
		}

		if ret != tst.Result {
			t.Fatalf("Found unexpected result running script")
		}
	}
}

// TestBool tests a struct-member can be boolean
func TestBool(t *testing.T) {

	// Dummy structure to test field-access.
	type Object struct {
		Valid bool
	}

	// Instances of object
	var object Object
	object.Valid = true

	type Test struct {
		Input  string
		Result bool
	}

	tests := []Test{
		{Input: `if ( Valid == true ) { return true; } return false;`, Result: true},
	}

	for _, tst := range tests {

		obj := New(tst.Input)

		ret, err := obj.Run(object)
		if err != nil {
			t.Fatalf("Found unexpected error running test %s\n", err.Error())
		}

		if ret != tst.Result {
			t.Fatalf("Found unexpected result running script")
		}
	}
}

// TestVariable sets a variable.
func TestVariable(t *testing.T) {

	type Test struct {
		Input  string
		Result bool
	}

	tests := []Test{
		{Input: `if ( len( $name ) == 5 ) { return true; } return false;`, Result: true},
	}

	for _, tst := range tests {

		obj := New(tst.Input)
		obj.SetVariable("name", "Steve")

		ret, err := obj.Run(nil)
		if err != nil {
			t.Fatalf("Found unexpected error running test %s\n", err.Error())
		}

		if ret != tst.Result {
			t.Fatalf("Found unexpected result running script")
		}
	}
}

// TestParseErrors parses bogus programs, and ensures that their errors
// are caught.
func TestParseErrors(t *testing.T) {
	tests := []string{
		`return false`,
		`if`,
		`if ( 1 != 1 `,
		`if ( 1 != 1 )`,
		`if ( 1 != 1 ) {`,
		`if ( 1 != 1 ) { return true ;`,
		`if ( 1 == 1 ) { return true ; } else `,
	}

	for _, tst := range tests {

		fmt.Printf("Parsing: %s\n", tst)
		obj := New(tst)
		_, err := obj.Run(nil)

		if err == nil {
			t.Fatalf("Found no error, but expected to do so")
		}

		if !strings.Contains(err.Error(), "expected") &&
			!strings.Contains(err.Error(), "EOF") {
			t.Fatalf("The error we found didn't match what we expected: %s", err.Error())
		}
	}
}

// TestReturn ensures that scripts with no return are bogus.
func TestReturn(t *testing.T) {
	tests := []string{
		`trim();`,
		``,
	}

	for _, tst := range tests {

		obj := New(tst)
		_, err := obj.Run(nil)

		if err == nil {
			t.Fatalf("Found no error, but expected to do so")
		}

		if !strings.Contains(err.Error(), "failed to terminate with a return statement") {
			t.Fatalf("The error we found didn't match what we expected: %s", err.Error())
		}
	}
}
