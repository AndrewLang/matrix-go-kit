package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	id := NewUUID()

	assert.Equal(t, "", id.id)

	actual := NewUUID().From("jupiter-neptune")

	assert.Equal(t, "jupiter-neptune", actual.id)
}

func TestUUIDGenerate(t *testing.T) {
	actual := NewUUID().Generate()
	value := "e545bc24-d528-f87f-4519-0f0e2c1c5b9c"

	assert.Equal(t, len(value), len(actual.ID()))
}
