package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	given_ts := timeZone()
	secondsEastOfUTC := int((time.Duration(10) * time.Minute).Seconds())
	timezone := time.FixedZone("", secondsEastOfUTC)
	expected := t.In(timezone)
	
	if given_ts.String() != expected {
		t.Errorf("Got '%v', expected '%v'", given_ts, expected)
	}
}

