package lyra2rev2_hash

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -llyra2rev2_hash
// #include "Lyra2RE.h"
import "C"

import (
	"unsafe"
)

func GetPowHash(hash []byte) []byte {
	powhash := make([]byte, 32)
	C.lyra2re2_hash((*C.char)(unsafe.Pointer(&hash[0])), (*C.char)(unsafe.Pointer(&powhash[0])))
	return powhash
}
