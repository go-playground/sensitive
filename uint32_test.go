package sensitive_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/powerman/sensitive"
)

func TestUint32Formatting(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	value := sensitive.Uint32(100)
	var empty *sensitive.Uint32

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Uint32 %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Uint32 %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Uint32 %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Uint32 %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Uint32 %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Uint32 %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Uint32 %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Uint32",
		},
		{
			name:       "Ptr Uint32 %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint32 %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint32 %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint32 %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint32 %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint32 %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Uint32 %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Uint32",
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

func TestUint32_MarshalText(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Uint32(100)

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestUint32JSON(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	value := sensitive.Uint32(100)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *sensitive.Uint32
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
