# Deep Diving 

Before starting, I want to clarify some of my notations. I will abide by the conventions of Linux man pages. In those:

`[ ]` means that anything within (any values, information, or parameters) is optional. You can choose one, more than one or none at all.

`< >` means that the element within (again any parameter, information, or a value) is mandatory. We require replacing the text in between angle brackets with actual appropriate information.

`-` means the options. It does not stand by itself, it will be followed by some form of a value ( or information or parameter).

`--` We will not use that one I guess. Anyway while we are on it, a little extra does not kill anyone ... They are explicitly written as their names. It is pretty common and you may stumble upon them anywhere. For example, an option `-l` indicating an option as a 'list', will be given `--list` as a long option as well. Both accomplish the same thing. When you use heavily on options, it helps to use long options for visibility and readability for you and others.

`{ }` I am not sure if this is common. This means within are required options that can be in any relation such as mutual exclusion and so on. One of the specified options or some in conjunction with others *must* be present. For example, filtring some results with `{ --time $TODAY | --priority 9 }` means "I want the ones which are from today (assume that entries have times associated with them) **OR** those who are of priority 9". Note that one of these conditions must be fulfilled. If none of them is satisfied, the returned result is empty.


# Container
## A subtle deblurring of the concepts

### What is a Container?

Containers are structured files a.k.a. file system layers and configuration files as JSON blobs resulting with a packaging scheme and runtime isolation. This is a lot to take in. Let us divide and conquer this.

![file system meme...](/images/fs_meme.png)
---
**What do structured files or a file system mean?**

![file system meme...](/images/fs_meme_2.jpeg)


Information is contained within a file. A directory itself is a file too. The only difference is that it can contain other files/directories. Directories have children in a tabular form while their content is null. Thus their size is 0.

if we `mdls images/` in this repository, we will see the following. `mdls` stands for Metadata List, and as its name suggests it is for listing the metadata of the specified **file**. I drew an arrow to indicate that their starting points are actually corresponding. They are not aligned so to reflect otherwise. Let us avoid confusion if there is any. Note that they are different in the <span style="color:orange">orange box</span>, as we mentioned a directory do not have a size. The markdown only has additional fields as in the <span style="color:blue">blue box</span>.

![metadata comparison](/images/metadata_file_and_dir.png)


You can see that they are not awfully different. It is trivial to see that a directory is a file itself that can contain multiple other files.

We can say that containers are a collection of Linux technologies that run one or more Linux processes. That is why we will discover 'structured files/file systems' in Linux itself. 

If I recklessly type `tree /` (
`/` *refers to the root directory. The root directory is - as its name suggests - the directory from which all the other directories branch off.* 
), my life will flash before my eyes. 

https://user-images.githubusercontent.com/40325021/174617365-eb663412-afc4-4f77-80e2-62787f7bfff6.mov

We will not cover all of this obviously ... Let us limit the `tree` command by telling it to show only the 1st children of the root directory by `tree -L 1 /`.

![First Degree Childs of the root directory](/images/from_the_root_of_the_ubuntu_fs.png)

---
**`/bin`** includes the binaries - programs that you can run - which are essential user commands. `cat`, `mv`, `kill`, `ls`, `rm` and so on all of them are within **`/bin`**. There can be other **`/bin`** directories within the file system.

