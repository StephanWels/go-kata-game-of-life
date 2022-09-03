package gol

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGameOfLife(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GameOfLife Test Suite")
}

var _ = Describe("game of life", func() {
	It("a single live cell should die", func() {
		startingConfiguration := []string{
			".....",
			".x...",
			".....",
		}
		expectedNextGeneration := []string{
			".....",
			".....",
			".....",
		}
		Expect(NextGeneration(startingConfiguration)).To(Equal(expectedNextGeneration))
	})

	It("a stable configuration", func() {
		startingConfiguration := []string{
			".xx..",
			".xx..",
			".....",
		}
		expectedNextGeneration := []string{
			".xx..",
			".xx..",
			".....",
		}
		Expect(NextGeneration(startingConfiguration)).To(Equal(expectedNextGeneration))
	})
})
