export type GenerateDidKeyResult = {
  didKey: string;
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
