package main

import (
	"testing"
)

func TestParseRestartPolicy(test *testing.T) {
	tests := map[string]string{
		"always":         "",
		"unless-stopped": "",
		"no":             "none",
		"other":          "other",
	}

	for t, e := range tests {
		v := parseRestartPolicy(t)
		if v != e {
			test.Errorf("Value from %s: Expected %s, got %s", t, e, v)
		}
	}
}
