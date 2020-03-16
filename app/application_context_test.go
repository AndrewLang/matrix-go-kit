package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApplicationContext(t *testing.T) {
	context := NewApplicationContext()

	assert.NotNil(t, context)
	assert.NotNil(t, context.Configuration)

}
