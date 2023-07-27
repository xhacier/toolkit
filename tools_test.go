package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("wrong leght")
	}
<<<<<<< HEAD
		ede

someme
de
=======

>>>>>>> parent of 211e3ea (some changes)
}
