package gangdoupkg

type DataTypeBind interface {
	string | int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64 | bool | complex64 | complex128
}

// 判断数组中是否包含特定元素
func ListContain[T DataTypeBind](items []T, item T) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

// 删除数组中的特定元素
func DelItem[T DataTypeBind](items []T, item T) []T {
	for index, i := range items {
		if i == item {
			items = append(items[:index], items[index+1:]...)
		}
	}
	return items
}

// 数组去重
func Deduplicate[T DataTypeBind](items []T, item T) []T {
	set := make(map[T]struct{}, len(items))
	j := 0
	for _, v := range items {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		items[j] = v
		j++
	}
	return items[:j]
}
