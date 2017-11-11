# Window Operations for MS-Windows

Example to enumerate visible top level windows:

```go
package main

import (
	"fmt"

	"github.com/koron/go-winop"
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
