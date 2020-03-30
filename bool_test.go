package sensitive

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoolFormatting(t *testing.T) {
	assert := require.New(t)
	value := Bool(true)
	var empty *Bool

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Bool %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Bool %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Bool %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Bool %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Bool %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Bool %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Bool %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Bool",
		},
		{
			name:       "Ptr Bool %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bool %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bool %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bool %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bool %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bool %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Bool %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Bool",
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

func TestBoolJSON(t *testing.T) {
	assert := require.New(t)
	value := Bool(true)

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *Bool
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}

func TestBoolCustomFormatFn(t *testing.T) {
	assert := require.New(t)

	oldFn := FormatBoolFn
	defer func() {
		FormatBoolFn = oldFn
	}()
	FormatBoolFn = func(s Bool, f fmt.State, c rune) {
		_, _ = f.Write([]byte("blah"))
	}

	value := Bool(true)
	assert.Equal("blah", fmt.Sprintf("%s", value))
}
