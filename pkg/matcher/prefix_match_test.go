package matcher

import (
	"sort"
	"testing"
)

func TestLongestMatchingPrefix1(t *testing.T) {
	// this is testcase for success case
	prefixes := []string{"rish", "wjpjf", "sbefkw", "erfds", "feefds"}
	sort.Strings(prefixes)
	word := "rishabh"
	got := LongestPrefix(prefixes, word)
	want := "rish"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLongestMatchingPrefix2(t *testing.T) {
	// this is testcase for success case with multiple matching prefixes
	prefixes := []string{"rish", "risha", "rishab", "rishab", "bhardwaj", "zaiohfoe"}
	sort.Strings(prefixes)
	word := "rishabh"
	got := LongestPrefix(prefixes, word)
	want := "rishab"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLongestMatchingPrefix3(t *testing.T) {
	// this is testcase for failure case
	prefixes := []string{"rish", "ewjpjf", "befkw"}
	sort.Strings(prefixes)
	word := "bhardwaj"
	got := LongestPrefix(prefixes, word)
	want := "No prefix found"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLongestMatchingPrefix4(t *testing.T) {
	// this is testcase for failure case
	prefixes := []string{"rish", "ewjpjf", "befkw"}
	sort.Strings(prefixes)
	word := "bha"
	got := LongestPrefix(prefixes, word)
	want := "No prefix found"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLongestMatchingPrefix5(t *testing.T) {
	// this is testcase for failure case
	var prefixes []string
	sort.Strings(prefixes)
	word := "rishabh"
	got := LongestPrefix(prefixes, word)
	want := "No prefix found"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLongestMatchingPrefix6(t *testing.T) {
	// this is testcase for failure case
	prefixes := []string{"abcd", "efgh"}
	sort.Strings(prefixes)
	word := ""
	got := LongestPrefix(prefixes, word)
	want := "No prefix found"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
