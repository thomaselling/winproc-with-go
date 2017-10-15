package main

import (
		"fmt"
		"syscall"
		//"unsafe"
)

const (
		ALL_PROC = 0x00000002
		MAX_PATH = 260
)

type ProcessEntry32 struct {
  	Size            uint32
  	Usage           uint32
  	ProcessID       uint32
  	DefaultHeapID   uintptr
  	ModuleID        uint32
  	Threads         uint32
  	ParentProcessID uint32
  	PriClassBase    int32
  	Flags           uint32
  	ExeFile         [MAX_PATH]uint16
  }

func Getallproc() (syscall.Handle, error) {

	//https://msdn.microsoft.com/en-us/library/windows/desktop/ms682489(v=vs.85).aspx
	//0x00000002 includes all processes
	//0 includes all processes
	//returns open handle on success, error on failure
	openhandle, err := syscall.CreateToolhelp32Snapshot(ALL_PROC, 0)
	//inside a function, the := short assignment statement with implicit type

	//print err details
	if err != nil {
		syscall.GetLastError()
		fmt.Printf("%d, %d\n", openhandle, err)
  		return openhandle, err
  	}
	//ProcessEntry32
	fmt.Printf("%d, %d\n", openhandle, err)
	return openhandle, err
}

func main() {
	Getallproc()
}