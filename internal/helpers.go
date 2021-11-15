package internal

import "time"

func GenerateID() uint64 {
	return uint64(time.Now().UnixNano() / (1 << 22))
}
