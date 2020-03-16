package matrix

import (
	"crypto/rand"
	"fmt"
)

// UUID represent a UUID
type UUID struct {
	id string
}

// NewUUID create a new uuid from give string
func NewUUID() *UUID {
	uuid := &UUID{}
	return uuid
}

// ID get id value
func (id *UUID) ID() string {
	return id.id
}

// From a exist value
func (id *UUID) From(value string) *UUID {
	id.id = value
	return id
}

// Generate a new UUID
func (id *UUID) Generate() *UUID {
	b := make([]byte, 16)
	rand.Read(b)

	id.id = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return id
}
