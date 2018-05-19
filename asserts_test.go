package biff

import "testing"

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

func TestA_AssertEqual(t *testing.T) {

	Alternative("Assert equal", func(a *A) {
		one := 1
		other := 1

		if !a.AssertEqual(one, other) {
			t.Error("AssertEqual should return true")
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

func TestA_AssertNil(t *testing.T) {

	Alternative("Assert nil", func(a *A) {
		one := interface{}(nil)

		if !a.AssertNil(one) {
			t.Error("AssertNil should return true")
		}
	})

}

func TestA_AssertNotEqual(t *testing.T) {

	Alternative("Assert nil", func(a *A) {
		one := 1
		two := 2

		if !a.AssertNotEqual(one, two) {
			t.Error("AssertNotEqual should return true")
		}
	})

}

func TestA_AssertNotNil(t *testing.T) {

	Alternative("Assert equal", func(a *A) {
		one := []string{}

		if !a.AssertNotNil(one) {
			t.Error("AssertEqual should return true")
		}
	})

}
