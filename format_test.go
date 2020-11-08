package sensitive

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	oldFn := FormatStringFn
	defer func() {
		FormatStringFn = oldFn
	}()
	FormatStringFn = func(s String, f fmt.State, c rune) {
		Format(f, c, string(s))
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
			result := fmt.Sprintf(tc.formatting, String("value"))
			assert.Equal(want, result)
		})
	}
}
