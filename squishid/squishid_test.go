package squishid

import (
	"math"
	"testing"
)

func TestSquish(t *testing.T) {
	result := Squish(0)
	if result != "a" {
		t.Errorf("got: %s, want: a", result)
	}

	result = Squish(4 << 60)
	if result != "nnae" {
		t.Errorf("got: %s, want: nnae", result)
	}

	result = Squish(123 | (4 << 60))
	if result != "5d0nae" {
		t.Errorf("got: %s, want: 5d0nae", result)
	}

	result = Squish(14242959524133701664)
	if result != "abcdefghijkmo" {
		t.Errorf("got: %s, want: abcdefghijkmo", result)
	}

	result = Squish(math.MaxUint64)
	if result != "nn9r" {
		t.Errorf("got: %s, want: a", result)
	}
}

func TestRestore(t *testing.T) {
	result, _ := Restore("a")
	if result != 0 {
		t.Errorf("got: %d, want: 0", result)
	}

	result, _ = Restore("nnae")
	if result != 4<<60 {
		t.Errorf("got: %d, want: %d", result, 4<<60)
	}

	result, _ = Restore("NNAE")
	if result != 4<<60 {
		t.Errorf("got: %d, want: %d", result, 4<<60)
	}

	result, _ = Restore("5d0nae")
	if result != 123|(4<<60) {
		t.Errorf("got: %d, want: %d", result, 123|(4<<60))
	}

	result, _ = Restore("abcdefghijkmo")
	if result != 14242959524133701664 {
		t.Errorf("got: %d, want: 14242959524133701664", result)
	}

	result, _ = Restore("nn9r")
	if result != math.MaxUint64 {
		t.Errorf("got: %d, want: %d", result, math.MaxUint64)
	}

	result, err := Restore("nnnae")
	if err == nil {
		t.Errorf("got: %d, want: DecodingError", result)
	}

	result, err = Restore("aaaaaaaaaaaaaa")
	if err == nil {
		t.Errorf("got: %d, want: DecodingError", result)
	}

	result, err = Restore("nnan")
	if err == nil {
		t.Errorf("got: %d, want: DecodingError", result)
	}

	result, err = Restore("nnaaaa")
	if err == nil {
		t.Errorf("got: %d, want: DecodingError", result)
	}
}
