package sensitive_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestFormat(t *testing.T) {
	oldFn := sensitive.FormatStringFn
	defer func() {
		sensitive.FormatStringFn = oldFn
	}()
	sensitive.FormatStringFn = func(s sensitive.String, f fmt.State, c rune) {
		sensitive.Format(f, c, string(s))
	}

	tests := []struct {
		formatting string
	}{
		{"%s"},
		{"%q"},
		{"%10s"},
		{"%.3[1]q"},
		{"%#-10v"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.formatting, func(t *testing.T) {
			assert := require.New(t)
			want := fmt.Sprintf(tc.formatting, "value")
			result := fmt.Sprintf(tc.formatting, sensitive.String("value"))
			assert.Equal(want, result)
		})
	}
}
