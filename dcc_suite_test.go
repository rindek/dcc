package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDcc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dcc Suite")
}
