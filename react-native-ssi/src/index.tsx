import { NativeModules, Platform } from 'react-native';
import type { Spec } from './NativeRNSsi';
import type { VerifiableCredential } from './types';

const LINKING_ERROR =
  `The package 'react-native-ssi' doesn't seem to be linked. Make sure: \n\n` +
  Platform.select({ ios: "- You have run 'pod install'\n", default: '' }) +
  '- You rebuilt the app after installing the package\n' +
  '- You are not using Expo Go\n';

// @ts-expect-error
const isTurboModuleEnabled = global.__turboModuleProxy != null;

const SsiModule = isTurboModuleEnabled
  ? require('./NativeRNSsi').default
  : NativeModules.RNSsi;

const Ssi: Spec = SsiModule
  ? SsiModule
  : new Proxy(
      {},
      {
        get() {
          throw new Error(LINKING_ERROR);
        },
      }
    );

type KeyType =
  | 'Ed25519'
  | 'secp256k1'
  | 'X25519'
  | 'RSA'
  | 'P-224'
  | 'P-256'
  | 'P-384'
  | 'P-521';

export function generateDidKey(keyType: KeyType) {
  return Ssi.generateDidKey(keyType);
}

export function expandDidKey(didKey: string) {
  return Ssi.expandDidKey(didKey);
}

export function signVerifiableCredentialJWT(
  keyID: string,
  privateJwk: Record<string, unknown>,
  vc: VerifiableCredential
) {
  return Ssi.signVerifiableCredentialJWT(keyID, privateJwk, vc);
}

export function verifyVerifiableCredentialJWT(
  keyID: string,
  publicJwk: Record<string, unknown>,
  jwt: string
) {
  return Ssi.verifyVerifiableCredentialJWT(keyID, publicJwk, jwt);
}
