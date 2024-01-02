package utils

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	arr := []string{"1", "2", "3"}
	Reverse(arr)
	fmt.Println(arr)
}
