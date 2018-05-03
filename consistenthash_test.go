package consistenthash

import (
	"testing"
)

var testCases = []struct {
	input  []int
	search int
	output int
}{
	{
		input:  []int{},
		search: 0,
		output: -1,
	},
	{
		input:  []int{1},
		search: 1,
		output: 0,
	},
	{
		input:  []int{1},
		search: 0,
		output: -1,
	},
	{
		input:  []int{1, 2, 3, 4},
		search: 3,
		output: 2,
	},
	{
		input:  []int{1, 2, 3, 4},
		search: 5,
		output: -1,
	},
	{
		input:  []int{1, 2, 3, 4},
		search: 4,
		output: 3,
	},
}

func TestBinarySearchInt(t *testing.T) {
	for i, testCase := range testCases {
		out := binarySearchInt(testCase.input, testCase.search)
		if out != testCase.output {
			t.Errorf("err: for case %d,  expect %d, but got %d", i, testCase.output, out)
			continue
		}
		t.Logf("test pass for case %d", i)
	}
}
