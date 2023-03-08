import type { TurboModule } from 'react-native';
import { TurboModuleRegistry } from 'react-native';
import type { DidDocument, GenerateDidKeyResult } from './types';

export interface Spec extends TurboModule {
  generateDidKey(keyType: string): Promise<GenerateDidKeyResult>;
  expandDidKey(didKey: string): Promise<DidDocument>;
}

export default TurboModuleRegistry.getEnforcing<Spec>('RNSsi');
