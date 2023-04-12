package com.ssi;

import androidx.annotation.NonNull;

import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.ReadableArray;
import com.facebook.react.bridge.ReadableMap;
import com.facebook.react.bridge.ReadableMapKeySetIterator;
import com.facebook.react.bridge.WritableArray;
import com.facebook.react.bridge.WritableMap;

import com.facebook.react.bridge.WritableNativeArray;
import com.facebook.react.bridge.WritableNativeMap;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.nio.charset.StandardCharsets;
import java.util.Iterator;

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
      byte[] result = Ssi.generateDIDKey(keyType);
      promise.resolve(convertBytesToMap(result));
    } catch (Exception e) {
      promise.reject(e);
    }
  }
  @ReactMethod
  public void expandDidKey(String key, Promise promise) {
    try {
      byte[] result = Ssi.expandDIDKey(key);
      promise.resolve(convertBytesToMap(result));
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void signVerifiableCredentialJWT(String did, String keyId, ReadableMap privateJwk, ReadableMap verifiableCredential, Promise promise) {
    try {
      String jwt = Ssi.signVerifiableCredentialJWT(did, keyId, convertMapToBytes(privateJwk), convertMapToBytes(verifiableCredential));
      promise.resolve(jwt);
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void verifyVerifiableCredentialJWT(String did, ReadableMap publicJwk, String jwt, Promise promise) {
    try {
      byte[] result = Ssi.verifyVerifiableCredentialJWT(did, convertMapToBytes(publicJwk), jwt);
      promise.resolve(convertBytesToMap(result));
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  private WritableMap convertBytesToMap(byte[] bytes) throws JSONException {
    String jsonString = new String(bytes, StandardCharsets.UTF_8);
    JSONObject jObject = new JSONObject(jsonString);
    return convertJsonToMap(jObject);
  }

  private byte[] convertMapToBytes(ReadableMap map) throws JSONException {
    JSONObject jObject = convertMapToJson(map);
    return jObject.toString().getBytes(StandardCharsets.UTF_8);
  }

  // Below was taken from:
  // https://gist.github.com/viperwarp/2beb6bbefcc268dee7ad
  private static WritableMap convertJsonToMap(JSONObject jsonObject) throws JSONException {
    WritableMap map = new WritableNativeMap();

    Iterator<String> iterator = jsonObject.keys();
    while (iterator.hasNext()) {
      String key = iterator.next();
      Object value = jsonObject.get(key);
      if (value instanceof JSONObject) {
        map.putMap(key, convertJsonToMap((JSONObject) value));
      } else if (value instanceof  JSONArray) {
        map.putArray(key, convertJsonToArray((JSONArray) value));
      } else if (value instanceof  Boolean) {
        map.putBoolean(key, (Boolean) value);
      } else if (value instanceof  Integer) {
        map.putInt(key, (Integer) value);
      } else if (value instanceof  Double) {
        map.putDouble(key, (Double) value);
      } else if (value instanceof String)  {
        map.putString(key, (String) value);
      } else {
        map.putString(key, value.toString());
      }
    }
    return map;
  }

  private static WritableArray convertJsonToArray(JSONArray jsonArray) throws JSONException {
    WritableArray array = new WritableNativeArray();

    for (int i = 0; i < jsonArray.length(); i++) {
      Object value = jsonArray.get(i);
      if (value instanceof JSONObject) {
        array.pushMap(convertJsonToMap((JSONObject) value));
      } else if (value instanceof  JSONArray) {
        array.pushArray(convertJsonToArray((JSONArray) value));
      } else if (value instanceof  Boolean) {
        array.pushBoolean((Boolean) value);
      } else if (value instanceof  Integer) {
        array.pushInt((Integer) value);
      } else if (value instanceof  Double) {
        array.pushDouble((Double) value);
      } else if (value instanceof String)  {
        array.pushString((String) value);
      } else {
        array.pushString(value.toString());
      }
    }
    return array;
  }

  private static JSONObject convertMapToJson(ReadableMap readableMap) throws JSONException {
    JSONObject object = new JSONObject();
    ReadableMapKeySetIterator iterator = readableMap.keySetIterator();
    while (iterator.hasNextKey()) {
      String key = iterator.nextKey();
      switch (readableMap.getType(key)) {
        case Null:
          object.put(key, JSONObject.NULL);
          break;
        case Boolean:
          object.put(key, readableMap.getBoolean(key));
          break;
        case Number:
          object.put(key, readableMap.getDouble(key));
          break;
        case String:
          object.put(key, readableMap.getString(key));
          break;
        case Map:
          object.put(key, convertMapToJson(readableMap.getMap(key)));
          break;
        case Array:
          object.put(key, convertArrayToJson(readableMap.getArray(key)));
          break;
      }
    }
    return object;
  }

  private static JSONArray convertArrayToJson(ReadableArray readableArray) throws JSONException {
    JSONArray array = new JSONArray();
    for (int i = 0; i < readableArray.size(); i++) {
      switch (readableArray.getType(i)) {
        case Null:
          break;
        case Boolean:
          array.put(readableArray.getBoolean(i));
          break;
        case Number:
          array.put(readableArray.getDouble(i));
          break;
        case String:
          array.put(readableArray.getString(i));
          break;
        case Map:
          array.put(convertMapToJson(readableArray.getMap(i)));
          break;
        case Array:
          array.put(convertArrayToJson(readableArray.getArray(i)));
          break;
      }
    }
    return array;
  }
}
