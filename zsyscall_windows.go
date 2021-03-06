// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package winop

import (
	"syscall"
	"unsafe"

	//	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	//moduser32 = windows.NewLazySystemDLL("user32.dll")
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procEnumWindows          = moduser32.NewProc("EnumWindows")
	procGetWindowTextLengthW = moduser32.NewProc("GetWindowTextLengthW")
	procGetWindowTextW       = moduser32.NewProc("GetWindowTextW")
	procGetWindowInfo        = moduser32.NewProc("GetWindowInfo")
	procIsWindowVisible      = moduser32.NewProc("IsWindowVisible")

	procEnumChildWindows     = moduser32.NewProc("EnumChildWindows")
	procSetWindowTextW       = moduser32.NewProc("SetWindowTextW")
	procGetClassNameW        = moduser32.NewProc("GetClassNameW")
	procGetComboBoxInfo      = moduser32.NewProc("GetComboBoxInfo")
)

func enumWindows(fn uintptr, param uintptr) (result bool) {
	r0, _, _ := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(fn), uintptr(param), 0)
	result = r0 != 0
	return
}

func GetWindowTextLength(hwnd Hwnd) (length int) {
	r0, _, _ := syscall.Syscall(procGetWindowTextLengthW.Addr(), 1, uintptr(hwnd), 0, 0)
	length = int(r0)
	return
}
func getWindowText(hwnd Hwnd, buf []uint16, size int) (length int) {
	var _p0 *uint16
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, _ := syscall.Syscall6(procGetWindowTextW.Addr(), 4, uintptr(hwnd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(size), 0, 0)
	length = int(r0)
	return
}
func GetWindowInfo(hwnd Hwnd, info *WindowInfo) (result bool) {
	r0, _, _ := syscall.Syscall(procGetWindowInfo.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(info)), 0)
	result = r0 != 0
	return
}

func enumChildWindows(hwnd Hwnd, fn uintptr, param uintptr) (result bool) {
	r0, _, _ := procEnumChildWindows.Call(uintptr(hwnd), uintptr(fn), uintptr(param), 0)
	result = r0 != 0
	return
}

func getClassName(hwnd Hwnd, buf []uint16, size int) (length int) {
	var _p0 *uint16
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, _ := syscall.Syscall6(procGetClassNameW.Addr(), 4, uintptr(hwnd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(size), 0, 0)
	length = int(r0)
	return
}

func GetComboBoxInfo(hwnd Hwnd, info *ComboBoxInfo) (result bool) {
	r0, _, _ := syscall.Syscall(procGetComboBoxInfo.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(info)), 0)
	result = r0 != 0
	return
}
func setWindowText(hwnd Hwnd, text string) (bool) {
	r0, _, _ := procSetWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(syscall.StringBytePtr(text))))
	return r0 != 0
}


func IsWindowVisible(hwnd Hwnd) (result bool) {
	r0, _, _ := syscall.Syscall(procIsWindowVisible.Addr(), 1, uintptr(hwnd), 0, 0)
	result = r0 != 0
	return
}
