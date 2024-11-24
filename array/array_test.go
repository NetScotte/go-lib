package array

import "testing"

func TestUniSlice(t *testing.T) {
	var s = []string{"aa", "bb", "aa", "cc", "bb"}
	t.Log(UniqSlice(s))
}
