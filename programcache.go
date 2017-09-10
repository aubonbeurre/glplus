package glplus

import (
	"crypto/md5"
	"encoding/hex"
)

var sProgCache map[string]*ProgramCache

// ProgramCache ...
type ProgramCache struct {
	ReleasingReferenceCount

	program *GPProgram
}

// Delete ...
func (b *ProgramCache) Delete() {
	if b.program != nil {
		b.program.DeleteProgram()
	}
}

func getMD5Hash(text1 string, text2 string) string {
	hasher := md5.New()
	hasher.Write([]byte(text1))
	hasher.Write([]byte(text2))
	return hex.EncodeToString(hasher.Sum(nil))
}
