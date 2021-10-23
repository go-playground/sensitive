package sensitive_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestIntFormatting(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.Int(100)
	var empty *sensitive.Int

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Int %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Int %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Int %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Int %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Int %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Int %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Int %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Int",
		},
		{
			name:       "Ptr Int %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Int %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Int",
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

func TestInt_MarshalText(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Int(100)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestIntJSON(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Int(100)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *sensitive.Int
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
