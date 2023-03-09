package pandora_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPandora(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pandora Suite")
}
