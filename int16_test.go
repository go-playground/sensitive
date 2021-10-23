package sensitive_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestInt16Formatting(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.Int16(100)
	var empty *sensitive.Int16

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Int16 %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Int16 %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Int16 %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Int16 %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Int16 %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Int16 %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Int16 %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Int16",
		},
		{
			name:       "Ptr Int16 %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int16 %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int16 %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int16 %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int16 %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int16 %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int16 %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Int16",
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

func TestInt16_MarshalText(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Int16(100)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestInt16JSON(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Int16(100)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *sensitive.Int16
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
