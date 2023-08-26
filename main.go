package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32           = syscall.MustLoadDLL("kernel32.dll")
	procGetConsoleWindow = kernel32.MustFindProc("GetConsoleWindow")
)

func hideConsole() {
	consoleWindow, _, _ := procGetConsoleWindow.Call()
	if consoleWindow != 0 {
		user32 := syscall.MustLoadDLL("user32.dll")
		procShowWindow := user32.MustFindProc("ShowWindow")
		procShowWindow.Call(consoleWindow, SW_HIDE)
	}
}

const (
	SW_HIDE = 0
)

func runInBackground() {
	for {
		fmt.Println("Running in the background...")
	}
}

func main() {
	hideConsole()


  // code here to run in the background
	runInBackground()
}
