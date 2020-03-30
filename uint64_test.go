package sensitive

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUint64Formatting(t *testing.T) {
	assert := require.New(t)
	value := Uint64(100)
	var empty *Uint64

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Uint64 %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Uint64 %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Uint64 %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Uint64 %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Uint64 %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Uint64 %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Uint64 %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Uint64",
		},
		{
			name:       "Ptr Uint64 %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint64 %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint64 %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint64 %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint64 %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint64 %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint64 %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Uint64",
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

func TestUint64_MarshalText(t *testing.T) {
	assert := require.New(t)

	value := Uint64(100)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestUint64JSON(t *testing.T) {
	assert := require.New(t)

	value := Uint64(100)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *Uint64
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
