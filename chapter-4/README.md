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

### 4.3 Docker Engine 

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

### 4.5 Docker Containers 

Now we know about docker images which are technically stopped containers, let's see those images in action:- **containers**.  

A **container** is the runtime instance of an image. In the same way that you can start a virtual machine (VM) from
a virtual machine template, you can start one or more containers from a single image.


#### 4.5.1 Running Containers

You can start a container with commands `docker container run` or `docker run`.

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
This is because with `-d` option, the container runs as daemon (background process).
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

Furthermore, we will investigate their lifecycle.

---

#### 4.5.2 Stopping Containers Gracefully

As I mentioned before, when you kill a container with `docker container rm` the container is killed without warning. 
The procedure is quite violent and poor container has no chance to complete its process and therefore, can't exit gracefully.

Under the hood, it's all about sending Linux/POSIX signals. 
- `docker container stop` command sends a SIGTERM signal to main process inside container (PID 1). 
  This gives container a chance to stop its process and finish its life-cycle gracefully.
  If the container doesn't stop within some amount of time, Docker sends SIGKILL signal.
- `docker cotainer rm` command sends SIGKILL commands straight. 

---

#### 4.5.3 Restart policies 

The thing about modularizing your application and running as one-of processes in container is 
it becomes easier to track in lower level but then when you are scaling up (what you definitely want and will do in microservice architecture),
it also becomes extremely hard to manage.
When you are managing micro-services, it is crucial to have some sort of self-healing nature to avoid higher level complexities. 

Restart policies are applied per-container, and can be configured imperatively on the command line as part of
docker-container run commands, or declaratively in YAML files for use with higher-level tools such as Docker
Swarm and Docker Compose (in later chapters) or Kubernetes.

In Docker, following restart policies exist;
- no (default)
- always
- unless-stopped
- on-failure

**No**: Do not automatically restart the container. (default)

**Always**: restart the container if it stops. 
If it is manually stopped, it is restarted only when Docker daemon restarts or the container itself is manually restarted. 

**On-failure**: Restart the container if it exits due to an error, which manifests as a non-zero exit code. 
Optionally, you can limit the number of times the Docker daemon attempts to restart the container using the :max-retries option.

**Unless-stopped**: Similar to always, except that when the container is stopped (manually or otherwise), 
it is not restarted even after Docker daemon restarts.

*Restart policies only apply to containers. Restart policies for swarm services are configured differently.* 

An easy way to demonstrate this is to start a new interactive container, with the --restart always policy, and tell it to run a shell process.   
When the container starts you will be attached to its shell. 
Typing exit from the shell will kill the container’s PID 1 process and kill the container.
However, Docker will automatically restart it because it has the `--restart` always policy. 

Let's make a nginx-server with restart always policies and attach to its terminal.
````
docker run --name my-nginx-server -it --restart always nginx bash
````
You will be attached to its shell by `it` options and then you `exit` from its shell and check again with 
`docker ps` , you will see it automatically restarted itself.
````
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS     NAMES
de5585e43b64   nginx     "/docker-entrypoint.…"   23 seconds ago   Up 15 seconds   80/tcp    my-nginx
````
Now you have to manually stop with `docker stop` command or force it with `docker rm`. 

Difference between the always and unless-stopped policies is that containers with the `--restart
unless-stopped` policy **will not be restarted** when the daemon restarts if they were in the Stopped (Exited)
state. 

Let's try and compare restart-always policy and unless-stopped policy.
````
docker run --name restart-always-nginx -d --restart always nginx 
docker run --name unless-stopped-nginx -d --restart unless-stopped nginx
````
And if you restart your docker daemon right away, your containers will restart regardless of restart policies. 
But if your container, somehow, exits or you stop manually with `docker stop` the containers will act according to restart policies. 

Let's restart the docker daemon with;

Using Docker Desktop UI or

On Linux
````
sudo systemctl restart docker
````
On Mac
````
sudo launchctl restart docker
````
And then check their status with 
````
docker ps -a 
````
You will notice that all of your containers restarted.

Output will be similar to;
````
CONTAINER ID   IMAGE     COMMAND                  CREATED              STATUS         PORTS     NAMES
c6b38bf5a3c2   nginx     "/docker-entrypoint.…"   About a minute ago   Up 3 seconds   80/tcp    unless-stopped-nginx
68ab3eb4c0e7   nginx     "/docker-entrypoint.…"   About a minute ago   Up 3 seconds   80/tcp    restart-always-nginx
````
But, if you stop them befrorehand, 
````
docker stop restart-always-nginx unless-stopped-nginx
````
And restart your daemon, then check the status.
````
docker ps -a
````
You will see, only the container with restart-always policies restarted.          
Output will be similar to;
````
CONTAINER ID   IMAGE                                 COMMAND                  CREATED         STATUS                      PORTS     NAMES
c6b38bf5a3c2   nginx                                 "/docker-entrypoint.…"   2 minutes ago   Exited (0) 42 seconds ago             unless-stopped-nginx
68ab3eb4c0e7   nginx                                 "/docker-entrypoint.…"   2 minutes ago   Up 19 seconds               80/tcp    restart-always-nginx
````
You can see that only unless-stopped-nginx has been restarted. 

