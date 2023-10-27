package sandbox

import (
	"math/rand"
	"time"
)

func BypassSandbox() {
	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, 1024*1024*1024)
	rand.Read(buf)
}
