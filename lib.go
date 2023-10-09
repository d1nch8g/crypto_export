// filename: lib.go
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"strings"

	"golang.org/x/crypto/blake2b"
)

import "C"

//export Hash
func Hash(cString *C.char) *C.char {
	goString := C.GoString(cString)
	bytes, decodeErr := base64.RawStdEncoding.DecodeString(goString)
	if decodeErr != nil {
		return C.CString("decode error occured, wrong input")
	}
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write(bytes)
	hash := hasher.Sum(nil)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)
	return C.CString(hashBase64)
}

//export Keys
func Keys() *C.char {
	persKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	mesKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	totalString := base64.RawStdEncoding.EncodeToString(
		x509.MarshalPKCS1PrivateKey(persKey)) + "|" +
		base64.RawStdEncoding.EncodeToString(
			x509.MarshalPKCS1PublicKey(&persKey.PublicKey)) + "|" +
		base64.RawStdEncoding.EncodeToString(
			x509.MarshalPKCS1PrivateKey(mesKey)) + "|" +
		base64.RawStdEncoding.EncodeToString(
			x509.MarshalPKCS1PublicKey(&mesKey.PublicKey))
	return C.CString(totalString)
}

//export Sign
func Sign(input *C.char) *C.char {
	split := strings.Split(C.GoString(input), "|")
	mesBytes, mesErr := base64.RawStdEncoding.DecodeString(split[1])
	if mesErr != nil {
		return C.CString("error decoding public key base64")
	}
	privBytes, keyErr := base64.RawStdEncoding.DecodeString(split[1])
	if keyErr != nil {
		return C.CString("error decoding private key base64")
	}
	privKey, privErr := x509.ParsePKCS1PrivateKey(privBytes)
	if privErr != nil {
		return C.CString("error parsing priv key")
	}
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write(mesBytes)
	hash := hasher.Sum(nil)
	signatureBytes, _ := rsa.SignPSS(
		rand.Reader,
		privKey,
		crypto.BLAKE2b_512,
		hash,
		nil,
	)
	return C.CString(base64.RawStdEncoding.EncodeToString(signatureBytes))
}

// //export Encode
// //export Decode

func main() {}
