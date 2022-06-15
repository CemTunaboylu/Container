package main 

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("WHAT!?")
	}
}
func run() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	cmd.SysProcAttr = &syscall.SysProcAttr{
		//NEWUTS: Unix Timesharing System -> isolation around the hostname
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	must(cmd.Run())

}
func must(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(3)
	}
}
