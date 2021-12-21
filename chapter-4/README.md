#### Now let's take a closer look at Docker components. 

Before we start to look at Docker components. Let's learn about what actually is a container.

### 4.1 What is a Linux container 

As I mentioned before, in earlier years, Docker directly uses [Linux containers (LXC)](https://linuxcontainers.org/) as its runtime and then later dropped it.        
 So, let's take a look it LXC before.

![LXC architecture](https://www.baeldung.com/wp-content/uploads/sites/2/2020/11/Linux-Container-Architecture-1.jpg)

  _Figure 4.1.2 LXC architecture_

A Linux container, at its core, is made up of; 
- Namespaces
- Cgroups (Control groups) 
- Other Kernel capabilities 

---
#### 4.1.1 Namespaces

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
#### 4.1.2 Cgroups (Control Groups)

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

Here's the blog from CNCF that further explained about containers: [demystifying-containers](https://github.com/saschagrunert/demystifying-containers).

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
 _Figure 4.1.1 docker implmentation with libcontainer and LXC underneath_
 
Omg! There are even more components there but hey, don't worry, these are just extra study I did because I think it is awesome. 
You won't need to know about these underlying layers just to start working with Docker which we will be doing later.
But in order to master it, these concepts and components are worth taking your time and hey, it is cool to be a nerd about it.  

After this section, you might as well be confused about this LXC and Docker things. 
So, I will sum up with this diagram. 
You don't need to know about Podman at all but think I will mention it in later episodes or chapters but for now you just need to know that
[Podman](https://podman.io/) is popular docker alternative from [RedHat org](https://www.redhat.com/). 

![LXC-docker-podman](https://user-images.githubusercontent.com/47061262/146827193-c7d68883-90d5-455e-ac48-4ba1979063b2.png)

 *Figure 4.1.4 most popular available containerization solutions*

---

### 4.2 Open Container Initiative

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

#### 4.3 Docker Engine 

The Docker engine is the core of the Docker that runs and manages Docker containers and images.

> **Docker Engine** is an open source containerization technology for building and containerizing your applications. 
 Docker Engine acts as a client-server application with:   
> - A server with a long-running daemon process dockerd.
> - APIs which specify interfaces that programs can use to talk to and instruct the Docker daemon.
> - A command line interface (CLI) client docker.

*[Official Docker Engine Docs](https://docs.docker.com/engine/)*

![docker-engine](https://user-images.githubusercontent.com/47061262/146833772-bf336eba-6658-4b37-90c9-9dba4479229b.png)

 *Figure 4.3.1 Docker engine overview*

Like real engines, the Docker engine is modular in design and built from many small specialised tools. Where possible, these are
based on open standards such as those maintained by the Open Container Initiative (OCI).         
Docker underwent many design changes over the years to achieve what we had today.

---

### 4.3.1 Docker Client 

The Docker client `docker` is the primary way that many Docker users interact with Docker. When you use commands such as  `docker run` , the client sends these commands to dockerd, which carries them out.

---

### 4.3.2 Docker Engine API 

Docker provides an RESTful API for interacting with the Docker daemon (the Docker Engine API), as well as SDKs for Go and Python. 
The SDKs allow you to build and scale Docker apps and solutions quickly and easily. 
If Go or Python don’t work for you, you can use the Docker Engine API directly via `wget` or `curl`, or the HTTP library which is part of most modern programming languages.
*For more information about SDKs, visits: https://docs.docker.com/engine/api/sdk/*

You can simply send requests to your Docker engine API with following steps;

On Linux:

````
curl --unix-socket /var/run/docker.sock localhost:2375/$your_api_version/containers/json
````

With Docker Desktop, you will need to go to Docker Desktop setting and find "Expose daemon on `tcp://localhost:2375` without TLS" option and enable it.

And point your browser to : `localhost:2375/$your_api_version/containers/json`

---

### 4.3.3 Docker Daemon

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

  _Figure 4.3.1 docker daemon_
  

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

 *Figure 4.3.2 Docker workflow overview*

Let's meet our new components.

---

### 4.3.4 Containerd 

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

### 4.3.5 runC and shim 

**runC** is a lightweight, portable container runtime which implements OCI-runtime standard.
It includes all of the plumbing code used by Docker to interact with system features related to containers. 
In docker, it is the thing that actually creates your containers with instructions from containerd.
runC is just here to create containers and after this, it exits and shim become container's parent process.             

runC is the de facto implementation of OCI-runtime and is donated to [CNCF](https://www.cncf.io/) by Docker. 


The **shim** is integral to the implementation of daemonless containers (what we just mentioned about decoupling
running containers from the daemon for things like daemon upgrades).

Some of the responsibilities the shim performs as a container’s parent include:
- Keeping any STDIN and STDOUT streams open so that when the daemon is restarted, the container
doesn’t terminate due to pipes being closed etc.
- Reports the container’s exit status back to the daemon.

--- 

After this modular decomposition of Docker daemon, you might be thinking what responsibilities or features left in docker daemon.  
Some of the major functionality that still exists in the daemon includes; 
- image management 
- image builds
- the REST API
- authentication
- security
- core networking
- orchestration

---

### 4.4 Docker Images 

A Docker image is a unit of packaging that contains everything required for an application to run.  
This includes; application code, application dependencies, and OS constructs. If you have an application’s Docker image, the
only other thing you need to run that application is a computer running Docker (or other container runtime).

Docker images are like VM templates, stopped containers or object classes from programming languages. It defines your container. 

Images are made up of multiple layers that are stacked on top of each other and represented as a single object.
Inside of the image is a cut-down operating system (OS) and all of the files and dependencies required to run
an application. Because containers are intended to be fast and lightweight, images tend to be small.

There are mainly 2 ways to get your images;
- pulling from image registry (docker's default registry is [Docker Hub](https://hub.docker.com/))
- building one from Dockerfile 

--- 

### 4.4.1 Image Registries

Like we store out code in GitHub we store our images in [Docker hub](https://hub.docker.com/).
[Docker Hub](https://hub.docker.com/) is one of the most popular registry and also default in Docker.

You can search image from [Docker hub](https://hub.docker.com/) with command;
````
docker search hello-world
````
Output is similar to; 
```
NAME                                       DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
hello-world                                Hello World! (an example of minimal Dockeriz…   1599      [OK]       
kitematic/hello-world-nginx                A light-weight nginx container that demonstr…   151                  
tutum/hello-world                          Image to test docker deployments. Has Apache…   87                   [OK]
...
````

You can pull an image from [Docker hub](https://hub.docker.com/) with command;
````
docker pull hello-world 
````

Output is similar to; 
````
Using default tag: latest
latest: Pulling from library/hello-world
Digest: sha256:2498fce14358aa50ead0cc6c19990fc6ff866ce72aeb5546e1d59caac3d0d60f
Status: Image is up to date for hello-world:latest
docker.io/library/hello-world:latest
````
This pulls a image that prints hello-world to terminal from [Docker hub](https://hub.docker.com/)

You can list all docker images with command;
````
docker images
````
or 
````
docker image ls
````
and remove images with command;
````
docker rmi hello-world
````
or 
````
docker image rm hello-world
````
You will learn to build images from your own Dockerfile in next chapter.

---

#### 4.4.2 Inspecting Images

A container only needs the code and dependencies of the application or service it is running — it does not need anything else. 
This results in small images stripped of all non-essential parts.
Gernerally, Docker images do not ship with many different shells for you to choose from.   

The good rule of thumb is **"if your application doesn't need it, you better not include it"**.   
So, many application images ship without a shell or basic command line tools that you're familiar with on your Linux system. 
Your image should define your application only.    
Image also don’t contain a kernel — all containers running on a Docker host share access to the host’s kernel. For
these reasons, we sometimes say images contain just enough operating system (usually just OS-related files and
filesystem objects).

Let's take a look at difference between images.
First, we will pull some images from Docker hub.
````
docker pull alpine  
docker pull ubuntu
docker pull golang
````
and we will list them. 
````
docker images
````
Output is similar to; 
````
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
alpine       latest    c059bfaa849c   3 weeks ago   5.59MB
ubuntu       latest    ba6acccedd29   2 months ago  72.8MB
golang       latest    4d9c15f5493b   4 weeks ago   941MB
````
Notice how golang image is like 168 times larger than alpine image.

So, you'd clearly better be choosing and building right image for production.

Further more inspection,we will inspect each one of them.

`docker image inspect` command allows you to obtain all the information about specific image.

We will look at alpine image first;

````
docker image inspect alpine
````

You will see something like this in output;

````
...
"Type": "layers",
            "Layers": [
                "sha256:8d3ac3489996423f53d6087c81180006263b79f206d3fdec9e66f0e27ceb8759"
            ]
...
````

This indicates the image only has a single layer.
Let's look at golang image;
````
docker image inspect golang
````
In output;
````
...
"Type": "layers",
            "Layers": [
                "sha256:a36ba9e322f719a0c71425bd832fc905cac3f9faedcb123c8f6aba13f7b0731b",
                "sha256:5499f2905579e85017f919e25be9e7a50bcc30c61294f12479b289708ebb31fa",
                "sha256:a4aba4e59b40caa040cc3ccfa42a84bbe64e3da8d1e7e0da69100c837afd215a",
                "sha256:8a5844586fdb00f07529ad1b3eb20167ba3a176302ecccbae1fbb45620acb89f",
                "sha256:fcd5459f6d07e8f21ca20db8d9872d61ae0e63064de5cebdae30ccf870d58706",
                "sha256:b584572b0aabd77494ee94f0244dd2e186fa3ee5b5159985b6705115f72b7438",
                "sha256:32185d066b1479d6463b3f09004ed139263dd6242587dda5cfa7503db588504f"
            ]
...           
````
This indicates there are 7 layers in the images. So, what are image layers?

### 4.4.3 Images Layers

A Docker image is just a bunch of loosely-connected read-only layers, with each layer comprising one or more
files. 

![docker image layers](https://user-images.githubusercontent.com/47061262/146947646-690a8c83-a4bd-43dd-b517-f34db0ba1723.png)

 *Figure 4.4.1 Docker Image Layering* 

`docker history` command is another way of inspecting an image and seeing layer data. However, it shows
the build history of an image and is not a strict list of layers in the final image. For example, some Dockerfile
instructions (“ENV”, “EXPOSE”, “CMD”, and “ENTRYPOINT”) add metadata to the image and do not result in
permanent layers being created.
All Docker images start with a base layer, and as changes are made and new content is added, new layers are
added on top.

You can also get more information about images and layers with command; 
````
docker history golang
```` 
Output may look like; 
````
IMAGE          CREATED       CREATED BY                                      SIZE      COMMENT
4d9c15f5493b   4 weeks ago   /bin/sh -c #(nop) WORKDIR /go                   0B        
<missing>      4 weeks ago   /bin/sh -c mkdir -p "$GOPATH/src" "$GOPATH/b…   0B        
<missing>      4 weeks ago   /bin/sh -c #(nop)  ENV PATH=/go/bin:/usr/loc…   0B        
<missing>      4 weeks ago   /bin/sh -c #(nop)  ENV GOPATH=/go               0B        
<missing>      4 weeks ago   /bin/sh -c set -eux;  arch="$(dpkg --print-a…   408MB     
<missing>      4 weeks ago   /bin/sh -c #(nop)  ENV GOLANG_VERSION=1.17.3    0B        
<missing>      4 weeks ago   /bin/sh -c #(nop)  ENV PATH=/usr/local/go/bi…   0B        
<missing>      4 weeks ago   /bin/sh -c set -eux;  apt-get update;  apt-g…   227MB     
<missing>      4 weeks ago   /bin/sh -c apt-get update && apt-get install…   152MB     
<missing>      4 weeks ago   /bin/sh -c set -ex;  if ! command -v gpg > /…   18.9MB    
<missing>      4 weeks ago   /bin/sh -c set -eux;  apt-get update;  apt-g…   10.7MB    
<missing>      4 weeks ago   /bin/sh -c #(nop)  CMD ["bash"]                 0B        
<missing>      4 weeks ago   /bin/sh -c #(nop) ADD file:5259fc086e8295ddb…   124MB     
````

As you can see in output, the layers with size>0 are the ones shown in `docker inspect` outputs (top-most layer being the container image itself)
and also the total size of the underlying layers combined equals the size of the image itself.

Now we can know, underlying layers have their own filesystem.

![image-filesystem](https://user-images.githubusercontent.com/47061262/146952045-0e114759-5c84-4f4f-821d-6a6c6d9ac63b.png)

 *Figure 4.4.2 Docker Image Layering and Filesystem*

We will later learn how to manage them in proper manners with our home-brew Dockerfiles. 
So far we only have to know that 
- images can be nested with one-dimensional read-only layers
- images stack on top of each other and they have their own filesystems. 
- the bottom-most layer is called the base layer and top-most layer represents the container image itself.

---

### 4.4.4 Image Names and Tags 

**Image names** usually come in the format: `registry/path/app-name:tag`. (neither of them is required)

For example, 
- if your image comes from another registry like gitlab;
  ````
  REPOSITORY                                                          TAG              
  registry.gitlab.com/htetpainghtun/my-repo/my-app                    v1.0
  ````
- if your image comes from dockerhub;
  ````
  REPOSITORY                                                          TAG              
  htetpainghtun/my-app                                                v1.0
  ````
 -if your images comes from official dockerhub repo (top level images);
  ````
  REPOSITORY                                                          TAG
  golang                                                              latest          
  ````

**Tags** serves as version control for docker images. They are
arbitrary alpha-numeric values that are stored as metadata alongside the image. 

You can tag images as you like with command; 
````
docker tag hello-world my-app:v1
````
and check with;
````
docker images 
````

*Notice how the image ID stays the same as long as you don't modify its contents.*
*Also if you didn't specify the tag, docker will assumes 'latest'.*

---

### 4.5 Containers 

Now we know about docker images which are technically stopped containers, let's see those images in action:- **containers**.  

A **container** is the runtime instance of an image. In the same way that you can start a virtual machine (VM) from
a virtual machine template, you start one or more containers from a single

You can start a container with commands `docker container run` or `docker run`

Let's run our hello-world image which we pulled before. If docker can't find local image, Docker will search from docker hub. 
````
docker run hello-world
````
Output will be similar to; 
````
Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
````

You can find your currently running containers on your machine with commands; 
`docker container ps` or `docker container ls` or `docker ps`

Let's see; 
````
docker ps
````
The output doesn't show anything.
This is because the hello-world container runs a one-of process that prints such output and exits.

You can find all of your containers (both running or exited) with command; 
````
docker ps -a
````
Now, there you go;
````
CONTAINER ID   IMAGE                                 COMMAND                  CREATED         STATUS                     PORTS     NAMES
e3424a2c7dc7   hello-world                           "/hello"                 8 seconds ago   Exited (0) 7 seconds ago             amazing_pasteur
````
You can now see its exit status along with other information. 
Also notice how Docker automatically names your container in runtime which is "amazing_pasteur" in my case. 
You can specify your container name with `--name my-container` option.
The complete command would be ;
````
docker run --name my-container hello-world
````

In similar fashion to images, you can inspect containers with `docker inspect` command.
````
docker inspect my-container
````
It will show all the information about the container you specified.
You can delete with `docker container rm` or `docker rm` command
````
docker rm my-container
````

Now let's try running ubuntu container.
````
docker run -it ubuntu bash
````
- `-t` option means allocate a pseudo-tty
- `-i` option means keep STDIN open even if not attached
- `ubuntu` is the image we will be running 
- `bash` is the command we start our container with.
You can specify commands to start your container with. 
Let's say we don't want to run bash, instead we want to see filesystem in the containers.
The Unix commands for that is `ls -a`.
Let's see.
````
docker run --it --rm --name ubuntu-ls ls -a
````
- `--rm` option means clean up the container(`docker rm`) after it exited, so that we don't have messy footprints of containers histroy.
- `ls -a` is the command we passed 

Output will be similar to:
````
.   .dockerenv	boot  etc   lib    lib64   media  opt	root  sbin  sys  usr
..  bin		dev   home  lib32  libx32  mnt	  proc	run   srv   tmp  var
````
So far, we have been working with the containers that are running in foreground and exited after its process ends.
Let's run a container running in backgroud.
````
docker run -dit --rm --name ubuntu  
````
You will get prompted with container id and immediately back to your terminal. 
This is because with `d` option, the container runs as daemon (background process).
You can check the container running with `docker ps` command. 
Output will be like:
````
CONTAINER ID   IMAGE     COMMAND   CREATED          STATUS          PORTS     NAMES
33673b08a165   ubuntu    "bash"    13 seconds ago   Up 12 seconds             my-ubuntu
````
And if we want to get back to it, we can use `docker exec` and `docker attach` commands.
Let's try `docker exec`:
````
docker exec -it my-ubuntu bash
````
`bash` is necessary because what docker exec does is executing a command in a container.

In this case, we will be running `bash` command with options `-it` in our container named "my-ubuntu". 
You will get the same terminal as before. 

Another way is to attach to its entrypoint with `docker attach` command.
````
docker attach my-ubuntu
````
You will get to the same terminal as well. 
What `attach` command does is attaching your terminal’s standard input, output, and error (or any combination of the three) to a running container. 
This allows you to view its ongoing output or to control it interactively, as though the commands were running directly in your terminal.

Notice if you `exit` the container, the main process exited and the container is considered to be completed his process.
So, our container will exit and later be removed by `--rm` option. 
If you just want to deattach the container, escape sequence is `Ctrl+ p q`. 

In this example we are running ubuntu container and it has no pratical usage to include `-d` option. 
We will try more practial usage by running a real nginx server with `-d` "runs in background" option.

It is necessary practice to visit the documentation or, at least, docker hub page before using images from remote repositories.
In our case, it would be
[nginx official documentation](https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-docker/) and 
[docker hub page](https://hub.docker.com/_/nginx). 

After read the instructions, you'll be feel free to use it.
````
 docker run -d --name my-nginx-server -p 8080:80 nginx
```` 
- `-p` option publish a container's port(s) to the host with your defined port `$HOST_PORT:$CONTAINER_POT`.
- `-P` option publish all exposed ports to random ports.

First, let's check container status with `docker ps` and output will be like:
````
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS                                   NAMES
b31e6a998ca7   nginx     "/docker-entrypoint.…"   5 seconds ago   Up 4 seconds   0.0.0.0:8080->80/tcp, :::8080->80/tcp   my-nginx-server
````
Let's see if it's actually running on our localhost.
You can directly and simply point your browser to `localhost:8080` or `curl localhost:8080` in terminal.

![Screenshot 2021-12-21 at 23-43-11 Welcome to nginx ](https://user-images.githubusercontent.com/47061262/146971145-c85fbd43-22c5-438f-945d-7e72bfa81c4d.png)

Congratulation! We just ran our first practically useable nginx server with Docker.

You can stop and remove it with;
````
docker stop my-nginx-server 
docker rm my-nginx-server
````
You can run many containers with a single image as long as you have enough host ports available and your container name is unique.
````
docker run -d --name my-nginx-server-1 -p 8080:80 nginx
docker run -d --name my-nginx-server-2 -p 8081:80 nginx
docker run -d --name my-nginx-server-3 -p 8082:80 nginx
````
Clean up:
````
docker rm -f my-nginx-server-1 my-nginx-server-2 my-nginx-server-3
````
`-f` option means force and it terminates containers without noticing them. (**not optimal and recommended**) 

Tips: 
> You should try and play around different containers using basic Linux commands such as `bash`, `ps`, `ls`, `apt`,`hostname` and more 
 via `docker exec -it` and `docker attach` to make youself familiar with containers. 
> Also don't forget to visit https://docs.docker.com/reference/ and use `docker --help` everything you need is there.

---










