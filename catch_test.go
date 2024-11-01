package main

import "testing"

func TestAttemptCapture(t *testing.T) {
	cases := []struct {
		input             int
		executeAttempts   int
		expectedThreshold int
	}{
		{
			input:             10,
			executeAttempts:   5,
			expectedThreshold: 5,
		}, {
			input:             300,
			executeAttempts:   1000,
			expectedThreshold: 60,
		}, {
			input:             100,
			executeAttempts:   1000,
			expectedThreshold: 200,
		},
	}

	for _, c := range cases {
		successCapture := 0
		for range c.executeAttempts {
			if attemptCapture(c.input) {
				successCapture++
			}
		}
		if c.executeAttempts == c.expectedThreshold {
			if successCapture != c.expectedThreshold {
				t.Errorf("successful captured did not match expected result. %v != %v", successCapture, c.expectedThreshold)
			}
		} else {
			if successCapture > c.expectedThreshold {
				t.Errorf("successful captures exceeded expected threshold. %v vs %v", successCapture, c.expectedThreshold)
			}
		}

	}

}
