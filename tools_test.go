package toolkit

import "testing"

//


if(){
	some data here
}

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("wrong leght")
	}
}
