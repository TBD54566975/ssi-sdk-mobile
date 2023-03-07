package ssi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDIDKey(t *testing.T) {
	resultBytes, err := GenerateDIDKey("RSA")
	assert.NoError(t, err)
	assert.NotEmpty(t, resultBytes)
}
