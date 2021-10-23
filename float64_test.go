package sensitive_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestFloat64Formatting(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.Float64(100.1)
	var empty *sensitive.Float64

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Float64 %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Float64 %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Float64 %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Float64 %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Float64 %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Float64 %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Float64 %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Float64",
		},
		{
			name:       "Ptr Float64 %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float64 %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float64 %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float64 %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float64 %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float64 %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float64 %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Float64",
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

func TestFloat64_MarshalText(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Float64(100.1)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestFloat64JSON(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Float64(100.1)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *sensitive.Float64
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
