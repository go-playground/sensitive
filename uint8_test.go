package sensitive_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestUint8Formatting(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.Uint8(100)
	var empty *sensitive.Uint8

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Uint8 %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Uint8 %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Uint8 %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Uint8 %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Uint8 %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Uint8 %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Uint8 %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Uint8",
		},
		{
			name:       "Ptr Uint8 %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint8 %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint8 %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint8 %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint8 %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint8 %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint8 %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Uint8",
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

func TestUint8_MarshalText(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Uint8(100)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestUint8JSON(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Uint8(100)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *sensitive.Uint8
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
