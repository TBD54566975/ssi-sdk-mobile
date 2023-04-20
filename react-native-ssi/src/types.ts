export type GenerateDidKeyResult = {
  did: string;
  publicJwk: Record<string, unknown>;
  privateJwk: Record<string, unknown>;
};

export type DidDocument = {
  '@context'?: string | string[];
  'id'?: string;
  'controller'?: string;
  'alsoKnownAs'?: string;
  'verificationMethod'?: VerificationMethod[];
  'authentication'?: VerificationMethodSet;
  'keyAgreement'?: VerificationMethodSet;
  'capabilityInvocation'?: VerificationMethodSet;
  'capabilityDelegation'?: VerificationMethodSet;
  'services'?: Service[];
};

export type VerificationMethod = {
  id: string;
  type: string;
  controller: string;
  publicKeyBase58?: string;
  publicKeyJwk?: PublicKeyJWK;
  publicKeyMultibase?: string;
  blockchainAccountId?: string;
};

export type VerificationMethodSet = (string | string[] | VerificationMethod)[];

export type PublicKeyJWK = {
  kty: string;
  crv?: string;
  x?: string;
  y?: string;
  n?: string;
  e?: string;
  use?: string;
  key_ops?: string;
  alg?: string;
  kid?: string;
};

export type Service = {
  id: string;
  type: string;
  serviceEndpoint?: string | string[];
  routingKeys?: string[];
  accept?: string[];
};

export type VerifiableCredential = {
  '@context': string | string[];
  'id'?: string;
  'type': string | string[];
  'issuer': string | Issuer;
  'issuanceDate': string;
  'expirationDate'?: string;
  'credentialStatus'?: CredentialStatus;
  'credentialSubject': CredentialSubject | CredentialSubject[];
  'credentialSchema'?: CredentialSchema;
  'refreshService'?: RefreshService;
  'termsOfUse'?: TermsOfUse[];
  'proof'?: Proof | Proof[];
};

export type Issuer = {
  id: string;
};

export type CredentialStatus = {
  id: string;
  type: string | string[];
};

export type CredentialSubject = {
  id?: string;
};

export type CredentialSchema = {
  id: string;
  type: string | string[];
};

export type RefreshService = {
  id: string;
  type: string | string[];
};

export type TermsOfUse = {
  id?: string;
  type: string | string[];
};

export type Proof = {
  type: string | string[];
};

export type CredentialManifest = {
  id: string;
  spec_version: string;
  issuer: CredentialIssuer;
  output_descriptors: OutputDescriptor[];
  name?: string;
  description?: string;
  format?: ClaimFormatDesignation;
  presentation_definition?: PresentationDefinition;
};

export type CredentialIssuer = {
  id: string;
  name?: string;
  styles?: EntityStyle;
  comment?: string;
};

export type EntityStyleImage = {
  uri: string;
  alt?: string;
};

export type EntityStyleColor = {
  color: string;
};

export type EntityStyle = {
  thumbnail?: EntityStyleImage;
  hero?: EntityStyleImage;
  background?: EntityStyleColor;
  text?: EntityStyleColor;
};

export type OutputDescriptor = {
  id: string;
  schema: string;
  name?: string;
  description?: string;
  styles?: EntityStyle;
  display?: DataDisplay;
};

export type DataMappingSchemaNonString = {
  type: 'boolean' | 'number' | 'integer';
};

export type DataMappingSchemaString = {
  type: 'string';
  format?:
    | 'date-time'
    | 'time'
    | 'date'
    | 'email'
    | 'idn-email'
    | 'hostname'
    | 'idn-hostname'
    | 'ipv4'
    | 'ipv6'
    | 'uri'
    | 'uri-reference'
    | 'iri'
    | 'iri-reference';
};

export type DataMappingSchema =
  | DataMappingSchemaString
  | DataMappingSchemaNonString;

export type DataMappingText = {
  text: string;
};

export type DataMappingPath = {
  path: string[];
  fallback?: string;
  schema: DataMappingSchema;
};

export type DisplayMapping = DataMappingPath | DataMappingText;

export type LabeledDisplayMapping = DisplayMapping & {
  label: string;
};

export type DataDisplay = {
  title?: DisplayMapping;
  subtitle?: DisplayMapping;
  description?: DisplayMapping | { text: string };
  properties?: LabeledDisplayMapping[];
};

export type ClaimFormatDesignationAlg = {
  alg: string[];
};

export type ClaimFormatDesignationProof = {
  proof_type: string[];
};

export type ClaimFormatDesignation = {
  jwt?: ClaimFormatDesignationAlg;
  jwt_vc?: ClaimFormatDesignationAlg;
  jwt_vp?: ClaimFormatDesignationAlg;
  ldp_vc?: ClaimFormatDesignationProof;
  ldp_vp?: ClaimFormatDesignationProof;
  ldp?: ClaimFormatDesignationProof;
};

export type PresentationDefinition = {
  id: string;
  input_descriptors: InputDescriptor[];
  name?: string;
  purpose?: string;
  format?: ClaimFormatDesignation;
};

export enum InputDescriptorConstraintStatusDirective {
  REQUIRED = 'required',
  ALLOWED = 'allowed',
  DISALLOWED = 'disallowed',
}

export enum InputDescriptorConstraintDirective {
  REQUIRED = 'required',
  PREFERRED = 'preferred',
}

export type InputDescriptorConstraintStatus = {
  directive: InputDescriptorConstraintStatusDirective;
};

export type InputDescriptorConstraintStatuses = {
  active?: InputDescriptorConstraintStatus;
  suspended?: InputDescriptorConstraintStatus;
  revoked?: InputDescriptorConstraintStatus;
};

export type InputDescriptorConstraintSubjectConstraint = {
  field_id: string[];
  directive: InputDescriptorConstraintDirective;
};

export type InputDescriptorConstraintFilter = {
  type: string;
  format?: string;
  pattern?: string;
  minimum?: string | number;
  minLength?: number;
  maxLength?: number;
  exclusiveMinimum?: string | number;
  exclusiveMaximum?: string | number;
  maximum?: string | number;
  const?: string | number;
  enum?: string[] | number[];
  not?: InputDescriptorConstraintFilter;
};

export type InputDescriptorConstraintField = {
  path: string[];
  id?: string;
  purpose?: string;
  filter?: InputDescriptorConstraintFilter;
  predicate?: InputDescriptorConstraintDirective;
};

export type InputDescriptorConstraints = {
  limit_disclosure?: InputDescriptorConstraintDirective;
  statuses?: InputDescriptorConstraintStatuses;
  subject_is_issuer?: InputDescriptorConstraintDirective;
  is_holder?: InputDescriptorConstraintSubjectConstraint[];
  same_subject?: InputDescriptorConstraintSubjectConstraint[];
  fields?: InputDescriptorConstraintField[];
};

export type InputDescriptor = {
  id: string;
  group?: string;
  name?: string;
  purpose?: string;
  constraints?: InputDescriptorConstraints;
  format?: ClaimFormatDesignation;
};
