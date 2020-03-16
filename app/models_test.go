package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockObject struct {
	Name string `json:"name"`
}

func TestSerializeToJSON(t *testing.T) {
	obj := &MockObject{}

	actual := SerializeToJSON(obj)

	assert.NotEmpty(t, actual)
}

func TestDeserializeFromJSON(t *testing.T) {

}
