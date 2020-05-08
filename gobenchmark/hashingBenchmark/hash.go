package hash

import (
	"path/filepath"

	"github.com/cespare/xxhash"
)

func naiveStringHash(str1, str2 string) string {
	return filepath.Join(str1, str2)
}

func XXhashStringHash(str string) uint64 {
	return xxhash.Sum64String(str)
}
