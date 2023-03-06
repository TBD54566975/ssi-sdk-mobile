import type { TurboModule } from 'react-native';
import { TurboModuleRegistry } from 'react-native';
import type {
  DidDocument,
  GenerateDidKeyResult,
  VerifiableCredential,
} from './types';

export interface Spec extends TurboModule {
  generateDidKey(keyType: string): Promise<GenerateDidKeyResult>;
  expandDidKey(didKey: string): Promise<DidDocument>;
  createVerifiableCredential(): Promise<VerifiableCredential>;
  signVerifiableCredentialJWT(
    keyID: string,
    privateJwk: Record<string, unknown>,
    vc: VerifiableCredential
  ): Promise<string>;
  verifyVerifiableCredentialJWT(
    keyID: string,
    publicJwk: Record<string, unknown>,
    jwt: string
  ): Promise<VerifiableCredential>;
}

export default TurboModuleRegistry.getEnforcing<Spec>('RNSsi');
