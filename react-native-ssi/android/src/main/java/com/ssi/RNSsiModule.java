package com.ssi;

import android.util.Log;

import androidx.annotation.NonNull;

import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactMethod;

import ssi.DIDKeyWrapper;
import ssi.Ssi;

public class RNSsiModule extends com.ssi.RNSsiSpec {
  public static final String NAME = "RNSsi";

  RNSsiModule(ReactApplicationContext context) {
    super(context);
  }

  @Override
  @NonNull
  public String getName() {
    return NAME;
  }


  // Example method
  // See https://reactnative.dev/docs/native-modules-android
  @ReactMethod
  public void multiply(double a, double b, Promise promise) {
    promise.resolve(a * b);
  }

  @ReactMethod
  public void generateDidKey(String keyType, Promise promise) {
    try {
      DIDKeyWrapper didKeyWrapper = Ssi.generateDIDKey(keyType);
      String didKey = didKeyWrapper.getDIDKey();
      promise.resolve(didKey);
    } catch (Exception e) {
      promise.reject(e);
    }
  }
}
