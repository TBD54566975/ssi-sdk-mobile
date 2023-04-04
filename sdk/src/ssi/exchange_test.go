package ssi

import (
	"encoding/json"
	"testing"

	"github.com/TBD54566975/ssi-sdk/credential"
	"github.com/TBD54566975/ssi-sdk/credential/exchange"
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/TBD54566975/ssi-sdk/cryptosuite"
	"github.com/TBD54566975/ssi-sdk/did"
	"github.com/stretchr/testify/assert"
)

func TestBuildPresentationSubmission(t *testing.T) {
	privateKey, _, _ := did.GenerateDIDKey("RSA")
	privateJWK, _ := crypto.PrivateKeyToJWK(privateKey)
	kid := "test-key"

	privateJWKBytes, err := json.Marshal(privateJWK)
	assert.NoError(t, err)
	assert.NotEmpty(t, privateJWKBytes)

	pd := getTestPresentationDefinition()
	claim := getTestPresentationClaim()

	pdBytes, err := json.Marshal(&pd)
	assert.NoError(t, err)
	assert.NotEmpty(t, pdBytes)

	claimsBytes, err := json.Marshal([]exchange.PresentationClaim{claim})
	assert.NoError(t, err)
	assert.NotEmpty(t, claimsBytes)

	bytes, err := BuildPresentationSubmission(kid, privateJWKBytes, pdBytes, claimsBytes, string(exchange.JWTVPTarget))
	assert.NoError(t, err)
	assert.NotEmpty(t, bytes)
}

func getTestPresentationDefinition() exchange.PresentationDefinition {
	return exchange.PresentationDefinition{
		ID: "test-id",
		InputDescriptors: []exchange.InputDescriptor{
			{
				ID: "id-1",
				Constraints: &exchange.Constraints{
					Fields: []exchange.Field{
						{
							Path:    []string{"$.vc.issuer", "$.issuer"},
							ID:      "issuer-input-descriptor",
							Purpose: "need to check the issuer",
						},
					},
				},
			},
		},
	}
}

func getTestPresentationClaim() exchange.PresentationClaim {
	testVC := getTestVerifiableCredential()

	return exchange.PresentationClaim{
		Credential:                    &testVC,
		LDPFormat:                     exchange.LDPVC.Ptr(),
		SignatureAlgorithmOrProofType: string(cryptosuite.JSONWebSignature2020),
	}
}

func getTestVerifiableCredential() credential.VerifiableCredential {
	return credential.VerifiableCredential{
		Context: []any{"https://www.w3.org/2018/credentials/v1",
			"https://w3id.org/security/suites/jws-2020/v1"},
		ID:           "test-verifiable-credential",
		Type:         []string{"VerifiableCredential"},
		Issuer:       "test-issuer",
		IssuanceDate: "2021-01-01T19:23:24Z",
		CredentialSubject: map[string]any{
			"id":      "test-vc-id",
			"company": "Block",
			"website": "https://block.xyz",
		},
	}
}
