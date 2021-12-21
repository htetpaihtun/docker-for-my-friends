#### Now let's take a closer look at Docker components. 

Before we start to look at Docker components. Let's learn about what actually is a container.

### 3.1 What is a Linux container 

As I mentioned before, in earlier years, Docker directly uses [Linux containers (LXC)](https://linuxcontainers.org/) as its runtime and then later dropped it.        
 So, let's take a look it LXC before.

![LXC architecture](https://www.baeldung.com/wp-content/uploads/sites/2/2020/11/Linux-Container-Architecture-1.jpg)

  _Figure 3.1.2 LXC architecture_

A Linux container, at its core, is made up of; 
- Namespaces
- Cgroups (Control groups) 
- Other Kernel capabilities 

---
#### 3.1.1 Namespaces

- Namespaces tell processes **what they can see**.
 
Namespaces is a Linux kernel feature that carry out the distinction between kernel resources.
Examples of resources are process IDs, hostnames, files, usernames, network access names, UNIX time sharing (UTS) and inter-process communications (IPC). 
Any individual process can only view or use the namespace associated with that particular process.

On Linux, you can check your system's namespace by following command:
        
        lsns

example output:
````
        NS TYPE   NPROCS   PID USER          COMMAND
4026531835 cgroup    110  2032 htetpainghtun /lib/systemd/systemd --user
4026531836 pid       108  2032 htetpainghtun /lib/systemd/systemd --user
4026531837 user       88  2032 htetpainghtun /lib/systemd/systemd --user
4026531838 uts       110  2032 htetpainghtun /lib/systemd/systemd --user
4026531839 ipc        90  2032 htetpainghtun /lib/systemd/systemd --user
4026531840 mnt       109  2032 htetpainghtun /lib/systemd/systemd --user
4026532008 net        89  2032 htetpainghtun /lib/systemd/systemd --user
````
_Following outputs are parent namespaces started when booting with systemd._

---
#### 3.1.2 Cgroups (Control Groups)

- Cgroups tell processes **what they can use**.

Cgroups is a Linux kernel feature that limits, accounts for, and isolates the resource usage (CPU, memory, disk I/O, network, etc.) of a collection of processes.
One of the design goals of cgroups is to provide a unified interface to many different use cases, from controlling single processes 
(by using nice, for example) to full operating system-level virtualization (as provided by OpenVZ, Linux-VServer or LXC, for example). 

Cgroups provides:
- Resource limiting
    - groups can be set to not exceed a configured memory limit
- Prioritization
    - some groups may get a larger share of CPU utilization or disk I/O throughput
- Accounting
    - measures a group's resource usage
- Control
    - freezing groups of processes, their checkpointing and restarting

On Linux system with systemd, you can check your system's namespace by following command:

    systemd-cgls 

Outputs may look like this:
````
├─system
│ ├─1 /usr/lib/systemd/systemd --switched-root --system --deserialize 20  
│ ...
│      
├─user
│ ├─user-1000
│ │ └─ ...
│ 
│ └─1 /sbin/init splash
└─system.slice 
  ├─irqbalance.service 
  │ └─911 /usr/sbin/irqbalance --foreground
  ├─containerd.service 
  │ └─1008 /usr/bin/containerd
  ├─systemd-udevd.service 
  │ └─334 /lib/systemd/systemd-udevd
  ├─cron.service 
  │ └─899 /usr/sbin/cron -f
  ├─thermald.service 
  │ └─936 /usr/sbin/thermald --systemd --dbus-enable --adaptive
  ├─docker.service 
  │ └─1549 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
  ├─ssh.service 
  │ └─1032 sshd: /usr/sbin/sshd -D [listener] 0 of 10-100 startups
  ...
````
---

#### *Further Reading*

If you want to learn more about Cgroups Redhat has
[documentation on how to manage resources on RHEL7 Linux](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/resource_management_guide/index)
(but the concept is same for Linux distros).

I also recommend this [super cool video](https://www.youtube.com/watch?v=sK5i-N34im8) (54mins) from Jérôme Petazzoni @PyCon 2016 

which explains Namespaces and Cgroups and demonstration on to make a bare minimum linux container from scratch.
If you like this, you can also check out
[full video on his Introduction to Docker talk @PyCon 2016](https://www.youtube.com/watch?v=ZVaRK10HBjo&list=PLkCdZRqnOdeX9_zQ0lVmTAv5lyR0PN5kg) (3hrs).

[This one "Building a container from scratch in Go"](https://www.youtube.com/watch?v=Utf-A4rODH8) (20mins) is from Liz Rice 
where she demonstrated a bare container from scratch with GoLang.

---

Now, we know what a Linux container is all about.             
And I also mentioned about Docker dropping LXC from its core, why and how?

> First up, LXC is Linux-specific. It is was a problem for a project that had aspirations of being multi-platform.
  Second up, being reliant on an external tool for something so core to the project was a huge risk that could hinder
  development.
  As a result, Docker. Inc. developed their own tool called libcontainer as a replacement for LXC. The goal of
  libcontainer was to be a platform-agnostic tool that provided Docker with access to the fundamental container
  building-blocks that exist in the host kernel.
  Libcontainer replaced LXC as the default execution driver in Docker 0.9.

A quotation from Nigel Poulton's [Docker Deep Dive](https://www.amazon.com/Docker-Deep-Dive-Nigel-Poulton/dp/1521822808)
 
![dockerwithlxc](https://user-images.githubusercontent.com/47061262/146825434-3d214b14-d1cb-405f-b915-f7cd66929683.png)
 _Figure 3.1.1 docker implmentation with libcontainer and LXC underneath_
 
Omg! There are even more components there but hey, don't worry, these are just extra study I did because I think it is awesome. 
You won't need to know about these underlying layers just to start working with Docker which we will be doing later.
But in order to master it, these concepts and components are worth taking your time and hey, it is cool to be a nerd about it.  

After this section, you might as well be confused about this LXC and Docker things. 
So, I will sum up with this diagram. 
You don't need to know about Podman at all but think I will mention it in later episodes or chapters but for now you just need to know that
[Podman](https://podman.io/) is popular docker alternative from [RedHat org](https://www.redhat.com/). 

![LXC-docker-podman](https://user-images.githubusercontent.com/47061262/146827193-c7d68883-90d5-455e-ac48-4ba1979063b2.png)

 *Figure 3.1.4 most popular available containerization solutions*

---

### 3.2 Open Container Initiative

When talking about Docker and containers, **Open Container Initiative (OCI)** also play important part throughout the history of containers.

The *Open Container Initiative* is an open governance structure 
for the express purpose of creating open industry standards around container formats and runtimes.

The Open Container Initiative (OCI) is a lightweight, open governance structure (project), formed under the auspices of the Linux Foundation, for the express purpose of creating open industry standards around container formats and runtime. The OCI was launched on June 22nd 2015 by Docker, CoreOS and other leaders in the container industry.

The OCI currently contains two specifications: the Runtime Specification (runtime-spec) and the Image Specification (image-spec). 
- Runtime-spec outlines how a container run time should look like. 
(note: Docker's container runtime "[runC](https://www.docker.com/blog/runc/)" is actual implementation of runtime-spec and is donated to OCI)
- Image-spec outlines how to a container image should look like.

_for full information, you can visit https://opencontainers.org/about/overview/_

---

#### 3.3 Docker Engine 

The Docker engine is the core of the Docker that runs and manages Docker containers and images.

> **Docker Engine** is an open source containerization technology for building and containerizing your applications. 
 Docker Engine acts as a client-server application with:   
> - A server with a long-running daemon process dockerd.
> - APIs which specify interfaces that programs can use to talk to and instruct the Docker daemon.
> - A command line interface (CLI) client docker.

*[Official Docker Engine Docs](https://docs.docker.com/engine/)*

![docker-engine](https://user-images.githubusercontent.com/47061262/146833772-bf336eba-6658-4b37-90c9-9dba4479229b.png)

 *Figure 3.3.1 Docker engine overview*

Like real engines, the Docker engine is modular in design and built from many small specialised tools. Where possible, these are
based on open standards such as those maintained by the Open Container Initiative (OCI).         
Docker underwent many design changes over the years to achieve what we had today.

---

### 3.3.1 Docker Client 

The Docker client `docker` is the primary way that many Docker users interact with Docker. When you use commands such as  `docker run` , the client sends these commands to dockerd, which carries them out.

---

### 3.3.2 Docker Engine API 

Docker provides an RESTful API for interacting with the Docker daemon (the Docker Engine API), as well as SDKs for Go and Python. 
The SDKs allow you to build and scale Docker apps and solutions quickly and easily. 
If Go or Python don’t work for you, you can use the Docker Engine API directly via `wget` or `curl`, or the HTTP library which is part of most modern programming languages.
*For more information about SDKs, visits: https://docs.docker.com/engine/api/sdk/*


### 3.3.3 Docker Daemon

The Docker daemon (dockerd) listens for Docker API requests and manages Docker objects such as images, containers, networks, and volumes.
It mainly implements:
Remote API that we use to mange with our containers.
Internal networking where containers can interact between them or with the host.
Volumes where our data live and mount points within host.
Images mangaement (pulling, building, managing and pushing).
You can communicate with your docker daemon with following commands;

On linux:
   ````
   systemctl status docker 
   ````
On Mac:
   ````
   launchctl status docker
   ````
   
On any system with docker installed:
   ````
   dockerd help 
   ````

*We don't normally communicate with Docker daemon this way. This is rather for debugging and advanced configuration purpose. We do it via docker-cli.*

![docker daemon](http://blog.itaysk.com/images/2018-02-06-the-hitchhickers-guide-to-the-container-galaxy_2.png)

  _Figure 2.3.2 docker daemon_
  

Initially, Docker daemon is the one making all the jobs around Docker and thus leading to monolithic nature of Docker. 
This monolithic feature of Docker daemon becomes problematic because: 
- It's hard to innovate on
- It got slower
- It wasn't what ecosystem wanted

So, Docker, Inc. had to tear this docker daemon apart to shift workloads away from daemon as much as possible and modularize into smaller specialised tools.
These specialized tools can be swapped out, as well as easily re-used by third parties to build other tools.

>*This plan follows the tried-and-tested Unix philosophy of building
small specialized tools that can be pieced together into larger tools
The major components that make up the Docker engine are the Docker daemon, containerd, runc and other networking and storage plugins.*


![docker_oci](https://user-images.githubusercontent.com/47061262/146834416-c754c05d-634b-4f0f-895d-3c6c076b4943.png)

 *Figure 3.3.2 Docker workflow overview*

Let's meet our new components.

---

### 3.3.4 Containerd 

**[Containerd](https://containerd.io/)** is an industry-standard container runtime with an emphasis on simplicity, robustness and portability.
Containerd is available as a daemon for Linux and Windows. It manages the complete container lifecycle of its host system, from image transfer and storage to container execution and supervision to low-level storage to network attachments and beyond.

Containerd is the high level runtime used on top of runC in Docker (and many other projects).
It is responsible for making containers and managing them. 
It achieves this via two steps;     
- recive instruction to create containers 
- instruct runC to create container 

Getting started with Containerd? 
https://containerd.io/docs/getting-started/

---

### 3.3.5 runC and shim 

**runC** is a lightweight, portable container runtime which implements OCI-runtime standard.
It includes all of the plumbing code used by Docker to interact with system features related to containers. 
In docker, it is the thing that actually creates your containers with instructions from containerd.
runC is just here to create containers and after this, it exits and shim become container's parent process.

runC is the implementation of OCI-runtime and is donated to [CNCF](https://www.cncf.io/) by Docker. 

**shim(docker shim)**

--- 

