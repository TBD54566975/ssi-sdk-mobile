package ssi

import (
	"strings"
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestGenerateDIDKey(t *testing.T) {
	resultBytes, err := GenerateDIDKey("RSA")
	assert.NoError(t, err)
	assert.NotEmpty(t, resultBytes)

	var result map[string]any
	assert.NoError(t, json.Unmarshal(resultBytes, &result))
	assert.True(t, strings.HasPrefix(result["didKey"].(string), "did:key:"))
	assert.NotEmpty(t, result["privateJwk"])
}
