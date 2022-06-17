## Short Term
- [] For TDD, first there should be tests that cover the current version. 
- [] After refactoring, eliminating defects or unsecure parts of the current state, each step and task should be planned.
 
## Image Production
- [] Creating an image from scratch that complies with [OCI Image Format Specification](https://github.com/opencontainers/image-spec/blob/main/spec.md) [v 1.0.2 release](https://opencontainers.org/release-notices/v1-0-2-image-spec/)
    - [] Try to make it architecture agnostic (arm64, amd64 x86)
    - [] Only can work with Linux since `syscall.Cloneflags` (namespaces, cgroups) works only with Linux 
    - [] config
    - [] image index
    - [] image manifest
    - [] layers
    - [] should be docker runnable `docker run <my image>` should work.

## Containerization 
- [] By using our custom built image, we will establish its namespaces, cgroups, volumes, and networking. 
    - [] If possible, should be somehow rootless. (*cgroups can be problematic here*)
    - [] Should be secure, on top of the fork bomb, every cgroup restriction should be tested.
    - [] Namespaces should be tested too.
    - [] The container should be able to gather its dependencies 
    - [] The container should be able to run a realistic process. (In our case, our container will run a server process)
    - [] this will be the part where instead of `docker run` we use our program `<my container> run <cmd> [args]` or with a dockerfile-like file we can specify the entry point ourselves 

## Creating a very in depth documentation as a resource for Images and Containers
- [] Every concept should be desribed.
- [] Every resource should be presented in an appropriate format.
- [] While learning the docker phases, like a surgery, I plan to dissect every step that docker takes. 
    - [] I should put all the tools and their resources as well. 
    - [] I plan to put all the images, screenshots of what I do in these phases.