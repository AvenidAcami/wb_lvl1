package main

import "fmt"

type Linux struct {
}

func (ro *Linux) DoLinuxStuff() {
	fmt.Println("Doing Linux stuff")
}

type LinuxOperation interface {
	DoLinuxStuff()
}

type Windows struct {
}

func (wind *Windows) DoWindowsStuff() {
	fmt.Println("Doing Windows stuff")
}

type LinuxToWindowsAdapter struct {
	windows *Windows
}

func (ltwa *LinuxToWindowsAdapter) DoLinuxStuff() {
	fmt.Println("Adapter does Linux work on Windows")
	ltwa.windows.DoWindowsStuff()
}

func doSomeLinuxWork(op LinuxOperation) {
	op.DoLinuxStuff()
}

func main() {

	linux := &Linux{}
	windows := &Windows{}

	adapter := &LinuxToWindowsAdapter{
		windows: windows}

	doSomeLinuxWork(linux)
	doSomeLinuxWork(adapter)
}
