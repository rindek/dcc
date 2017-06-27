package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("V1", func() {
	Describe("asCpus", func() {

		tests := map[CpuQuota]string{
			250000: "2.5",
			11000:  "0.11",
			123456: "1.23456",
		}

		It("returns proper values", func() {
			for t, e := range tests {
				Expect(t.AsCpus()).To(Equal(e))
			}
		})
	})
})
