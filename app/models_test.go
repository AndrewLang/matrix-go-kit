package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeToJSON(t *testing.T) {
	station := NewTestStationInfo()

	actual := SerializeToJSON(station)
	assert.NotEmpty(t, actual)
}

func TestDeserializeFromJSON(t *testing.T) {
	content := NewTestStationInfo().ToJSON()

	raw := NewTestStationInfo()
	raw.Name = "barsoom"

	any := DeserializeFromJSON(raw, content)

	assert.NotNil(t, any)

	info, ok := any.(*TestStationInfo)
	assert.True(t, ok)
	assert.NotNil(t, info)
	assert.Equal(t, raw.Name, info.Name)
}
