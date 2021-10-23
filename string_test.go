package sensitive_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestStringFormatting(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.String("value")
	var empty *sensitive.String

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "String %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "String %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "String %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "String %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "String %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "String %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "String %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.String",
		},
		{
			name:       "Ptr String %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr String %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr String %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr String %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr String %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr String %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr String %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.String",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := fmt.Sprintf(tc.formatting, tc.value)
			assert.Equal(tc.expected, result)
		})
	}
}

func TestStringJSON(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.String("value")

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("\"\"", string(b))

	var empty *sensitive.String
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}

func TestStringCustomFormatFn(t *testing.T) {
	assert := require.New(t)

	oldFn := sensitive.FormatStringFn
	defer func() {
		sensitive.FormatStringFn = oldFn
	}()
	sensitive.FormatStringFn = func(s sensitive.String, f fmt.State, c rune) {
		sensitive.Format(f, c, "blah")
	}

	value := sensitive.String("value")
	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("\"blah\"", string(b))
}

func BenchmarkString_Format(b *testing.B) {
	value := sensitive.String("value")
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", value)
	}
}

func BenchmarkString_FormatNative(b *testing.B) {
	value := "value"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", value) //nolint:gosimple // Benchmark.
	}
}

func BenchmarkStringJSON(b *testing.B) {
	value := sensitive.String("value")
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(value)
	}
}

func BenchmarkString_JSONNative(b *testing.B) {
	value := "value"
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(value)
	}
}
