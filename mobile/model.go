package mobile

type DIDDocumentMobile struct {
	Context              *StringArray                `json:"@context,omitempty"`
	ID                   string                      `json:"id,omitempty"`
	Controller           string                      `json:"controller,omitempty" json:"controller,omitempty"`
	AlsoKnownAs          string                      `json:"alsoKnownAs,omitempty" json:"also-known-as,omitempty"`
	VerificationMethod   *VerificationMethodArray    `json:"verificationMethod,omitempty" validate:"dive" json:"verification-method,omitempty"`
	Authentication       *VerificationMethodSetArray `json:"authentication,omitempty" validate:"dive" json:"authentication,omitempty"`
	AssertionMethod      *VerificationMethodSetArray `json:"assertionMethod,omitempty" validate:"dive" json:"assertion-method,omitempty"`
	KeyAgreement         *VerificationMethodSetArray `json:"keyAgreement,omitempty" validate:"dive" json:"key-agreement,omitempty"`
	CapabilityInvocation *VerificationMethodSetArray `json:"capabilityInvocation,omitempty" validate:"dive" json:"capability-invocation,omitempty"`
	CapabilityDelegation *VerificationMethodSetArray `json:"capabilityDelegation,omitempty" validate:"dive" json:"capability-delegation,omitempty"`
	Services             *ServiceSetArray            `json:"service,omitempty" validate:"dive" json:"services,omitempty"`
}
