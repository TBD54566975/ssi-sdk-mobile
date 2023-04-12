package com.ssi;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReadableMap;

abstract class RNSsiSpec extends ReactContextBaseJavaModule {
  RNSsiSpec(ReactApplicationContext context) {
    super(context);
  }

  public abstract void multiply(double a, double b, Promise promise);
  public abstract void generateDidKey(String keyType, Promise promise);
  public abstract void expandDidKey(String key, Promise promise);
  public abstract void signVerifiableCredentialJWT(String did, String keyId, ReadableMap privateJwk, ReadableMap verifiableCredential, Promise promise);
  public abstract void verifyVerifiableCredentialJWT(String did, ReadableMap publicJwk, String jwt, Promise promise);
}
