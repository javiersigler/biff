package biff

import (
	"strings"
	"testing"
)

func Example_jsonEquality() {

	Alternative("Json equality", func(a *A) {

		i := map[string]interface{}{
			"number": int(33),
		}
		f := map[string]interface{}{
			"number": float64(33),
		}

		a.AssertEqualJson(i, f)

	})

	// Output:
	// Case: Json equality
	//     i is map[string]interface {}{"number":33}
	// -------------------------------
}

func newExitFunction() *bool {

	exited := false
	exit = func() {
		exited = true
	}

	return &exited
}

func TestA_AssertEqual(t *testing.T) {

	Alternative("Assert equal", func(a *A) {
		one := 1
		other := 1

		if !a.AssertEqual(one, other) {
			t.Error("AssertEqual should return true")
		}
	})

}

func TestA_AssertEqualFailed(t *testing.T) {

	exited := newExitFunction()

	Alternative("Assert equal", func(a *A) {
		one := 1
		other := 2

		if a.AssertEqual(one, other) {
			t.Error("AssertEqual should return false when fail")
		}

		if !*exited {
			t.Error("Exited should be true when AssertEqual fails")
		}
	})

}

func TestA_AssertEqualJson(t *testing.T) {

	Alternative("Assert equal Json", func(a *A) {
		one := map[string]interface{}{
			"number": int(33),
		}
		other := map[string]interface{}{
			"number": float64(33),
		}

		if !a.AssertEqualJson(one, other) {
			t.Error("AssertEqualJson should return true")
		}
	})

}

func TestA_AssertEqualJsonFailed(t *testing.T) {

	exited := newExitFunction()

	Alternative("Assert equal Json", func(a *A) {
		one := map[string]interface{}{
			"number": 1,
		}
		other := map[string]interface{}{
			"number": 2,
		}

		if a.AssertEqualJson(one, other) {
			t.Error("AssertEqualJson should return false")
		}

		if !*exited {
			t.Error("Exited should be true when AssertEqualJson fails")
		}

	})

}

func TestA_AssertNil(t *testing.T) {

	Alternative("Assert nil", func(a *A) {
		one := interface{}(nil)

		if !a.AssertNil(one) {
			t.Error("AssertNil should return true")
		}
	})

}

func TestA_AssertNilFailed(t *testing.T) {

	exited := newExitFunction()

	Alternative("Assert nil", func(a *A) {
		one := []string{"1"}

		if a.AssertNil(one) {
			t.Error("AssertNil should return false")
		}

		if !*exited {
			t.Error("Exited should be true when AssertNil fails")
		}

	})

}

func TestA_AssertNotEqual(t *testing.T) {

	Alternative("Assert not equal", func(a *A) {
		one := 1
		two := 2

		if !a.AssertNotEqual(one, two) {
			t.Error("AssertNotEqual should return true")
		}
	})

}

func TestA_AssertNotEqualFailed(t *testing.T) {

	exited := newExitFunction()

	Alternative("Assert not equal", func(a *A) {
		one := 1
		two := 1

		if a.AssertNotEqual(one, two) {
			t.Error("AssertNotEqual should return false")
		}

		if !*exited {
			t.Error("Exited should be true when AssertNotEqual fails")
		}

	})

}

func TestA_AssertNotNil(t *testing.T) {

	Alternative("Assert not nil", func(a *A) {
		one := []string{}

		if !a.AssertNotNil(one) {
			t.Error("AssertEqual should return true")
		}
	})

}

func TestA_AssertNotNilFailed(t *testing.T) {

	exited := newExitFunction()

	Alternative("Assert not nil", func(a *A) {
		one := interface{}(nil)

		if a.AssertNotNil(one) {
			t.Error("AssertEqual should return false")
		}

		if !*exited {
			t.Error("Exited should be true when AssertNotNil fails")
		}
	})

}

func TestA_AssertInArray(t *testing.T) {

	Alternative("Assert in array", func(a *A) {

		colors := []string{"red", "green", "blue"}

		if !a.AssertInArray("red", colors) {
			t.Error("AssertInArray should return true")
		}
	})

}

func TestA_AssertInArrayFailed(t *testing.T) {

	exited := newExitFunction()

	Alternative("Assert in array", func(a *A) {

		colors := []string{"red", "green", "blue"}

		if a.AssertInArray("orange", colors) {
			t.Error("AssertInArray should return false")
		}

		if !*exited {
			t.Error("Exited should be true when AssertInArray fails")
		}

	})

}

func TestA_AssertTrue(t *testing.T) {

	Alternative("Assert true", func(a *A) {

		value := 1 == 1

		if !a.AssertTrue(value) {
			t.Error("AssertTrue should return true")
		}
	})

}

func TestA_AssertTrueFailed(t *testing.T) {

	exited := newExitFunction()

	Alternative("Assert true", func(a *A) {

		value := false

		if a.AssertTrue(value) {
			t.Error("AssertTrue should return false")
		}

		if !*exited {
			t.Error("Exited should be true when AssertTrue fails")
		}

	})

}

func TestA_AssertFalse(t *testing.T) {

	Alternative("Assert false", func(a *A) {

		value := 1 == 2

		if !a.AssertFalse(value) {
			t.Error("AssertTrue should return true")
		}
	})

}

func TestA_AssertFalseFailed(t *testing.T) {

	exited := newExitFunction()

	Alternative("Assert false", func(a *A) {

		value := true

		if a.AssertFalse(value) {
			t.Error("AssertTrue should return false")
		}

		if !*exited {
			t.Error("Exited should be true when AssertFalse fails")
		}
	})

}

func Test_getStackLine(t *testing.T) {

	func() {
		l := getStackLine(2)
		if !strings.HasPrefix(l, "github.com/fulldump/biff.Test_getStackLine(") {
			t.Error("getStackLine should start by 'github.com/fulldump/biff.Test_getStackLine('")
		}
	}()

}
