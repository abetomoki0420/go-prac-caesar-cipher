package cipher

import "testing"

func TestCaesar(t *testing.T) {
	res1 := Caesar("abc", 1)

	if res1 != "bcd" {
		t.Error("expected:", "bcd")
	}

	res2 := Caesar("xAz", 5)

	if res2 != "cFe" {
		t.Error("expected:", "cFe")
	}

	t.Log("Test done")
}
