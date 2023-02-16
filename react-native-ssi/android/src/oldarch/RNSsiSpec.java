package com.ssi;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.Promise;

abstract class RNSsiSpec extends ReactContextBaseJavaModule {
  RNSsiSpec(ReactApplicationContext context) {
    super(context);
  }

  public abstract void multiply(double a, double b, Promise promise);
  public abstract void generateDidKey(String keyType, Promise promise);
}
