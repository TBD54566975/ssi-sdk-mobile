#include <jni.h>
#include "react-native-ssi.h"

extern "C"
JNIEXPORT jint JNICALL
Java_com_ssi_SsiModule_nativeMultiply(JNIEnv *env, jclass type, jdouble a, jdouble b) {
    return ssi::multiply(a, b);
}
