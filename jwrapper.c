#include "jwrapper.h"

#ifdef JWRAPPER
#define JWRAPPER
#include <jni.h>


jstring createJString( JNIEnv *env, const char* cstring ) {
    return (*env)->NewStringUTF(env, cstring);
}
#endif