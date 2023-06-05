package logic

import (
	"github.com/stretchr/testify/require"

	"testing"
)

type TestComputeShmac struct {
	name           string
	secret         string
	strToCrypt     string
	wantSMACSHA512 string
}

func Test_testLogic_computeHmac512(t *testing.T) {

	var tests = []TestComputeShmac{
		{
			name:           "positive",
			secret:         "secret",
			strToCrypt:     "password for entry",
			wantSMACSHA512: "YMGxJWuEFvnNm3mzniyV2r21oH+LqzbIH6kybGfbIP55rtnsOipMI7AgcUra9zEI7mNMGLWKxIG6jfVZe/PK1A==",
		},
		{
			name:           "empty string to crypt",
			secret:         "secret",
			strToCrypt:     "",
			wantSMACSHA512: "kq/ZajBPDX9f/Iot50YWbQynEpLisYEF690p1Fc4zZY8HVkYt8ZIOtdT1Qid7NJp2gGUd873ETc263YkfRgEhA==",
		},
		{
			name:           "empty secret",
			secret:         "",
			strToCrypt:     "password for entry",
			wantSMACSHA512: "6204fc4n8ss+nCgPYigYk4ZvB9/HW61Lk14NFq7fy7E/FxAIntRgHchD4ZvopBBBGwjrrQkTD4Y5X1JAIDKxBw==",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := computeHmac512(tt.strToCrypt, tt.name)
			require.Equal(t, tt.wantSMACSHA512, got)

		})
	}
}
