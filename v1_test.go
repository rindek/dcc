package main

import (
	"testing"
)

func TestAsCpus(test *testing.T) {
	tests := map[CpuQuota]string{
		250000: "2.5",
		11000:  "0.11",
		123456: "1.23456",
	}

	for t, e := range tests {
		v := t.asCpus()
		if v != e {
			test.Errorf("Value from %d: Expected %s, got %s", t, e, v)
		}
	}
}
