// +build windows

package confirmbox

import (
	"syscall"
	"unsafe"
)

func msgBox(hwnd uintptr, content, title string, flags uint) int {
	uContent, _ := syscall.UTF16PtrFromString(content)
	uTitle, _ := syscall.UTF16PtrFromString(title)
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		hwnd,
		uintptr(unsafe.Pointer(uContent)),
		uintptr(unsafe.Pointer(uTitle)),
		uintptr(flags))
	return int(ret)
}

func isConfirmed(title, content string) bool {
	const (
		null  = 0
		mbYesNO = 0x00000004
		yesResp = 6
	)
	return msgBox(null, content, title, mbYesNO) == yesResp
}