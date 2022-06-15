package main 

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("WHAT!?")
	}
}
func run() {
	// /proc/self/exe refers to this executable, 
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	//cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	cmd.SysProcAttr = &syscall.SysProcAttr{
		//NEWUTS: Unix Timesharing System -> isolation around the hostname
		//NEWPID: Giving container new PIDs starting from scratch 
		//NEWUSER: Using a new user namespace for rootless container
		//NEWNS: A new namespace for mounts
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS, 
		// User id and group id, inside the container I want to be root 
		// root inside the container corresponds to this user id (the container former)
		// containers view of my user and group id will be different than the hosts.
		Credential: &syscall.Credential{Uid:0, Gid:0},
		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getuid(), Size: 1},
		},
		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getgid(), Size: 1},
		},


		//Unshareflags: syscall.CLONE_NEWNS, 
	}
	// We cannot set the hostname here with the following
	// must(syscall.Sethostname([]byte("container")))
	// because the command will be run after, the process is not created yet 
	//namespaces are not established yet. So if I do this, I will effect the host.
	must(cmd.Run())

}

func child() {
	fmt.Printf("[%v]", os.Getpid())
	
	c_group()
	
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// I no longer need to setup the namespaces anymore, we have already done that in 'run'	
	 must(syscall.Sethostname([]byte("container")))
	
	//(1) TO DO: We need a ubuntu file system for this to work
	//must(syscall.Chroot("~/media/psf/ToUbuntu/ubuntu_file_system"))
	// now we are in an undefined location, that is why we are changin to root directory
	//must(syscall.Chdir("/"))
	// then we should mount the /proc
	//must(syscall.Mount("proc", "proc", "proc",0, ""))
	// to tidy up afterwards 
	//defer must(syscall.Unmount("proc", 0))
	// (1) END

	must(cmd.Run())

}

func c_group() {
	cgroups := "/sys/fs/cgroup"
	pids := filepath.Join(cgroups, "pids")
	u_name := "container"

	must(os.Mkdir(filepath.Join(pids, u_name), 0755))
	must(ioutil.WriteFile(filepath.Join(pids, fmt.Sprintf("%s/pids.max", u_name)), []byte("20"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, fmt.Sprintf("%s/notify_on_release", u_name)), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, fmt.Sprintf("%s/cgroup.procs", u_name)), []byte(strconv.Itoa(os.Getpid())), 0700))


}

func must(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(3)
	}
}
