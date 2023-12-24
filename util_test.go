package advantageair

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetRecursively(t *testing.T) {
	expected := map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": "value"}}}
	actual := setRecursively(map[string]any{}, "value", "foo", "bar", "baz")
	assert.Equal(t, expected, actual)
}

func FuzzSetRecursively(f *testing.F) {
	f.Add("value,foo,bar,baz")
	f.Add("value,lorem,ipsum,dolor,amet")
	f.Add("value,a,b,c,d,e,f,g,h,i")
	f.Add("value,1,a,2,b,3,c")
	f.Fuzz(func(t *testing.T, input string) {
		keysAndValue := strings.Split(input, ",")
		if len(keysAndValue) < 2 {
			t.Skip() // we need at least one key and a value
		}
		value := keysAndValue[0]
		keys := keysAndValue[1:]

		expectedBuilder := strings.Builder{}
		for _, key := range keys {
			expectedBuilder.WriteString("map[")
			expectedBuilder.WriteString(key)
			expectedBuilder.WriteString(":")
		}
		expectedBuilder.WriteString(value)
		for range keys {
			expectedBuilder.WriteString("]")
		}

		expected := expectedBuilder.String()
		actual := fmt.Sprintf("%+v", setRecursively(map[string]any{}, value, keys...))
		assert.Equal(t, expected, actual)
	})
}
