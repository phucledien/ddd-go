package product_test

import (
	"errors"
	"testing"

	"github.com/phucledien/tavern/domain/product"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "should return error if name is empty",
			name:        "",
			expectedErr: product.ErrMissingValues,
		}, {
			test:        "valid values",
			name:        "test",
			description: "test",
			price:       1.0,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		_, err := product.New(tc.name, tc.description, tc.price)
		if !errors.Is(err, tc.expectedErr) {
			t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
		}
	}
}
