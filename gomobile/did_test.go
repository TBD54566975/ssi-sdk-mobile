package gomobile

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDIDKey(t *testing.T) {
	didWrapper, err := GenerateDIDKey("RSA")
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(didWrapper.DIDKey, "did:key:"))
	assert.NotEmpty(t, didWrapper.PrivateJSONWebKey)
}
