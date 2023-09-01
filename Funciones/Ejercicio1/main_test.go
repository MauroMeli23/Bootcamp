package main

import (
	"fmt"
	"testing"
)

func TestTaxesCalculator(t *testing.T) {
	testCases := []struct {
		salary        float64
		expectedTaxes float64
	}{
		{25000, 0},
		{60000, 10200},
		{160000, 36800},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Salary %.2f", tc.salary), func(t *testing.T) {
			actualTaxes := taxesCalculator(tc.salary)
			if actualTaxes != tc.expectedTaxes {
				t.Errorf("For salary %.2f, expected %.2f taxes, but got %.2f", tc.salary, tc.expectedTaxes, actualTaxes)
			}
		})
	}
}
