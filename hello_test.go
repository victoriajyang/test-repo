package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	ts := timeZone()
	expected := "2019-07-10 10:44:42.564651745 +1000 AEST"
	if ts.String() != expected {
		t.Errorf("Got '%v', expected '%v'", ts, expected)
	}
}

