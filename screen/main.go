package screen

import (
	"syscall"
)

var (
	user32           = syscall.NewLazyDLL("User32.dll")
	getSystemMetrics = user32.NewProc("GetSystemMetrics")
)

func GetSystemMetrics(nIndex int) int {
	index := uintptr(nIndex)
	ret, _, _ := getSystemMetrics.Call(index)
	return int(ret)
}

const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1
)
