#include "jwrapper.h"

#ifdef JWRAPPER
#define JWRAPPER

jstring createJString( JNIEnv *env, const char* cstring ) {
    return (*env)->NewStringUTF(env, cstring);
}

#endif