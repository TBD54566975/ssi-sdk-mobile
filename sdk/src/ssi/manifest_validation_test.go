package ssi

import (
	"embed"
	"testing"

	"github.com/TBD54566975/ssi-sdk/credential"
	"github.com/TBD54566975/ssi-sdk/credential/manifest"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata
	testVectors embed.FS
)

const (
	FullApplicationVector string = "full-application.json"
	FullCredentialVector  string = "full-credential.json"
	FullManifestVector    string = "full-manifest.json"
)

func TestIsValidCredentialApplicationForManifest(t *testing.T) {

	t.Run("Credential Application and Credential Manifest Pair Valid", func(tt *testing.T) {
		cm, ca := getValidTestCredManifestCredApplication(tt)

		applicationAndCredsJSONBytes, err := json.Marshal(ca)
		assert.NoError(tt, err)

		cmBytes, err := json.Marshal(cm)
		assert.NoError(tt, err)

		unfulfilledInputDescriptorsJSONBytes, err := IsValidCredentialApplicationForManifest(cmBytes, applicationAndCredsJSONBytes)
		assert.NoError(tt, err)

		unfulfilledInputDescriptorsJSON := make(map[string]any)
		err = json.Unmarshal(unfulfilledInputDescriptorsJSONBytes, &unfulfilledInputDescriptorsJSON)
		assert.NoError(tt, err)
		assert.Empty(tt, unfulfilledInputDescriptorsJSON)
	})

	t.Run("PresentationSubmission DescriptorMap mismatch id", func(tt *testing.T) {
		cm, ca := getValidTestCredManifestCredApplication(tt)
		ca.CredentialApplication.PresentationSubmission.DescriptorMap[0].ID = "badbadid"

		applicationAndCredsJSONBytes, err := json.Marshal(ca)
		assert.NoError(tt, err)

		cmBytes, err := json.Marshal(cm)
		assert.NoError(tt, err)

		unfulfilledInputDescriptorsJSONBytes, err := IsValidCredentialApplicationForManifest(cmBytes, applicationAndCredsJSONBytes)
		assert.Contains(t, err.Error(), "unfulfilled input descriptor")

		unfulfilledInputDescriptorsJSON := make(map[string]any)
		unmarshalJSONErr := json.Unmarshal(unfulfilledInputDescriptorsJSONBytes, &unfulfilledInputDescriptorsJSON)
		assert.NoError(tt, unmarshalJSONErr)
		assert.Len(tt, unfulfilledInputDescriptorsJSON, 1)
	})
}

func getValidTestCredManifestCredApplication(t *testing.T) (manifest.CredentialManifest, manifest.CredentialApplicationWrapper) {
	// manifest
	manifestJSON, err := getTestVector(FullManifestVector)
	require.NoError(t, err)

	var cm manifest.CredentialManifest
	err = json.Unmarshal([]byte(manifestJSON), &cm)

	require.NoError(t, err)
	require.NotEmpty(t, cm)
	require.NoError(t, cm.IsValid())

	// application
	credAppJSON, err := getTestVector(FullApplicationVector)
	require.NoError(t, err)

	var ca manifest.CredentialApplication
	err = json.Unmarshal([]byte(credAppJSON), &ca)

	require.NoError(t, err)
	require.NotEmpty(t, ca)
	require.NoError(t, ca.IsValid())

	vcJSON, err := getTestVector(FullCredentialVector)
	require.NoError(t, err)

	var vc credential.VerifiableCredential
	err = json.Unmarshal([]byte(vcJSON), &vc)

	require.NoError(t, err)
	require.NotEmpty(t, vc)
	require.NoError(t, vc.IsValid())

	return cm, manifest.CredentialApplicationWrapper{CredentialApplication: ca, Credentials: []any{vc}}
}

func getTestVector(fileName string) (string, error) {
	b, err := testVectors.ReadFile("testdata/" + fileName)
	return string(b), err
}
