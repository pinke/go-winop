# Window Operations for MS-Windows

##### forked from <a href="https://github.com/koron/go-winop">koron/go-winop</a>

Example to enumerate visible top level windows:

```go
package main

import (
	"fmt"

	"github.com/pinke/go-winop"
)

func main() {
	winop.EnumWindows(func(hwnd winop.Hwnd) bool {
		info := new(winop.WindowInfo)
		if !winop.GetWindowInfo(hwnd, info) {
			return true
		}
		if !winop.IsWindowVisible(hwnd) {
			return true
		}
		if info.Window.IsZero() {
			return true
		}
		fmt.Printf("%x %q %+v\n", hwnd, winop.GetWindowText(hwnd), *info)
		return true
	})
}
```

## more func
*  func enumChildWindows(hwnd Hwnd, fn uintptr, param uintptr) (result bool)
*  func getClassName(hwnd Hwnd, buf []uint16, size int) (length int) 
*  func GetComboBoxInfo(hwnd Hwnd, info *ComboBoxInfo) (result bool)
*  func setWindowText(hwnd Hwnd, text string) (bool) 

