package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("v1tov32", func() {
	Describe("parseRestartPolicy", func() {

		tests := map[string]string{
			"always":         "",
			"unless-stopped": "",
			"no":             "none",
			"other":          "other",
		}

		It("returns proper values", func() {
			for t, e := range tests {
				Expect(parseRestartPolicy(t)).To(Equal(e))
			}
		})
	})

	Describe("getProto", func() {
		tests := map[string]string{
			"":    "tcp",
			"udp": "udp",
		}

		It("returns proper values", func() {
			for t, e := range tests {
				Expect(getProto(t)).To(Equal(e))
			}
		})
	})
})
