package main

import (
	"fmt"
	"syscall"
	"time"
)

var (
	user32 = syscall.MustLoadDLL("user32.dll")
	blockinput = user32.MustFindProc("BlockInput")
)
type(
	BOOL            int32
)

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}

func blockInput(fBlockIt bool) bool{
	ret, _, err := blockinput.Call(uintptr(BoolToBOOL(fBlockIt)))
	fmt.Println(err)
	return ret != 0
}

func looper() {
	for{
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	go looper()
fmt.Println("Enabling...")
blockInput(true)
fmt.Println("Enabled try typing or using mouse.")
time.Sleep(10 * time.Second)

fmt.Println("Disabling...")
blockInput(false)
}
