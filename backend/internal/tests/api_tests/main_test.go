package api_tests

import (
	"testing"

	"github.com/KeishiIrisa/backend-go-template/internal/tests/testutils"
)

func TestMain(m *testing.M) {
	// Initialize the test suite with the database
	testutils.InitializeTestSuite(m)
}
