import { NativeModules, Platform } from 'react-native';

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

const Ssi = SsiModule
  ? SsiModule
  : new Proxy(
      {},
      {
        get() {
          throw new Error(LINKING_ERROR);
        },
      }
    );

// need to expand
type KeyType = 'ed25519' | 'secp256k1' | 'x25519';
export function generateDidKey(keyType: KeyType) {
  try {
    return Ssi.generateDidKey(keyType);
  } catch (e) {
    return e;
  }
}