---

#### 4.5.4 Health Check

The Docker engine  determine if the container is in a state of abnormality by whether the main process in the container exits. 
In many cases, this is fine, but if the program enters a deadlock state, or an infinite loop state,
the application process does not exit, but the container is no longer able to provide services. 
So, Docker did not detect this state of the container and would not reschedule it, 
causing some containers to be unable to serve, but still accepting user requests.

So, you might want to tell Docker how your container should be acting. 
Docker HEALTHCHECK allows you to tell  Docker how to determine if the state of the container is normal.

Let's create a nginx container with custom health check command.
````
docker run -dit --name nginx-health-care -p 8080:80 --health-cmd "curl localhost:80" nginx
````
You can see healthcheck status in status field of `docker ps`. After 30s, it will say "healthy". 
This way, you can know whether the container is acting normally or not.
This is useful when you are orchestrating many containers with higher level tools like Docker swarm or Kubernetes.

Let's try to make container unhealthy.
````
docker exec nginx-health-care rm /etc/nginx/conf.d/default.conf
docker exec nginx-health-care nginx -s reload 
````
The follwing is for demonstration purpose and you will rarely delete nginx's configuration file in real life.
- `rm etc/nginx/conf.d/default.conf` command delete nginx's default configuration file.
- `nginx -s reload` command restart nginx service inside container.
This generally makes nginx to behave imporperly. 
This is not detectable by Docker by default because Docker doesn't know the exact state of your process.
It only knows whether it exited or not.
When we check with `docker ps`, after some time, it will say "unhealthy".
````
````

Now, let's try to create a nginx-server with it's own self-healing mechanism.
````
docker run -d name self-healing-server -p 8080:80 --restart unless-stopped --health-cmd "curl localhost:80 || nginx -s stop" nginx 
````
- `curl localhost:80 || nginx -s stop` says try connect to localhost:80, if fails, stop the nginx service(which is PID in our case).

We will try to make nginx crashes.
````
docker exec nginx-health-care rm /etc/nginx/conf.d/default.conf
docker exec nginx-health-care nginx -s reload 
````
And let's watch it crashes and restart again.
````
watch docker ps -a 
````
Nice, we can monitor its whole lifecycle now.
We didn't specify its restart count limit, so it will keep restarting itself forever.

But, our way of crashing nginx server is not realistic since we just delete some config file from the container to make it fail.
So, restarting container doesn't actually fix the problem. 
This is not what happens in most case.

And also the way of making nginx stops in health check command is not correct at all.
So, you won't be doing health check this way. 
We will see how to actually use healthcheck in proper way in later chapters.
 
But, we do learn how to make basic healthcheck the Docker way. 
And, if you're paying attention, you will notice how the container's files perists through his restarts.
If we just remove container directly with `docker rm` and build exact same container with exact same command,
we will see it is working again. 
This is, in fact, how we exactly want our containers to act i.e. to die and completely replace with another one in its place.
So, like images, containers also have its own filesystems. 
We will learning about them in next chapter.

---

### 4.6 Docker Volumes

Stateful applications that persist data are important in the world of cloud-native and microservices applications.			
For this purpose, Docker provides 2 type of storage for us, non-persistent and persistent.


#### 4.6.1 Non-Persistent storage

When you are creating Docker containers, each Docker container gets its own non-persistent storage. 
This is automatically created for every container and is tightly coupled to the lifecycle of the container.

We have seen it in action in previous section. But let's play around a little bit more by creating ubuntu container.
````
docker run -it --name my-container ubuntu 
````
In the container, we will create some files there.
````
echo "Hello, I was here" >  myfile.txt
````
You can verify with `ls` command. And then we exit from the container.
Confirm with `docker ps -a`.
Let's start the container again.
````
docker start -ia my-container
````
- `a` means attach container's STDOUT, STDERR and SIG stream to your terminal
- `i` means attach container's STDIN to your terminal
And then we check with `ls` and read it's content with `cat myfile.txt`, We can see that our data persist through restarts.
If we create a new container from same image again, we won't be able to view our data we created.
		
So, it is visable that containers create a read-write layer on top of read-only layer of the container image it's based on. 
This writable layer of local storage is managed on every Docker host by a storage driver such as overlay2, aufs, btrfs and more. 

