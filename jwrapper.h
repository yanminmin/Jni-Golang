#ifndef JWRAPPER
#define JWRAPPER
#include <jni.h>

jstring createJString( JNIEnv *env, const char* cstring );

#endif