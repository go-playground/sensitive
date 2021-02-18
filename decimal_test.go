package sensitive

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestDecimalFormatting(t *testing.T) {
	assert := require.New(t)
	value := Decimal(decimal.NewFromFloat(100.1))
	var empty *Decimal

	tests := []struct {
		name       string
		formatting string
		expected   string
		value      interface{}
	}{
		{
			name:       "Decimal %s",
			formatting: "%s",
			value:      value,
		},
		{
			name:       "Decimal %q",
			formatting: "%q",
			value:      value,
		},
		{
			name:       "Decimal %v",
			formatting: "%v",
			value:      value,
		},
		{
			name:       "Decimal %#v",
			formatting: "%#v",
			value:      value,
		},
		{
			name:       "Decimal %x",
			formatting: "%x",
			value:      value,
		},
		{
			name:       "Decimal %X",
			formatting: "%X",
			value:      value,
		},
		{
			name:       "Decimal %T",
			formatting: "%T",
			value:      value,
			expected:   "sensitive.Decimal",
		},
		{
			name:       "Ptr Decimal %s",
			formatting: "%s",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Decimal %q",
			formatting: "%q",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Decimal %v",
			formatting: "%v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Decimal %#v",
			formatting: "%#v",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Decimal %x",
			formatting: "%x",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Decimal %X",
			formatting: "%X",
			value:      empty,
			expected:   "<nil>",
		},
		{
			name:       "Ptr Decimal %T",
			formatting: "%T",
			value:      empty,
			expected:   "*sensitive.Decimal",
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

func TestDecimal_MarshalText(t *testing.T) {
	assert := require.New(t)

	value := Decimal(decimal.NewFromFloat(100.1))

	b, err := value.MarshalText()
	assert.NoError(err)
	assert.Equal("", string(b))
}

func TestDecimalJSON(t *testing.T) {
	assert := require.New(t)

	value := Decimal(decimal.NewFromFloat(100.1))

	b, err := json.Marshal(value)
	assert.NoError(err)
	assert.Equal("null", string(b))

	var empty *Decimal
	b, err = json.Marshal(empty)
	assert.NoError(err)
	assert.Equal("null", string(b))
}
