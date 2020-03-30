package sensitive

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUintFormatting(t *testing.T) {
	assert := require.New(t)
	value := Uint(100)
	var empty *Uint

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Uint %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Uint %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Uint %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Uint %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Uint %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Uint %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Uint %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Uint",
		},
		{
			name:       "Ptr Uint %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Uint",
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

func TestUint_MarshalText(t *testing.T) {
	assert := require.New(t)

	value := Uint(100)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestUintJSON(t *testing.T) {
	assert := require.New(t)

	value := Uint(100)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *Uint
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
