package eval

import "testing"

func TestAlgo(t *testing.T) {
	tests := map[string]int{
		"2 + 2":                     4,
		"10 * 2":                    20,
		"5 + 10 * 2":                25,
		"100 + 200 + 300 + 400":     1000,
		"8 * 1024":                  8192,
		"4 * 1024 + 1024":           5120,
		"1024 + 1024 + 1024 + 1024": 4096,
		"2 - 2":                     0,
		"4 - 2 + 2":                 4,
		"100":                       100,
		"4 / 2":                     2,
		"( 100 + 200 ) * 2":         600,
	}

	for expression, result := range tests {
		exp, err := rpn(expression)
		if err != nil {
			t.Fatal(err)
		}
		v, err := eval(exp)
		if err != nil {
			t.Fatal(err)
		}
		if v != result {
			t.Errorf("expected %d but got %d", result, v)
		}
	}
}
