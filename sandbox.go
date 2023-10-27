package sandbox

import (
	"math/rand"
	"time"
)

func bypassSandbox() {
	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, 1024*1024*1024)
	rand.Read(buf)
}