You might be wondering why there is an arrow `->` after **`/bin`**. `tree` shows symbolic links in that format: `name -> real path. A symbolic link - you may have encountered it referred to as a symlink, shell link, or a soft link - points to another file or directory on the system or a connected file system. A fancy way of saying a `shortcut`. So my **`/bin`** directory actually points to **`usr/bin`** and the binaries that 'I' - as a user - use and run are in that folder.

You see where those programs or commands reside with `which`. For example, the command to move stuff around `mv`:

![moves place](/images/moves_place.png)

I would like to `ls` into **`/bin`** but it is huge, so you would have to do it yourself...

---

![lsa_boot](/images/lsa_boot.png)

**`/boot`** includes what your system needs when starting (a.k.a. booting). There is information for GRUB (**GR**and **U**nified **B**ootloader) - the boot loader package - and all the kernels - we will dive deeper into kernels, but for now, it is basically the core program of OS that has complete control over everything in the system (you have installed, their configuration files and so on) -. 

**`/boot`** is a very important directory, I don’t recommend touching any of these files...

---


**`/cdrom`** I am not entirely sure about this one. If you would check for yourself you may not see it. Normally, when you want to use CD(stands for compact disk) you must first mount it on an empty directory. Mounting means connecting and rendering another file system accessible at a certain point or node - such as an empty directory solely existing for this purpose-  in the Linux directory tree. We must do this because Unix-based systems have only one directory tree. Everything that is to be accessible should be somewhat connected to that tree from somewhere. It is not weird actually, since you are within this directory tree all along. You cannot go out of here. I think other systems such as Windows have other forms and numbers of trees but it is not important for now. Just note that the empty directory to mount to can be anywhere and have any name. You can mount a CD-ROM(compact disk read-only memory) to a directory named `USB` with the `mount` command.  It does not matter. So this directory probably is here because of compatibility reasons. There can be a `/floppy` directory too for the same reasons as `/cdrom`. Note that now, when you attach a USB to your system, it will be automatically mounted.

---

![lsa_boot](/images/lsa_dev.png)

**`/dev`** includes device files, and all your hardware components on your system. **`/dev`** must contain a command named `MAKEDEV`, which can create devices as needed. When you plug in a webcam, a new device entry will be created here.

---
This is what inside **`/etc`** looks like, it is huge so I could not fit all in here…

![inside etc ](/images/inside_etc.png)

**`/etc`** was your special drawer where you everything that you are not sure where to put just like “et cetera” I guess. It does not matter now because now it is "Everything to Configure". It is pronounced as 'etsy'. Contains most of the system-wide configuration files. (A "configuration file" is a local file used to control the operation of a program and it must be static and cannot be an executable binary.) Every name on the system ( name of the system itself, machines on the network ), users and their passwords, and when and where the partitions on hard disks to be mounted are all in here. Other examples can be SSH, Pipewire, systemD, and Firefox, which have all configuration files here. 

---

Here is all my personal stuff 😱

![dirty laundry ](/images/lsa_home.png)

**`/home`** ... where are all your personal belongings? At your home ... In **`/home`**, personal directories of the user reside. So in my `/home` directory, there would be a directory named `cem` which will contain all my directories. I am given permission to do whatever I want within this directory.

---

Here is my library 🤓📚 

![library ](/images/lsa_lib.png)

**`/lib`** contains all the code that applications use a.k.a. libraries. I think the closer a directory is to the root directory, the more important and special it gets. In this sense, **`/lib`** contains kernel modules that are crucial for the system. These shared library images(we will discover deeply what an image is in the next sections) are needed to boot the system and run the commands in the root filesystem, ie. by binaries in **`/bin`** and **`/sbin`**. In short, everything - such as WiFi, and your video card - works because of those modules. Again, you may stumble upon other **`/lib`** folders. Shared libraries that are only necessary for binaries in **`/usr`** must not be in **`/lib`**. Only the shared libraries required to run binaries in **`/bin`** and **`/sbin`** should be here. 

---

**`/lost+found`** is as its name suggests for data fragments there are not referenced anywhere anymore. Let us step back for a second and see it clearly. When your system is not properly shut down, due to a crash or a failure,at the next boot a diligent file system check using `fsck` will take place. Wait, what the hell is `fsck`? `fsck` is a system utility tool that helps detect inconsistencies of a file system and interactively repair if there are any in Unix-based systems, it stands for `Filesystem check`. In my mac, if I `man fsck`, it hints to me that the script `/etc/rc` uses `fsck` during automatic reboot for example. `fsck` scans your filesystem and performs some checks to ensure integrity. It will try to recover any corrupt files that it finds. Corrupted files are altered files - this can happen in various ways, do not think of only naive ways, even a virus can do that - such that the file is not readable to the hardware or indecipherable to the software anymore because its bits are rearranged. When such a file is encountered, it will be placed in **`/lost+found`** because even though they do not seem to be helpful anymore, there can be a chance that something worthwhile may be recovered. It can find a complete file that does not have a name on the file system which will render it inaccessible in any way. There is an interesting bit. If you happened to delete **`/lost+found`**, do not `mkdir lost+found`. Instead, there is a special command for that: `mklost+found`. This is because **`/lost+found`** is a special directory. It preallocates space beforehand for `fsck` to be able to place files. This is done so that when `fsck` runs, it does not have to allocate blocks in the file system during recovery.

---


**`/media`** is the point in the directory tree where external storage is automatically mounted when attached. This automatic mounting by detecting newly inserted storage is relatively new, which makes this directory fairly new too. To have separate mounting points for each removable media at the root directory would crowd the place, so gathering them in here tidies up mounting directories. `floppy`(floppy drive), `cdrom`(CD-ROM drive), `cdrecorder`(CD Writer), `zip`(Zip drive) should be in the **`/media`**.    

---

**`/mnt`** is a geezer one too. It was used to manually mount storage, and temporarily mount a filesystem as needed. This directory does not affect any execution or run of any program. If I am to farm in Supspace network for example, to increase my memory and to gain more tickets, I add an SSD to my system, I can set up a permanent mount point for it here.

---

**`/opt`** is reserved for add-on application software packages. Add-on means software that is not part of your system, any third party, or your custom software.
That software is not installed from distribution repositories. A distribution repository is a server hosting specific **verified** programs for some Linux OS. This is where your `apt upgrade` or such, your system pulls the packages from those repositories. You can find which ones your Linux system uses in ` /etc/apt/sources.list`. If I `cat` that, and cleanse the output from comments, we end up with the following : 

```bash
deb http://ports.ubuntu.com/ubuntu-ports/ focal main restricted
deb http://ports.ubuntu.com/ubuntu-ports/ focal-updates main restricted
deb http://ports.ubuntu.com/ubuntu-ports/ focal universe
deb http://ports.ubuntu.com/ubuntu-ports/ focal-updates universe
deb http://ports.ubuntu.com/ubuntu-ports/ focal multiverse
deb http://ports.ubuntu.com/ubuntu-ports/ focal-updates multiverse
deb http://ports.ubuntu.com/ubuntu-ports/ focal-backports main restricted universe multiverse
deb http://ports.ubuntu.com/ubuntu-ports/ focal-security main restricted
deb http://ports.ubuntu.com/ubuntu-ports/ focal-security universe
deb http://ports.ubuntu.com/ubuntu-ports/ focal-security multiverse
```

I had to mingle with these for a reason back then, but I do not remember when and why ... There is another file `sources.list.d` for any other files ending with `.list`, and they will be included in `sources.list` I guess. 


Without straying from the main point anymore, let us go back to **`/opt`**. When you compile software that you have built, it can land here. **`/opt`** will divide itself structurally such as an application a library will be placed **`/opt/bin`** and **`/opt/lib`** respectively. When a package is installed in **`/opt`**, its static files should be in **`/opt/<package>`**. An important distinction is that there are other places for such software to be placed such as  **`/usr/local/bin`** and **`/usr/local/lib`**. This behavior is dictated by how the developers have configured files that control the compilation and installing process.

---

**`/proc`** contains information about your computer such as your CPU or the kernel. It is a virtual directory - it does not have an actual physical component, every physical disk is treated as a file in Linux (even the disk that you are working off of). that is why it is called virtual -, a pseudo-file system that the OS uses to (kernel \<-\> user space) communicate information about computer such as processes. The content is generated on the fly. When you `ps` for example, the OS gathers the process information from that directory. Below we see what is inside, note that the directories with numbers as their names are processes running and those numbers are their IDs or `pid`s.

![listing the /proc directory content](/images/proc_dir.png)

In this sense, we did not specify that but you may notice that **`/dev`** has the same feel to it as **`/proc`** too. And yes, it is a virtual directory too. **`/proc`** is a vital directory for us because we will get in here and establish our own cgroups and more. We will come to them later with utmost detail.

---

Let's check if this 'super user' guy is really super or a lamer with no life... 🧐

![ls root](/images/lsa_root.png)

**`/root`** is equivalent to `root` - a.k.a. administrator or super - users **`/home`**. 

---

![run lsa](/images/run_ls.png)

**`/run`** is the run-time variable data. It is a temporary location for system processes to store their own temporary data. When you boot your system, this folder is probably empty and starts to get populated as your system operates.  It is relatively a new directory. Like **`/proc`**, it stores run-time volatile data.

---

This one is pretty huge too...

**`/sbin`** has the same duty as the **`/bin`** directory but it is for the superuser - 's' in the beginning for that reason -. You use those programs with the `sudo` command which **temporarily** renders you "super", and gives you powers of the superuser.
As you may infer, **`/sbin`** contains mutators, and programs that change something within the system such as installing, deleting, or formatting/changing a component or package. The essential binaries to boot, restore, recover, and repair the system are here too. For example, updating passwords for users `chgpasswd [group_name:password]` resides here. 

![chgpasswd place](/images/chgpasswd_place.png)

---

![lsa snap ](/images/lsa_snap.png)

**`/snap`** is for storing your favorite Snapchat stories. I am kidding. Let us read what the `README` says within the **`/snap`** directory. 

![snap readme](/images/snap_readme.png)

Let me refresh your memory: `apt` was a cli handling installation or removing the packages. It is a packaging tool. 'snap' - not the directory, the program - is similar, it is a packaging and deployment system. Snaps are a secure and scalable way to embed applications on Linux devices. A snap is an application **containerized** with all its dependencies. A snap can be installed using a single command on any device running Linux. Applications run fully isolated in their own sandbox, thus minimizing security risks. Though, there was an incident, where in a snap, the developer put cryptocurrency mining software and all hell broke loose...

Snaps are self-contained applications running in a sandbox with mediated access to the host system. Self-contained means they consist of a single, installable bundle containing the application and a copy of the run-time environment needed to run the application. When the application is installed, it behaves in the same way as any native application. The file format for snap is a single compressed filesystem with the extension `.snap`. This filesystem contains the application, libraries it depends on, and declarative metadata. This metadata is interpreted by snapd to set up an appropriately shaped secure sandbox for that application. After installation, the snap is mounted - this is mediated access we talked about above - by the host operating system and decompressed on the fly when the files are used. Although this has the advantage that snaps use less disk space, it also means some large applications start more slowly.

Those `.snap` files actually are in `var/lib/snapd/` directory. When they are run, as we said earlier they are mounted to **`/snap`**.
![var/lib/snapd](/images/var_lib_snapd.png)

If we look at **`/snap/core<number>/current`**, there is a filesystem just like the Linux filesystem!

![var/lib/snapd](/images/snap_core.png)

`snapd` is the daemon required to run those snaps. It downloads the snap from the store, mounts it somewhere in **`/snap`**, confines it, and runs apps out of it. The sandbox that those snaps run in is called 'Snapsandbox'.

To sum it up, this directory contains the mount-points for your snaps and several symlinks which are needed by snapd.

You may feel exhausted now. Why do I keep straying from the main thing, the Linux filesystem? I try to provide every bit of information so that confusion or shaky ground is avoided. I hope you did not get confused. Let us continue.

---

**`/srv`** contains data for the systems provided services. This includes your servers. Let us say we are running an FTP server
on our Linux machine. All the files that I am serving would go in **`/srv/ftp`**. Any service will have its data in **`/srv/<service>`**. Another example can be a web service. It would have its sites HTML files in **`/srv/http`** or **`/srv/www`**

---

**`/swap.img`** is a swap image file. It is an HDD(Hard Disk Drive) -  a non-volatile data storage device - partitioned space. In the event you run out of RAM, some of the unused data currently in memory can be "swapped out" to make room for what needs to run right now. You can use swap images for other reasons. As you may know, RAM is volatile. It is a temporary allocation of your data. You can save data in your system through a swap file since it is a partition on HDD - persistent memory -. Then, when you reboot your computer, it will transfer that data to RAM.

---

![lsa sys](/images/lsa_sys.png)

**`/sys`** contains information from devices connected to your computer. It is another virtual directory like **`/proc`** and **`/dev`**. You can adjust the brightness of your screen or volume of your speakers from the **`/sys`** directory if you know what you are doing and can locate the appropriate files...

---

![lsa tmp](/images/lsa_tmp.png)

**`/tmp`** is for programs that require temporary files. One should not assume that files in this directory persist. You can utilize **`/tmp`** for personal needs too without hesitation because you can interact with it without being the administrator. 

---

![lsa usr](/images/lsa_usr.png)

**`/usr`** was your old home. Users **`/home`** folder was in this directory in the past. Now contains a bundle of directories such as libraries, applications, icons, wallpapers, and much more. They are shared by applications and services. Note that the symlinks **`/bin`**, **`/lib`**, and **`/sbin`** point to their counterparts with the same name within this directory.

---

![lsa var](/images/lsa_var.png)

**`/var`** is very similar to **`/tmp`** directory. Although it is named after its contents to be variable, changing frequently, it generally holds information longer than **`/tmp`**. It may also store log files and KVM Virtual Machine disk images. If an outsider intruder tries to break into your system, your system’s firewall will log this attempt here. It may also contain **spools** for later tasks. **`/var/spool`** will include the data of those tasks which are awaiting later processing. A job sent to a busy shared printer, or a mail that is waiting to be delivered can be examples of those tasks.

---



<!-- TO DO BELOW 
### What is an Image?

### What are namespaces?
### What are cgroups?


## Containers from scratch in Go. 

To be precise; 


1. Image creation without Docker compliant with [OCI Image Format Specification](https://github.com/opencontainers/image-spec/blob/main/spec.md) [v 1.0.2 release](https://opencontainers.org/release-notices/v1-0-2-image-spec/). This broadly means, any  container runtime can `run`the images that we created, any registry should accept our images when we register.
2. Containerization with the image that we built, establishing : namespaces, cgroups, volumes, networking and if we can security. We will discover how far we can go when we try them rootless.

### We expand on the Talk of Liz Rice and her sequel on 'Containers from scratch'
- The first security test is introduced here : 
    - Fork bomb (this is what Liz tries after she establishes the CGroups)
    - child command is unsecure though

### Namespaces
- If you `go run main.go run /bin/bash` you will be running a shell inside it. With the `syscall.CLONE_NEWUTS` flag, hostname namespace is isolated. After you run the shell, if you type `hostname` you will see your Ubuntu hostname. You can change it like `hostname container` making its hostname 'container'. Check it with `hostname`, and you will see it is changed to 'container'. **BUT** if you check it from your terminal as you as the user, you will see that your hostname is still the same, your Ubuntu hostname.         


The container gets the other process information from `/proc`. 
If you `ps aux` you can see other processes running on the host. `/proc` is a pseudo-file system that the OS uses to (kernel \<-\> user space) communicate information about processes. `ps` goes here to find out the information about running processes.
So we need our container to have its own `/proc`. We need to give the container its own filesystem. Then the container should `chroot` (*changing the apparent root directory for the current running process and its children*). The container will not be able to name files outside of its designated directory. 

### Rootless Container

This is what you need when you do not have root privileges in the host machine. As a non-privileged user I am not allowed to form namespaces (syscall Cloneflags). With `syscall.CLONE_NEWUSER`, we will be able to use a new user namespace. As long as we don't mount anything - we are not allowed to mount when we are not *root* - we can create the container and namespaces with the `syscall.CLONE_NEWUSER` and new credential and mappings.

### CGroups
As namespaces can be summarized as what process can 'see', Cgroups are in that manner what the process can 'use'. It uses a pseudo-file system too. 
*Note: CGroups might not be working with rootless containers.*

Let's check cgroups in our system.
`mount | grep cgroups`

![An example of cgroups in Ubuntu with](images/cgroup_ss.png)


### Security
**Fork Bomb a.k.a. Rabbit**
It is a Denial-of-Service attack. You basically clone a function again and again until you consume all the system resources of the host. 

`:() { : | : & }; :` initiates a fork bomb. To be verbose, it means the following: 

`:()`define a function called ':' (*note that you can define it whichever name you want, ':' is concise.) 
`{` function body start
`: | :` call  the ':' recursively and pipe it to a new ':'. (*Piping basically is connecting the stdout of the current command/process to the stdin of the next command/process.*)
`&`run in the background
`};` function body and definition end
`:` call the colon, initiate the fork bomb.


 In a constraint container, it will not be able to find unlimited resources. When you check the processes you will see \<defunct\> ones which are the fork bombs, 
To stop the bomb, exit the container. 
-->

<details>
<summary><b>References</b></summary>

[mdls](https://ss64.com/osx/mdls.html)

[Linux Filesystem Hierarchy Standard](https://refspecs.linuxfoundation.org/FHS_3.0/fhs/index.html)

[Classic SysAdmin: The Linux Filesystem Explained](https://linuxfoundation.org/blog/classic-sysadmin-the-linux-filesystem-explained/)

[Understanding the Linux Virtual Directory Structure](https://www.maketecheasier.com/linux-virtual-directory-structure/)

[What is the purpose of the /cdrom directory in Ubuntu-based Linux](https://superuser.com/questions/1020383/what-is-the-purpose-of-the-cdrom-directory-in-ubuntu-based-linux)

[CD-ROMs](https://dsl.org/cookbook/cookbook_31.html#SEC426)

[Workflow Symbolic Links](https://www.digitalocean.com/community/tutorials/workflow-symbolic-links)

[lost+found](https://tldp.org/LDP/Linux-Filesystem-Hierarchy/html/lostfound.html)

[purpose of the lost+found](https://unix.stackexchange.com/questions/18154/what-is-the-purpose-of-the-lostfound-folder-in-linux-and-unix)

[mklost+found](
https://man7.org/linux/man-pages/man8/mklost+found.8.html
)

[Sources List](https://wiki.debian.org/SourcesList)

[Listing Repos on Linux](https://www.networkworld.com/article/3305810/how-to-list-repositories-on-linux.html)

[sources.list linux man page](https://linux.die.net/man/5/sources.list#:~:text=The%20package%20resource%20list%20is,%2Fetc%2Fapt%2Fsources.)

[virtual directory structure](https://www.maketecheasier.com/linux-virtual-directory-structure/)

[Virtual Files and Directories Under /proc](
https://docs.oracle.com/cd/E37670_01/E41138/html/ch04s02s01.html
)

[chgpasswd](https://man7.org/linux/man-pages/man8/chgpasswd.8.html)

[Managing Ubuntu Snaps](https://hackernoon.com/managing-ubuntu-snaps-the-stuff-no-one-tells-you-625dfbe4b26c)

[Deploying Self-Contained Applications](https://docs.oracle.com/javase/tutorial/deployment/selfContainedApps/index.html)


[Deploying Self-Contained Applications](https://docs.oracle.com/javase/tutorial/deployment/selfContainedApps/index.html)

[Self-Contained Packaging](https://docs.oracle.com/javase/8/docs/technotes/guides/deploy/self-contained-packaging.html)

[Snaps intro](https://ubuntu.com/core/services/guide/snaps-intro)

[Snap on ask ubuntu](https://askubuntu.com/questions/963404/what-do-snap-snapd-and-snappy-refer-to)

[swap file creation and removal in Ubuntu](https://docs.rackspace.com/support/how-to/create-remove-swap-file-in-ubuntu/)

[linode swap image](https://www.linode.com/community/questions/17264/whats-a-swap-image-and-what-is-for)

</details>

