package sensitive

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUint16Formatting(t *testing.T) {
	assert := require.New(t)
	value := Uint16(100)
	var empty *Uint16

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Uint16 %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Uint16 %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Uint16 %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Uint16 %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Uint16 %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Uint16 %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Uint16 %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Uint16",
		},
		{
			name:       "Ptr Uint16 %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint16 %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint16 %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint16 %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint16 %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint16 %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint16 %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Uint16",
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

func TestUint16_MarshalText(t *testing.T) {
	assert := require.New(t)

	value := Uint16(100)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestUint16JSON(t *testing.T) {
	assert := require.New(t)

	value := Uint16(100)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *Uint16
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
