package internal

import "time"

func GenerateID() uint64 {
	return uint64(time.Now().UnixNano() / (1 << 22))
}

func Contains(array []string, value string) bool {
	for _, str := range array {
		if str == value {
			return true
		}
	}
	return false
}
