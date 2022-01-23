package generator

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerator_StringWithCharset(t *testing.T) {
	type testCase struct {
		name           string
		input int
		randomizerReturns []int
		output string
	}

	testCases := []testCase{
		{
			name:           "BCD",
			input: 3,
			randomizerReturns: []int{1, 2, 3},
			output: "BCD",
		},
		{
			name:           "AAA",
			input: 3,
			randomizerReturns: []int{0, 0, 0},
			output: "AAA",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			randomizer := NewMockrandomizer(ctrl)

			calls := []*gomock.Call{{}}

			for _, randomizerReturn := range test.randomizerReturns {
				calls = append(calls, randomizer.
					EXPECT().
					GetRandomInt(len(charset)).Return(randomizerReturn))
			}

			gomock.InOrder(calls...)

			generator := New(randomizer)

			a := generator.StringWithCharset(test.input, charset)
			assert.Equal(t, test.output, a)
		})
	}
}

