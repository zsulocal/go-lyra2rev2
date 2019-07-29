package lyra2rev2_hash

import (
	"encoding/hex"
	"testing"
)

func TestX16r(t *testing.T) {
	header := "000000201453fd45edf4308e17f4a7509f803932e8da06b4dd56c5e8e6231468f9a48734eaa625ecd3f56d4faf989c3d39a534f53cff0d0065f79f8eea243f96ca5b5cbdd4343d5df78f131ae2e69710"
	header_bin, _ := hex.DecodeString(header)
	powhash := GetPowHash(header_bin)
	powhash_hex := hex.EncodeToString(powhash)
	if powhash_hex != "77ff69650bb734681411aee1c15644f65cd4c670b1d0b1f46f0c000000000000" {
		t.Error("Test powhash failed")
	}
}

//
