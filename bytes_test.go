package sensitive_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestBytesFormatting(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.Bytes("value")
	var empty *sensitive.Bytes

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Bytes %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Bytes %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Bytes %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Bytes %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Bytes %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Bytes %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Bytes %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Bytes",
		},
		{
			name:       "Ptr Bytes %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bytes %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bytes %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bytes %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bytes %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bytes %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bytes %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Bytes",
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

func TestBytesJSON(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.Bytes("value")

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *sensitive.Bytes
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}

func TestBytesCustomFormatFn(t *testing.T) {
	assert := require.New(t)

	oldFn := sensitive.FormatBytesFn
	defer func() {
		sensitive.FormatBytesFn = oldFn
	}()
	sensitive.FormatBytesFn = func(s sensitive.Bytes, f fmt.State, c rune) {
		_, _ = f.Write([]byte("blah"))
	}

	value := sensitive.Bytes("value")
	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("\"YmxhaA==\"", string(b))
}

func BenchmarkBytes_Format(b *testing.B) {
	value := sensitive.Bytes("value")
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", value)
	}
}

func BenchmarkBytes_FormatNative(b *testing.B) {
	value := "value"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", value) //nolint:gosimple // Benchmark.
	}
}

func BenchmarkBytesJSON(b *testing.B) {
	value := sensitive.Bytes("value")
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(value)
	}
}

func BenchmarkBytes_JSONNative(b *testing.B) {
	value := "value"
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(value)
	}
}
