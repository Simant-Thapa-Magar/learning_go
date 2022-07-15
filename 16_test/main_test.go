package main

import (
	"testing"
)

type test struct {
	Name     string
	Dividend float32
	Dividor  float32
	Expected float32
	IsError  bool
}

var params = []test{{"valid-data", 55.0, 11.0, 5, false}, {"invalid-data", 101.0, 0.0, 0.0, true}, {"expect-2", 4.0, 2.0, 2.0, false}}

func TestDivide(t *testing.T) {
	for _, d := range params {
		r, e := divide(d.Dividend, d.Dividor)

		if d.IsError {
			if e == nil {
				t.Error("Expecting error but no error occured")
			}
		} else {
			if e != nil {
				t.Error("Not expecting error but got error", e.Error())
			}
		}

		if r != d.Expected {
			t.Errorf("Expected %f but got %f for division %f/%f", d.Expected, r, d.Dividend, d.Dividor)
		}
	}
}
