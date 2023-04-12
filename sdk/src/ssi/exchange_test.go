package ssi

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/TBD54566975/ssi-sdk/credential"
	"github.com/TBD54566975/ssi-sdk/credential/exchange"
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/TBD54566975/ssi-sdk/cryptosuite"
	"github.com/TBD54566975/ssi-sdk/did"
	"github.com/TBD54566975/ssi-sdk/schema"
	"github.com/stretchr/testify/assert"
)

// TestMain is used to set up schema caching in order to load all schemas locally
func TestMain(m *testing.M) {
	localSchemas, err := schema.GetAllLocalSchemas()
	if err != nil {
		os.Exit(1)
	}
	loader, err := schema.NewCachingLoader(localSchemas)
	if err != nil {
		os.Exit(1)
	}
	loader.EnableHTTPCache()
	os.Exit(m.Run())
}

func TestBuildPresentationSubmission(t *testing.T) {
	privateKey, did, _ := did.GenerateDIDKey("RSA")
	privateJWK, _ := crypto.PrivateKeyToJWK(privateKey)
	kid := "test-key"
	requester := "requester"

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

	bytes, err := BuildPresentationSubmission(string(*did), kid, privateJWKBytes, requester, pdBytes, claimsBytes, string(exchange.JWTVPTarget))
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