Containers are supposed to be ephermeral and **immutable**, meaning they can easily be started, deleted, replaced and replicated.
They do not persist its data beyond its lifetime.
Like I said before, we don't do configuration to containers directly. 
We rather destroy it and then create new container with different configuration from either different or same image.
It's called immutablity of containers. 

#### 4.6.2 Persistent Storage

So, if needed, containers need to persist the application's data elsewhere. 
This is where persistent volumes came in.

 - Volumes are separate objects that have their lifecycles decoupled from containers. 
  This means you can create and manage volumes independently, and they’re not tied to the lifecycle of any container. 
  Net result, you can delete a container that’s using a volume, and the volume won’t be deleted.

 - Volumes can be mapped to specialized external storage systems

 - Volumes enable multiple containers on different Docker hosts to access and share the same data.

# ADDED FIGURES

Let's start by creating a volume.
````
docker volume create myvol 
````
and check with
````
docker volume ls
````
Output will look like;

````
DRIVER    VOLUME NAME
local     myvol
````
You can also inspect volumes like images and containers.
````
docker volume inspect myvol
````
Output will be similar to; 
````
[
    {
        "CreatedAt": "2021-12-23T03:51:07+06:30",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/myvol/_data",
        "Name": "myvol",
        "Options": {},
        "Scope": "local"
    }
]
````
- By default, docker uses 'local' driver for your volumes.
- You can use labels to group your containers and volumes together.
- Mountpoints shows where your volume actaully exists in your host system.

You can remove volumes with
````
docker volume rm myvol
````
or 
````
docker volume prune
````
which will delete all unused volumes in your system.

We can see that, we can treat our volume as separate object from our containers. 

Let's create a volume and use it in our containers.
````
docker run -dit --name vol-container-1 --mount source=myvol,target=/vol ubuntu
docker run -dit --name vol-container-2 --mount source=myvol,target=/vol ubuntu
````
- `source` is the volume you want to use. (docker will create new one if it doesn't find exisiting one)
- `target` is the location inside that volume which will then binded to your volume.
Basically, "myvol" and "/vol" directory are linked together now.
Other directories inside the containers will not be mounted to the volume.

Let's write something into volume from vol-container-1.
````
echo "Container 1 was here." > /vol/container1.txt
ls -al /vol/
````
Deattach from container-1 with `Ctrl+ P Q` and attach to vol-container-2 to check the file.
````
docker attach vol-container-2 
ls -al /vol/
cat /vol/container-1.txt
````
You can see the text we wrote earlier from container-1.

Now, let's try binding our host system directory with a new container. 
I will show you how to make a dev enviorment with docker container.
You can achieve this without having to install progamming languages and dependencies, 
also having your host system clean and isolated.

Let's quickly make a go app that say hello.
Create a file named main.go and edit with your favourite editor. 

````
package main

import (
	"fmt"
)

func main() {
	fmt.Println("HELLO WORLD!")
}
````
And let's point this current directory to a go lang container.
````
docker run -it --name go-dev-env -v $PWD/:/code/ golang 
````
- `-v` option allows us to mount our host directory to the container.

And then in container, 
````
go run /code/main.go
````
Output: 
````
HELLO WORLD!
```` 
There you have it. 
Whatever changes you make will take effect immediately.
This way, you can write codes without having to install programming languages locally.
You can even use them as compiler for your code.
````
cd /code/
go build main.go
````
You will get binary file on your system and later you can delete or whatever you like to your container.
Try and run your binary on your clean host. 
````
./main
````
I use this method a lot because I like to keep my system clean and isolated.
Also it feels better when dealing with many version releases and damn npm packages. 

Let's revist to our nginx server. 
This time, we will be using our config file outside of the container and mounts into it.

In default.conf file;
````
server {
    listen       81;
    listen  [::]:81;
    server_name  localhost;

    location / {
        root   /usr/share/nginx/html;
        index  my-index.html index.htm;
    }
}
```` 
This is bare minimum nginx configuration file that listens at container port 81 (default was 80) 
and say "serve this my-index.html file under /usr/share/nginx/html directory". That's all you need to know.

And in my-index.html file;
````
<!DOCTYPE html>
<html>
<head>
<title>HELLO !</title>
</head>
<body>
<h1> Hello from nginx server using custom config file and custom index.html!</h1>
</body>
</html> 
````
This prints hello-world with HTML.

Now we will create a nginx server container with these two files mounted into it.

````
docker run -dit -p 8080:81 \
-v $PWD/default.conf:/etc/nginx/conf.d/default.conf \
-v $PWD/my-index.html:/usr/share/nginx/html/my-index.html \
 nginx
````
- In first volume mount, we will overwrite default.conf file with out own default.conf file from host system.
- In second volume mount, we will added our own my-index.html and pass to nginx.

Now, we know how to talk to containers via storage.

Next step, is the network.

---
### 4.6 Docker Networks


ELP



