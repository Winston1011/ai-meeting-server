package utils

/**
 * @description: 数组去重RemoveDuplicate
 * @return {*}
 */
func RemoveDuplicate[T comparable](arr []T) []T {
	m := make(map[T]bool)
	result := []T{}
	for _, item := range arr {
		if _, ok := m[item]; !ok {
			m[item] = true
			result = append(result, item)
		}
	}
	return result
}

/**
 * @description: 数组反转
 * @return {*}
 */
func Reverse[T any](arr []T) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
