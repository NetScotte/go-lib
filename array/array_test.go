package array

import "testing"

func TestUniSlice(t *testing.T) {
	var s = []string{"aa", "bb", "aa", "cc", "bb"}
	t.Log(UniqSlice(s))
}

func GetMapKeys[T string | int | int32 | int64 | uint | uint32 | uint64, T2 any](m map[T]T2) []T{
	var reply = make([]T, 0, len(m))
	for key := range m {
		reply = append(reply, key)
	}
	return reply
}
