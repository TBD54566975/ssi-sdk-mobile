package ssi

import (
	"github.com/TBD54566975/ssi-sdk/credential"
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/goccy/go-json"
)

type DIDDocumentMobile struct {
	Context              *StringArray                `json:"@context,omitempty"`
	ID                   string                      `json:"id,omitempty"`
	Controller           string                      `json:"controller,omitempty"`
	AlsoKnownAs          string                      `json:"alsoKnownAs,omitempty"`
	VerificationMethod   *VerificationMethodArray    `json:"verificationMethod,omitempty" validate:"dive"`
	Authentication       *VerificationMethodSetArray `json:"authentication,omitempty" validate:"dive"`
	AssertionMethod      *VerificationMethodSetArray `json:"assertionMethod,omitempty" validate:"dive"`
	KeyAgreement         *VerificationMethodSetArray `json:"keyAgreement,omitempty" validate:"dive"`
	CapabilityInvocation *VerificationMethodSetArray `json:"capabilityInvocation,omitempty" validate:"dive"`
	CapabilityDelegation *VerificationMethodSetArray `json:"capabilityDelegation,omitempty" validate:"dive"`
	Services             *ServiceSetArray            `json:"service,omitempty" validate:"dive"`
}

type VerifiableCredentialMobile struct {
	Context *StringArray `json:"@context" validate:"required"`
	ID      string       `json:"id,omitempty"`
	Type    *StringArray `json:"type" validate:"required"`

	// Mika: Problem type
	// Either a string or an object
	Issuer []byte `json:"issuer" validate:"required"`

	IssuanceDate   string `json:"issuanceDate" validate:"required"`
	ExpirationDate string `json:"expirationDate,omitempty"`

	// Mika: Problem type
	// Requires id and type, but anything else is fair game
	CredentialStatus []byte `json:"credentialStatus,omitempty" validate:"omitempty,dive"`

	// Mika: Problem type
	// type is: map[string]interface{}
	CredentialSubject []byte `json:"credentialSubject" validate:"required"`

	// Mika: Problem type
	// These don't work, even though the structs only contain supported types
	CredentialSchema CredentialSchema `json:"credentialSchema,omitempty" validate:"omitempty,dive"`
	RefreshService   RefreshService   `json:"refreshService,omitempty" validate:"omitempty,dive"`

	// Mika: Problem type
	TermsOfUse []byte `json:"termsOfUse,omitempty" validate:"omitempty,dive"`

	// Mika: Problem type
	Evidence []byte `json:"evidence,omitempty" validate:"omitempty,dive"`

	// Mika: Problem type
	Proof []byte `json:"proof,omitempty"`
}

func (v *VerifiableCredentialMobile) CreateVerifiableCredential() error {
	var issuer any
	if issuerBytes, err := json.Marshal(v.Issuer); err != nil {
		if err = json.Unmarshal(issuerBytes, &issuer); err != nil {
			return err
		}
	}
	var credentialStatus any
	if credentialStatusBytes, err := json.Marshal(v.CredentialStatus); err != nil {
		if err = json.Unmarshal(credentialStatusBytes, &credentialStatus); err != nil {
			return err
		}
	}
	var credentialSubject credential.CredentialSubject
	if credentialSubjectBytes, err := json.Marshal(v.CredentialSubject); err != nil {
		if err = json.Unmarshal(credentialSubjectBytes, &credentialSubject); err != nil {
			return err
		}
	}
	var termsOfUse []credential.TermsOfUse
	if termsOfUseBytes, err := json.Marshal(v.TermsOfUse); err != nil {
		if err = json.Unmarshal(termsOfUseBytes, &termsOfUse); err != nil {
			return err
		}
	}
	var evidence []interface{}
	if evidenceBytes, err := json.Marshal(v.Evidence); err != nil {
		if err = json.Unmarshal(evidenceBytes, &evidence); err != nil {
			return err
		}
	}
	var proof crypto.Proof
	if proofBytes, err := json.Marshal(v.Proof); err != nil {
		if err = json.Unmarshal(proofBytes, &proof); err != nil {
			return err
		}
	}

	cred := &credential.VerifiableCredential{
		Context:           v.Context.fromJSON(),
		ID:                v.ID,
		Type:              v.Type.fromJSON(),
		Issuer:            issuer,
		IssuanceDate:      v.IssuanceDate,
		ExpirationDate:    v.ExpirationDate,
		CredentialStatus:  credentialStatus,
		CredentialSubject: credentialSubject,
		CredentialSchema:  v.CredentialSchema.fromJSON(),
		RefreshService:    v.RefreshService.fromJSON(),
		TermsOfUse:        termsOfUse,
		Evidence:          evidence,
		Proof:             &proof,
	}

	return cred.IsValid()
}

type CredentialSchema struct {
	ID   string `json:"id" validate:"required"`
	Type string `json:"type" validate:"required"`
}

func (c *CredentialSchema) fromJSON() *credential.CredentialSchema {
	return &credential.CredentialSchema{
		ID:   c.ID,
		Type: c.Type,
	}
}

type RefreshService struct {
	ID   string `json:"id" validate:"required"`
	Type string `json:"type" validate:"required"`
}

func (r RefreshService) fromJSON() *credential.RefreshService {
	return &credential.RefreshService{
		ID:   r.ID,
		Type: r.Type,
	}
}
