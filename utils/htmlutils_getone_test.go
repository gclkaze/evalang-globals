package utils

import (
	"strings"
	"testing"
)

const getOneHTML = `<html><body><ul><li class="a">one</li><li class="a">two</li></ul><p id="x">only</p></body></html>`

// No match used to panic with "index out of range [0] with length 0".
func TestParseHTMLByXPATHAndGetOneNoMatchIsAnError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked on no match instead of returning an error: %v", r)
		}
	}()
	got, err := ParseHTMLByXPATHAndGetOne(getOneHTML, "//zzz")
	if err == nil {
		t.Fatalf("want an error for an XPath that matches nothing, got %q", got)
	}
	if got != "" {
		t.Errorf("want an empty value alongside the error, got %q", got)
	}
	if !strings.Contains(err.Error(), "//zzz") {
		t.Errorf("the error should name the XPath so the script author can find it: %v", err)
	}
}

func TestParseHTMLByXPATHAndGetOneFindsExactlyOne(t *testing.T) {
	got, err := ParseHTMLByXPATHAndGetOne(getOneHTML, "//p[@id='x']")
	if err != nil || !strings.Contains(got, "only") {
		t.Fatalf("want the matching node, got %q err=%v", got, err)
	}
}

func TestParseHTMLByXPATHAndGetOneManyMatchesIsStillAnError(t *testing.T) {
	_, err := ParseHTMLByXPATHAndGetOne(getOneHTML, "//li")
	if err == nil || !strings.Contains(err.Error(), "more than one") {
		t.Fatalf("want the existing 'more than one' error, got %v", err)
	}
}
