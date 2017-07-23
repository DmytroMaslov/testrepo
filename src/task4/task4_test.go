package task4

import (
	"testing"
)
func TestGetMaxPalidrom2(t *testing.T) {
	var testData = [] struct{
		input uint64
		expectedResult string
		expectedStatus bool
	}{
		//positive case
		{123321123121156, "321123", true},
		{222222222222, "222222222222", true},
		{2233445566, "66", true},
		{123, "", false},
		//negative case
	}

	for _, test := range testData{
		out, status := GetMaxPalidrom(test.input)
		if (out != test.expectedResult) || (status != test.expectedStatus) {
			t.Errorf("GetMaxPalidrom(%d) Expected: %d status %t  Actual: %s, %t", test.input, test.expectedResult, test.expectedStatus, out, status)
		}
	}
}

