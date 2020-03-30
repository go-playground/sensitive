package sensitive

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt64Formatting(t *testing.T) {
	assert := require.New(t)
	value := Int64(100)
	var empty *Int64

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Int64 %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Int64 %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Int64 %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Int64 %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Int64 %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Int64 %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Int64 %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Int64",
		},
		{
			name:       "Ptr Int64 %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int64 %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int64 %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int64 %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int64 %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int64 %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int64 %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Int64",
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

func TestInt64_MarshalText(t *testing.T) {
	assert := require.New(t)

	value := Int64(100)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestInt64JSON(t *testing.T) {
	assert := require.New(t)

	value := Int64(100)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *Int64
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
