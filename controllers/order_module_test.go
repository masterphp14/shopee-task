package controllers_test

import (
	"os"
	"tax-calculator/test"
	"testing"
)

func TestMain(m *testing.M) {
	test.Setup()

	// run tests
	res := m.Run()

	// cleanup
	test.DbClean()

	os.Exit(res)
}

func TestARouting(t *testing.T) {
}
