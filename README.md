# Container
Container from scratch in go

We expand on the Talk of Liz Rice and her sequel on 'Containers from scratch'

- Namespaces
- Rootless containers
- Cgroups
- Attacks like : 
	- Fork bomb (this is what Liz tries after she establishes the CGroups)
	- child command is unsecure

### Namespaces
- If you `go run main.go run /bin/bash` you will be running a shell inside it. With the `syscall.CLONE_NEWUTS` flag, hostname namespace is isolated. After you run the shell, if you type 'hostname' you will see your Ubuntu hostname. You can change it like `hostname container` making its hostname 'container' now. Check it with `hostname`, and you will see it is 'container' now. BUT if you check it from your terminal, you will see that your hostname is still the same, your Ubuntu hostname. 		  


- We will transform this sample program into a strongly isolated container.

It will be run directly like `docker run`. 

The container gets the other process information from /proc. 
If you `ps aux` you can see other processes running on the host. /proc is a pseudo-file system that
OS uses to (kernel \<-\> user space) communicate information about processes. `ps` goes here to find out the information about running processes.
So we need our container to have its own /proc. We need to give the container its own filesystem.
Then the container should chroot (changing the apparent root directory for the current running process and its children). The container will not be able to name files outside of its designated directory. 

### Rootless Container
This is what you need when you do not have root privileges in the host machine. As a non-privileged user I am not allowed to form namespaces (syscall Cloneflags). With `syscall.CLONE_NEWUSER`, we will be able to use a new user namespace. As long as we don't mount anything we can create the container and namespaces with the `syscall.CLONE_NEWUSER` and new credential and mappings.

### CGroups
CGroups might not be working with rootless containers. As namespaces can be sumarized as what you process 'see', Cgroups are in that manner what the process can 'use'. It uses a pseudo-file system too. 

`mount | grep cgroups`

![An example of cgroups in Ubuntu with](images/cgroups_ss.png)

`:() { : | : & }; :` initiates a fork bomb. In a constraint container, it will not be able to find unlimited resources. When you check the processes you will see \<defunct\> ones which are the fork bombs, 
To stop the bomb, exit the container. 
