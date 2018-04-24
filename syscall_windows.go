package winop

import "syscall"

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go syscall_windows.go

//sys enumWindows(fn uintptr, param uintptr) (result bool) = user32.EnumWindows
//sys GetWindowTextLength(hwnd Hwnd) (length int) = user32.GetWindowTextLengthW
//sys getWindowText(hwnd Hwnd, buf []uint16, size int) (length int) = user32.GetWindowTextW
//sys GetWindowInfo(hwnd Hwnd, info *WindowInfo) (result bool) = user32.GetWindowInfo
//sys IsWindowVisible(hwnd Hwnd) (result bool) = user32.IsWindowVisible

type Hwnd syscall.Handle
type Dword uint32
type Word uint16
type Long int32
type Uint uint32
type Atom Word

type Rect struct {
	Left   Long
	Top    Long
	Right  Long
	Bottom Long
}

func (rc Rect) IsZero() bool {
	return rc.Left == 0 && rc.Top == 0 && rc.Right == 0 && rc.Bottom == 0
}

type WindowInfo struct {
	Size            Dword
	Window          Rect
	Client          Rect
	Style           Dword
	ExStyle         Dword
	WindowStatus    Dword
	CxWindowBorders Uint
	CyWindowBorders Uint
	WindowType      Atom
	CreatorVersion  Word
}

type ComboBoxInfo struct {
	cbSize      Dword
	rcItem      Rect
	rcButton    Rect
	stateButton Dword
	hwndCombo   Hwnd
	hwndItem    Hwnd
	hwndList    Hwnd
}
