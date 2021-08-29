package tt

import (
	"fmt"
	"math/rand"
	"time"
)

func GenRandStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
