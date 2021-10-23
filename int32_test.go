package sensitive_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestInt32Formatting(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.Int32(100)
	var empty *sensitive.Int32

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Int32 %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Int32 %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Int32 %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Int32 %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Int32 %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Int32 %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Int32 %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Int32",
		},
		{
			name:       "Ptr Int32 %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int32 %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int32 %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int32 %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int32 %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int32 %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int32 %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Int32",
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

func TestInt32_MarshalText(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Int32(100)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestInt32JSON(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Int32(100)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *sensitive.Int32
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
