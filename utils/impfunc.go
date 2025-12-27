package utils

import (
	//"os"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"runtime"
	"strconv"
	"strings"
	//"bufio"
)

// CompareWithHash compares two hex-encoded hash strings and returns true
// when they are equal. The function logs progress using the package logger.
func CompareWithHash(hash1, hash2 *string) bool {
	Infof("Starting to comapre for the string %s and %s", *hash1, *hash2)

	if *hash1 == *hash2 {
		Infof("Hashes match!")
		return true
	} else {
		Warnf("Hashes do not match!.. Ignoring")
		return false
	}

}

// ConvertToHash computes the SHA-256 hash of the provided password and
// returns the hex-encoded string pointer. The function logs debug/info
// messages to aid tracing during the demo runs.
func ConvertToHash(password *string) *string {
	Infof("Converting the %s to hash ", *password)
	Debugf("Initiating the sha256 algorithm to convert to hash bytes")
	hash_bytes := sha256.Sum256([]byte(*password))
	Debugf("Hashbytes created for %s", *password)
	Debugf("initiating the bytes to string encoder for %s", *password)
	hashString := hex.EncodeToString(hash_bytes[:])
	Debugf("Hash string created for %s", *password)
	Infof("Given text has been converted to has successfully")
	return &hashString

}

// EncodeWithBase64 returns a base64 (RawStdEncoding) representation of the
// provided string and logs the operation. It returns a pointer to the
// generated string to preserve the same signature style used in this demo.
func EncodeWithBase64(value *string) *string {
	Infof("Creating the base64 value for the passed text %s", *value)
	Debugf("Initiating the base64 encoder")
	encoded_string := base64.RawStdEncoding.EncodeToString([]byte(*value))
	Debugf("Encoded string generated")
	Infof("The encoding is done for the value %s ")
	return &encoded_string
}

// GetWorkerid attempts to derive a goroutine id by parsing the runtime
// stack. This is a heuristic method intended only for debugging/demo use in
// the example code and is not guaranteed to be stable across Go versions.
func GetWorkerid() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	stk := strings.TrimPrefix(string(buf[:n]), "goroutine ")
	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		Errorlog("cannot get goroutine id: %v", err)
		panic("cannot get goroutine id")
	}
	return int64(id)
}
