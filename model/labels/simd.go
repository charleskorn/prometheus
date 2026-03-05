package labels

import "github.com/ajroetker/go-highway/hwy"

// Generate implementations from this file with ./bin/hwygen -input simd.go -output . -targets avx2,avx512,neon:asm,fallback -keepc

func BaseFindFirstDifferentByte(shorter, longer []byte) int {
	sv := hwy.LoadSlice([]byte(shorter))
	lv := hwy.LoadSlice([]byte(longer[:len(shorter)]))

	notEqual := hwy.NotEqual(sv, lv)
	return hwy.FindFirstTrue(notEqual)
}
