package sensitive_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestFloat32Formatting(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.Float32(100.1)
	var empty *sensitive.Float32

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Float32 %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Float32 %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Float32 %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Float32 %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Float32 %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Float32 %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Float32 %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Float32",
		},
		{
			name:       "Ptr Float32 %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float32 %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float32 %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float32 %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float32 %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float32 %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Float32 %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Float32",
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

func TestFloat32_MarshalText(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Float32(100.1)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestFloat32JSON(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Float32(100.1)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *sensitive.Float32
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
