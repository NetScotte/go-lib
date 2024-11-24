package array

// slice去重处理， 返回新的无重复数据的slice，会保持原有顺序
func UniqSlice[T string | int | int32 | int64 | uint | uint32 | uint64](lst []T) []T {
	var reply = make([]T, 0, len(lst))
	var uni = make(map[T]struct{})
	for _, item := range lst {
		_, ok := uni[item]
		if !ok {
			reply = append(reply, item)
			uni[item] = struct{}{}
		}
	}
	return reply
}

func GetMapKeys[T string | int | int32 | int64 | uint | uint32 | uint64, T2 any](m map[T]T2) []T{
	var reply = make([]T, 0, len(m))
	for key := range m {
		reply = append(reply, key)
	}
	return reply
}
