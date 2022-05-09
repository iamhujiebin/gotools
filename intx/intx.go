package intx

import (
	"strconv"
	"strings"
)

// 是否包含int
func IntIndexOf(list []int, target int) int {
	index := -1
	for i, v := range list {
		if v == target {
			index = i
		}
	}
	return index
}

func IntJoin(arr []int, sep string) string {
	var str []string
	for _, v := range arr {
		str = append(str, strconv.Itoa(v))
	}
	return strings.Join(str, sep)
}
