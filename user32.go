package winop

import "syscall"

type EnumWindowsProc func(hwnd Hwnd) bool

func EnumWindows(fn EnumWindowsProc) bool {
	return enumWindows(syscall.NewCallback(func(hwnd Hwnd, data uintptr) uintptr {
		if !fn(hwnd) {
			return uintptr(0)
		}
		return uintptr(1)
	}), uintptr(0))
}

func GetWindowText(hwnd Hwnd) string {
	l := GetWindowTextLength(hwnd) + 1
	buf := make([]uint16, l)
	getWindowText(hwnd, buf, l)
	return syscall.UTF16ToString(buf)
}
