package ssi

import (
	"testing"

	"github.com/TBD54566975/ssi-sdk/credential"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestSignAndVerifyVCJWT(t *testing.T) {
	keyPair, err := GenerateEd25519Key()
	assert.NoError(t, err)
	assert.NotEmpty(t, keyPair)

	// Create a new Verifiable Credential
	vc := getSampleCredential()
	assert.NotEmpty(t, vc)

	vcBytes, err := json.Marshal(vc)
	assert.NoError(t, err)
	assert.NotEmpty(t, vcBytes)

	// sign it
	vcJWT := SignVerifiableCredentialJWT("test-key-id", keyPair.KeyType, keyPair.PrivKey, vcBytes)
	assert.NotEmpty(t, vcJWT)

	// verify it
	valid := VerifyVerifiableCredentialJWT("test-key-id", keyPair.KeyType, keyPair.PubKey, vcJWT)
	assert.True(t, valid)
}

func getSampleCredential() credential.VerifiableCredential {
	return credential.VerifiableCredential{
		Context: []any{"https://www.w3.org/2018/credentials/v1",
			"https://w3id.org/security/suites/jws-2020/v1"},
		ID:             "test-verifiable-credential",
		Type:           []string{"VerifiableCredential"},
		Issuer:         "test-issuer",
		ExpirationDate: "2021-01-01T00:00:00Z",
		IssuanceDate:   "2021-01-01T19:23:24Z",
		CredentialSubject: map[string]any{
			"id":      "test-vc-id",
			"company": "Block",
			"website": "https://block.xyz",
		},
	}
}
