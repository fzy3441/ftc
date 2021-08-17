package tt

import (
	"fmt"
	"math/rand"
)

func GenRandStr(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
