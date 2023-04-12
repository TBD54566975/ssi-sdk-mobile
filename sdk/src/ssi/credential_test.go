package ssi

import (
	"testing"

	"github.com/TBD54566975/ssi-sdk/credential"
	"github.com/TBD54566975/ssi-sdk/crypto"
	ssi "github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestSignAndVerifyVCJWT(t *testing.T) {
	did := "test-did"
	kid := "test-key-id"

	// Create a new public & private JWK
	_, privKey, err := ssi.GenerateEd25519Key()
	assert.NoError(t, err)
	assert.NotEmpty(t, privKey)

	privateJwk, err := crypto.PrivateKeyToJWK(privKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, privateJwk)

	publicJwk, err := privateJwk.PublicKey()
	assert.NoError(t, err)
	assert.NotEmpty(t, publicJwk)

	privateJwkBytes, err := json.Marshal(privateJwk)
	assert.NoError(t, err)
	assert.NotEmpty(t, privateJwkBytes)

	publicJwkBytes, err := json.Marshal(publicJwk)
	assert.NoError(t, err)
	assert.NotEmpty(t, publicJwkBytes)

	// Create a new Verifiable Credential
	vcBytes, err := json.Marshal(getSampleCredential())
	assert.NoError(t, err)
	assert.NotEmpty(t, vcBytes)

	// sign it
	jwt, err := SignVerifiableCredentialJWT(did, kid, privateJwkBytes, vcBytes)
	assert.NoError(t, err)
	assert.NotEmpty(t, jwt)

	// verify it
	vc, err := VerifyVerifiableCredentialJWT(did, publicJwkBytes, jwt)
	assert.NoError(t, err)
	assert.NotEmpty(t, vc)
}

func getSampleCredential() credential.VerifiableCredential {
	return credential.VerifiableCredential{
		Context: []any{"https://www.w3.org/2018/credentials/v1",
			"https://w3id.org/security/suites/jws-2020/v1"},
		ID:             "test-verifiable-credential",
		Type:           []string{"VerifiableCredential"},
		Issuer:         "test-issuer",
		ExpirationDate: "3000-01-01T00:00:00Z",
		IssuanceDate:   "2021-01-01T19:23:24Z",
		CredentialSubject: map[string]any{
			"id":      "test-vc-id",
			"company": "Block",
			"website": "https://block.xyz",
		},
	}
}
