package main

import (
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	given_ts := timeZone()
	secondsEastOfUTC := int((time.Duration(10) * time.Minute).Seconds())
	timezone := time.FixedZone("", secondsEastOfUTC)
	expected := time.Now().In(timezone)

	if given_ts.String() != expected.String() {
		t.Errorf("Got '%v', expected '%v'", given_ts, expected)
	}
}
