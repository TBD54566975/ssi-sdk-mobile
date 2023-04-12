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
  signVerifiableCredentialJWT(
    did: string,
    keyID: string,
    privateJwk: Record<string, unknown>,
    vc: VerifiableCredential
  ): Promise<string>;
  verifyVerifiableCredentialJWT(
    did: string,
    publicJwk: Record<string, unknown>,
    jwt: string
  ): Promise<VerifiableCredential>;
}

export default TurboModuleRegistry.getEnforcing<Spec>('RNSsi');
