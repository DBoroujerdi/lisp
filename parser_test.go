package main

import "testing"

func TestParseSingleSymbol(t *testing.T) {

	result, _ := parse("foo")

	if result != "foo" {
		t.Error("Result should be \"foo\" but was")
	}
}

func TestParseBoolean(t *testing.T) {

	tru, _ := parse("#t")

	if tru != true {
		t.Error("Result should be \"true\" but was")
	}
}
