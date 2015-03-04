// +build android

// Stubs for invoking JNI (Java Native Interface) methods, which all use
// function pointers.  These stubs are meant for languages such as GO
// that cannot invoke function pointers.

#include <jni.h>

// Constructs a new Java object. The method ID indicates which constructor method to invoke.
jobject NewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);

// Invokes an instance (nonstatic) method on a Java object, according to the specified class and method ID.
jobject CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
jint CallIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
jlong CallLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
jboolean CallBooleanMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
void CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);

// Invokes a static method on a Java object, according to the specified class and method ID.
jobject CallStaticObjectMethodA(JNIEnv *env, jclass cls, jmethodID methodID, jvalue *args);
jint CallStaticIntMethodA(JNIEnv *env, jclass cls, jmethodID methodID, jvalue *args);
jlong CallStaticLongMethodA(JNIEnv *env, jclass cls, jmethodID methodID, jvalue *args);
jboolean CallStaticBooleanMethodA(JNIEnv *env, jclass cls, jmethodID methodID, jvalue *args);
void CallStaticVoidMethodA(JNIEnv *env, jclass cls, jmethodID methodID, jvalue *args);

// Returns the class of an object.
jclass GetObjectClass(JNIEnv* env, jobject obj);

// Searches the directories and zip files specified by the CLASSPATH environment
// variable for the class with the specified name.
jclass FindClass(JNIEnv* env, const char* name);

// Returns the method ID for an instance (nonstatic) method of a class or
// interface.
jmethodID GetMethodID(JNIEnv* env, jclass class, const char* name, const char* args);

// Returns the method ID for a static method of a class or interface.
jmethodID GetStaticMethodID(JNIEnv* env, jclass class, const char* name, const char* args);

// Returns the field ID for an instance (nonstatic) field of a class.
jfieldID GetFieldID(JNIEnv* env, jclass class, const char* name, const char* sig);

// Returns the field ID for a static field of a class.
jfieldID GetStaticFieldID(JNIEnv* env, jclass class, const char* name, const char* sig);

// Return the values of an instance (nonstatic) fields of the provided object.
jobject GetObjectField(JNIEnv* env, jobject obj, jfieldID fieldID);
jboolean GetBooleanField(JNIEnv* env, jobject obj, jfieldID fieldID);
jint GetIntField(JNIEnv* env, jobject obj, jfieldID fieldID);
jobject GetStaticObjectField(JNIEnv* env, jclass cls, jfieldID fieldID);

// Constructs a new array holding objects of type jclass.
jobjectArray NewObjectArray(JNIEnv* env, jsize len, jclass class, jobject initialElement);

// Constructs a new byte array.
jbyteArray NewByteArray(JNIEnv* env, jsize len);

// Returns the number of elements in the array.
jsize GetArrayLength(JNIEnv* env, jarray array);

// Returns an element of an Object array.
jobject GetObjectArrayElement(JNIEnv* env, jobjectArray array, jsize index);

// Sets an element of an Object array.
void SetObjectArrayElement(JNIEnv* env, jobjectArray array, jsize index, jobject obj);

// Returns the contents of a provided Java byte array as a primitive array.
jbyte* GetByteArrayElements(JNIEnv* env, jbyteArray array, jboolean *isCopy);

// Informs the VM that the native code no longer needs access to elems.
// If necessary, this function copies back all changes made to elems to the
// original array.  Supported modes are:
//     0               copy back the content and free the elems buffer
//     JNI_COMMIT      copy back the content but do not free the elems buffer
//     JNI_ABORT       free the buffer without copying back the possible changes
void ReleaseByteArrayElements(JNIEnv* env, jbyteArray array, jbyte* elems, jint mode);

// Copies the data from a primitive array into the Java array.
void SetByteArrayRegion(JNIEnv* env, jbyteArray array, jsize start, jsize len, const jbyte* data);

// Returns a pointer to an array of bytes representing the string in modified
// UTF-8 encoding.
const char* GetStringUTFChars(JNIEnv* env, jstring str, jboolean* isCopy);

// Informs the VM that the native code no longer needs access to the byte array.
void ReleaseStringUTFChars(JNIEnv* env, jstring str, const char* utf);

// Constructs a new java.lang.String object from an array of characters in
// modified UTF-8 encoding.
jstring NewStringUTF(JNIEnv* env, const char* str);

// Causes a java.lang.Throwable object to be thrown.
jint Throw(JNIEnv* env, jthrowable obj);

// Constructs an exception object from the specified class with the given
// message and causes that exception to be thrown.
jint ThrowNew(JNIEnv* env, jclass class, const char* msg);

// Creates a new local reference to the object referred to by the obj argument.
jobject NewLocalRef(JNIEnv* env, jobject obj);

// Creates a new global reference to the object referred to by the obj argument.
jobject NewGlobalRef(JNIEnv* env, jobject obj);

// Deletes the global reference pointed to by globalRef.
void DeleteGlobalRef(JNIEnv* env, jobject globalRef);

// Returns the Java VM interface (used in the Invocation API) associated with
// the current thread.
jint GetJavaVM(JNIEnv* env, JavaVM** vm);

// Determines if an exception is being thrown.
jthrowable ExceptionOccurred(JNIEnv* env);

// Clears any exception that is currently being thrown.
void ExceptionClear(JNIEnv* env);

// Attaches the current thread to a Java VM.
jint AttachCurrentThread(JavaVM* jvm, JNIEnv** env, void* args);

// Detaches the current thread from a Java VM.
jint DetachCurrentThread(JavaVM* jvm);

// Gets the Java environment for the current thread. If successful, *env will
// point to the appropriate interface and JNI_OK status will be returned;
// otherwise, *env will be NULL and JNI_EDETACHED or JNI_EVERSION status will
// be returned.
jint GetEnv(JavaVM* jvm, JNIEnv** env, jint version);

// Tests whether an object is an instance of a class.
jboolean IsInstanceOf(JNIEnv *env, jobject obj, jclass class);