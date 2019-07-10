package types

import (
	"testing"
)

func TestHello(t *testing.T) {
	ts := timeZone()
	expected := "2019-07-10 10:44:42.564651745 +1000 AEST"
	if !expected.Compare(ts).Equal(IntZero).(Bool) {
		t.Errorf("Got '%v', expected '%v'", ts, expected)
	}
	else {
		t.Errorf("YAY!")
	}
}

