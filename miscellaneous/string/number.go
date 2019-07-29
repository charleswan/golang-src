package main

import "strconv"

// IsNumeric 检测字符串是否纯数字
func IsNumeric(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}
