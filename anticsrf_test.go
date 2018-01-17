package main

import (
	"strconv"
	"strings"
	"testing"
)

func Test_keygen(t *testing.T) {
	cases := []struct {
		in, out byte
	}{
		{2, 2},
		{6, 6},
		{32, 32},
		{42, 42},
		{52, 52},
		{126, 126},
	}

	for _, c := range cases {
		got := byte(len(random_key(c.in)))
		if got != c.out {
			t.Errorf("random_key(%q) == %q, want %q", c.in, got, c.out)
		}
	}
}

func Test_microtime(t *testing.T) {
	// currently it's 1,516,158,000 series
	// there are         31,557,600 every year
	// therefore the first two digits will stay 15 for a long time
	// the length must be 16 as well
	first2 := "15"
	needlen := 16

	for i := 0; i < 100; i++ {
		got := strconv.FormatInt(microtime(), 10)
		gotlen := len(got)
		if !strings.HasPrefix(got, first2) && needlen != gotlen {
			t.Errorf("microtime startswith %q != %q and len %q != %q", got[:2], first2, gotlen, needlen)
		}
	}
}
