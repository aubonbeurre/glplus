package glplus

const gPositionAttr uint32 = 1 // prog.GetAttribLocation("position")
const gUVsAttr uint32 = 2      // prog.GetAttribLocation("uvs")
const gNormalsAttr uint32 = 3  // prog.GetAttribLocation("normal")

// StringToArray ...
func StringToArray(s string) []uint8 {
	var srcArray = make([]uint8, len(s)+1)
	for i := 0; i < len(s); i++ {
		srcArray[i] = s[i]
	}
	srcArray[len(s)] = 0
	return srcArray
}
