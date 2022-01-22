At this point, we have learnt everything we need to build a real application.
But first, we need to master `docker cli`.

### 5.1 Hacking Docker CLI 

If you are new to docker and just started using docker to do something, we should always do 2 things; 
- `docker help` and 
- [Docker documentaion](https://docs.docker.com/) 

`docker help` is a really handy command that eases our Docker cli experience.	
When we are lost or we don't know how to do something we just simply type `docker help`.
For example, when you are trying to pull an image from docker hub, you just type `docker pull help` and all of your solution will be there in a second.
If `docker help` command isn't enough for you, Docker has [a really good documentation](https://docs.docker.com/) just one goole away from you.
For most of the time, these two steps solve just fine.

Docker first-class API objects such as images, containers, storage, networks and etc, come with their their subcommands.
These come in handy when you are dealing with the particular objects.
For example, when you want to deal with your custom volume, you can use `docker volume` subcommand to deal with them.
Also, you can always add `--help` to any commands.
If you want to deal with images, you use `docker images --help` and so on. 

Another handy commands that you might be using at the most are list commands;
- `docker ps` to list all containers running (it is almost always good to use `docker ps -a` instead of `docker ps`). 
When you are monitoring your containers `watch docker ps -a` command can save you some times.
- `docker images` to list all conatiner images on your system. 
- `docker volume ls` to list all volumes.
- `docker networks ls` to list all networks and more. 

All of these listing commands have 2 ways of filtering.
- Using `--filter` command : you need to pass key=value pair. 
For example, you can list only containers that exited by; 
````
docker ps -a --filter "status=exited" 
````
`--filter` option works great paring with "label" metadata on containers.
Remember, you can always visit [documentation page](https://docs.docker.com/) for more info on specify command's flags and options.
- `--format` pretty-prints containers or other objects using a Go template. 
With `--format` you can extract some piece of information out of particular object and display it very neat.
We can also use both of them combined to make more readable outputs.
For example, if we want to see only container-id and their respective image of all the containers that exited.
````
docker ps --filter "status=exited" --format "table {{.ID}}\t{{.Image}}"
````
To find out what data can be printed, check all content as json first:
````
 docker container ls --format='{{json .}}'
````
As usual, please visit [the documentation page's formatting and log output section](https://docs.docker.com/config/formatting/) for more combination and format.

You can also read about go-templates here : https://golangdocs.com/templates-in-golang.

Another command you will be using frequently is `docker inspect`.
You can inspect first-class API objects like images, containers, volumes, networks and more. 
It will return low-level information on Docker objects.
Using `--type json` formatting or using `--format` with go template with save you a lot of time while inspecting images in higly repetitive cases.

Mostly, you will be using `inspect` commands together with `docker logs` command to debug containers.

Some other tips :
- When you're spinning up a container for ephermeral purpose, 
you might want to use `--rm` option to make them clean up themselves, it will help you managing unused containers.
- Utilizing autocompletion makes things easier and faster. 

**"When you are lost, just throw `--help` in".**

You will be familiar with docker-cli in short time and you won't be even notice when it happens, due to its nice and clean documentaion experience.

Now that we warm ourselves up with docker-cli, let's build something.

---

### 5.2 Containerization your application

Containers are all about making apps simple to build, ship, and run.

Packaging your apps into Docker image is called containerization you app.

The process should be like this;
- start with your application code and dependencies 
- create Dockerfile that describes your application, its behaviours and its dependencies.
- build a Docker image from the Dockerfile
- push the image to registry
- run container from the image

Once, you've done this cycle, 
your application is now containerized and ready to share and run on everywhere 
as long as there's docker engine.(or other container runtime)

![application containerization process](https://user-images.githubusercontent.com/47061262/147288551-dc466f1a-0955-423f-bc2b-880371852249.png)

*Figure 5.1.1 Application containerization process*

---

Let's start with code.

We will create a simple Go server that keeps track how many times you hit the end point.

In our `server.go` file,

```go
package main

import (
	"fmt"
	"net/http"
)

var counter = 0

func count(w http.ResponseWriter, req *http.Request) {
	counter++
	fmt.Fprintf(w, "Counter: %d", counter)
}

func main() {
	http.HandleFunc("/", count)
	http.ListenAndServe(":8090", nil)
}
```
Let's run this with
````
go run server.go
````
you can also use method we used before with mounts binding.
````
docker run -dit --rm \
--name go-counter \
-p 8090:8090 \
-v /home/htetpainghtun/docker-for-my-friends/test-code/go-server-example/:/my-go-app/ \
golang \
go run /my-go-app/server.go
````
Then, you can point your browser to `localhost:8090` or `curl localhost:8090` and see it is working fine.

But we don't want to manually start a container by manually doing ports exposing, volume mapping and setting entry points.  
We want our container to actually start from pre-built docker image.

---

### 5.3 Dockerfile 

A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image. 
Using `docker build`, users can create an automated build that executes several command-line instructions in succession. 

Dockerfile is said to be a starting point of the image and it should specify all the requirements of the image/app.
This should include;
- the image filesystem (dependencies, volume mounts and application code itself)
- how to start running our application (port exposing, entry points and run commands) 
- how the application should behave (healthcheck and restart policies)
- some additional metadata (labels and maintainer)

Dockerfile gives birth to your application image and therefore, quality of your image/app varies greatly based on Dockerfile.
While exploring its contents, I will also try to explain and suggest best Dockerfile practices for performance, storage and security concerns.

Conatiners are supposed to be light-weighted and portable. 
So, while building docker images, we should try to reduce its size and layers as much as possible 
so that you can ship images around swiftly.
But there's argument saying sometimes rich-feature images/containers are better overall.
So, we also need to maintain necessary (or potentially important) features untouched while reducing it's image size.

Dockerfile contains instructions for our images.
There are kinds those manipulate filesystem and those doesn't and just provide metadata. 
- First ones add extra layers on top of other images since they manipulate with our image filesystem. 
(We learnt about docker images' filesystem nature in previous chapter). 
You will see them adding extra layer per instruction directly with `docker inspect --format "{{.RootFS.Layers}}"` command.
- Second ones defines our appilcation should be started and running and exited 
and also provides additional data about our images. 
These kinds doesn't directly add layers to our images but you can see that took place with `docker history` command.

Also Dockerfile is helpful in self-documenting since Dockerfile provides overview information over our application.
It's very easy when you are doing a lot of microservices since Dockerfiles are easy to read and understand.

Let's dive further..

---

#### 5.3.1 FROM Instruction 

The FROM INSTRUCTION initializes a new build stage and sets the Base Image for subsequent instructions.
Every Dockerfile starts with at least one FROM instruction usually at the start as based image.

The **base image** for your application provides the Linux libraries required by the application. 
The base image that you choose affects the versatility, security, and efficiency of your container.

You can use `FROM` instrcution with any image you like as long as it is pullable via Container Registry or locally available.

You will always want to use just possibly smallest image just enough to run our application. 

Let's create a Docker image with base ubuntu image. 
In `Dockerfile`, 
```Dockerfile
FROM ubuntu:20.04 
```
When using `FROM` instruction, a good practice is to include the exact image tag you want. 
By default, Docker will pull image from Dockerhub with 'latest` tag.
If you want to specify other registry, you can use something like this :
```Dockerfile
FROM registry.access.redhat.com/ubi8/ubi:8.1
```
Defining exact image tag and registry provides your application more stablity.
When using 'latest' tag, if provider releases another version with latest tag, docker will pull that image and 
will sometimes makes your application mulfunction and unstable.
Also, you should always visit respective image's Docker hub page or the official documentation page. 

We will talk about `docker build` in details, in next section, but first, let's build with `docker build` command,
````
docker build -t my-image . 
````
- `-t` defines image tag in format 'image:tag', we can obmit 'tag' part if we want (default is latest).
- `.` is 'build-context', we will revisit about this but for now regard it as location to your Dockerfile.

Let's see with `docker images`
````
docker images my-imge
````
Output will be similar to: 
````
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
my-image    latest    ba6acccedd29   2 months ago   72.8MB
````
Wait, my image is created 2months ago. 
This is because we didn't add any other instructions in Dockerfile and
Docker will treat this image the same as our original docker image which is 'ubuntu:20.04' in my case.
You can also investigate with `docker image inspect` and such, as we've already learnt.

We can also build a new image without using base image from scratch.
`FROM scratch` says 'let's build from scratch'.
Example: [Ubuntu 20.04 image's Dockerfile](https://github.com/tianon/docker-brew-ubuntu-core/blob/bf61e139e84e04f9d87fff5dc588a3f0398da627/focal/Dockerfile)

You can have more than one `FROM` instructions in your Dockerfile.
You can re-use image from previously built image in same Dockerfile.
We will learn about them in multistage builds section.

If you delete your image with 
````
docker rmi my-image
````
You will see the output says:
````
Untagged: my-image:latest
````
Because of the nature of docker image layering, deleting the image will not result in deleting its base image.
In our case, we didn't modify anything, we just simply re-tag it.

**Reminder**: If we have shared layers underneath, creating new images will not create whole new image stack from start 
but rathers writes on top of the existing image.
In other words, it will not inheritates it but rather depends on it.
and tag it as whole new image.

---

#### 5.3.2 ARG Instruction

`ARG` instructions let you define build-time arguments.
`ARG` instructions are usually used together with `FROM` instructions.
`FROM` instructions support variables that are declared by any `ARG` instructions that occur before the first FROM.
For example, you can do something like this in our `Dockerfile`:
```Dockerfile
ARG VERSION=20.04
FROM ubuntu:$VERSION
```
---

#### 5.3.3 RUN Instruction 

The RUN instruction will execute any commands in a new layer on top of the current image and commit the results. 
The resulting committed image will be used for the next step in the Dockerfile.
This is one of the instruction types that adds extra layers, since we manipulate the base image's filesystem and enviornment.
Let's include something like this in our `Dockerfile`:
```Dockerfile
ARG VERSION=20.04
FROM ubuntu:$VERSION
ARG VERSION
RUN echo $VERSION > image_version
```
- `echo $version > image_version` command writes `$VERSION` to `image_version` file.

Note : First argument we defined exists outside first FROM instruction. So, if we want to carry it over we need to define again. 

Build it with `docker build -t my-image .` 
Run with `docker run -it my-image cat image_version`.
Output : `20.04`

`RUN` has two from 
- Shell form : the form we used in previous example. Format : `RUN <command>`. 
In shell form, the command is run in a shell, with `/bin/sh -c` on Linux based contianers.
- Exec form : format: `RUN ["executable", "param1", "param2"]`.
The exec form makes it possible to avoid shell string munging, and to RUN commands using a base image that does not contain the specified shell executable.

So, following `RUN` instruction in shell form (say we use Linux-based containers); 
```Dockerfile
RUN echo $PWD
```
Would be the same as this in exec form; 
```Dockerfile
RUN [["/bin/bash", "-c", "echo", $PWD ]
```

The RUN instruction will execute any commands in a new layer on top of the current image and commit the results. 
The resulting committed image will be used for the next step in the Dockerfile.

Layering RUN instructions and generating commits conforms to the core concepts of Docker 
where commits are cheap and containers can be created from any point in an image’s history, much like source control. 
To simply, it add extra layer on stack and you can track and use it later. 

Since `RUN` adds extra layers to your image, you might want to reduce `RUN` instructions as much as possible.
You might want to do something like this most of the time;
```Dockerfile
RUN apk update ; apk upgrade  
```
or 
```Dockerfile
RUN apk update \ 
apk upgrade
```
---

#### 5.3.4 CMD Instruction

`CMD` instructions are similar to `RUN` but the difference is they don't immediately run during build stage and doesn't add extra image layer on top.
The main purpose of `CMD` is to provide default command for an executing container. 
These defaults can include an executable, or they can omit the executable, in which case you must specify an `ENTRYPOINT` instruction as well.			
Like `RUN` instructions, `CMD` instructions also have shell form and exec form.

Let's try using one in our `Dockerfile`;
```Dockerfile
ARG VERSION=latest
FROM ubuntu:$VERSION
ARG VERSION
RUN echo $VERSION > image_version
CMD cat image_version
```
Let's build `docker build -t my-image .` and run `docker run -it --rm my-iamge`
Output : `latest` 

---

#### 5.3.5 LABEL Instruction 

The LABEL instruction adds metadata to an image as a key-value pairs.

For examples:
```Dockerfile
LABEL "maintainer"="htetpainghtun" \
	"app_name"="test app" \ 
	"cool"="true" \
	"description"="Cool test app"
```
Now, your image is more documented and you can easily filter using labels.

"Docker, show me all the cool app images" 
````
docker images --filter "label=cool=true"
````

"Docker, show me ID and names of containers running the cool app, along with their maintainer.
````
docker ps --filter "label=cool=true" \
--format "table {{.ID}}\t{{.Names}}\t{{.Label "maintainer"}}"
````

---

#### 5.3.6 EXPOSE Instruction

The EXPOSE instruction informs Docker that the container listens on the specified network ports at runtime.
You can specify whether the port listens on TCP or UDP, and the default is TCP if the protocol is not specified.

The EXPOSE instruction does not actually publish the port. 
It functions as a type of documentation between the person who builds the image and the person who runs the container, 
about which ports are intended to be published. 
To actually publish the port when running the container (i.e. at run time), 
use the -p flag on docker run to publish and map one or more ports, 
or the -P flag to publish all exposed ports and map them to high-order ports.

It can be something like this.
```Dockerfile
EXPOSE 80/tcp
EXPOSE 80/udp
```
And then in `docker run` command;

````
docker run -p 80:80/tcp -p 80:80/udp my-image
````
---

#### 5.3.7 ENV Instruction

The ENV instruction sets the environment variable.
This value will be in the environment for all subsequent instructions in the build stage. 
The value will be interpreted for other environment variables, so quote characters will be removed if they are not escaped.

Example:
```Dockerfile
ARG VERSION=20.04
FROM ubuntu:$version
ENV FOO="BAR"
CMD echo $FOO
```
Difference between `ARG`and `ENV` is that `ARG` exists at build time only.

---
	
#### 5.3.8 ADD Instruction

The ADD instruction copies new files, directories or remote file URLs from source and adds them to the filesystem of the image at the path destination.

Multiple source resources may be specified but if they are files or directories, 
their paths are interpreted as relative to the source of the context of the build (Dockerfile).

Let's try building our go server app we wrote earlier.
Our file directory will be like this:
````
├── app
│   └── server.go
└── Dockerfile
````

In Dockerfile;
```Dockerfile
FROM golang:latest
ADD app/*.go /app/
EXPOSE 8090
CMD go run /app/server.go
```
And then build the image with;
````
docker build -t go-app . 
````
Then run it with;
````
docker run -d --name go-app -p 8090:8090 go-app
````
Test with;
````
curl localhost:8090
````
`COPY` instruction serves the same function as `ADD` but `ADD` provide two additional features.
- you can use a URL instead of a local file/directory but while using remote location `ADD` won't provide any authentication.
- you can extract tar from the source directory into the destination.

Because image size matters, using `ADD` to fetch packages from remote URLs is strongly discouraged; 
you should use `curl` or `wget` instead. 
That way you can delete the files you no longer need after they’ve been extracted and you don’t have to add another layer in your image. 

--- 
	
#### 5.3.9 COPY Instruction

`COPY` instructions are generally the same as `ADD` instructions. 

`ADD` and `COPY` instructions allow you to set permission on the files copied with `--chown` flag: 
```Dockerfile
COPY --chown appuser:appuser . .
```
Although you can use `RUN` instructions to manually create user and manipulate file permissions,
just for adding user, we can use `USER` instruction.
Some images ship with default users, you can check by running their container with `cat /etc/passwd` or simply reading their documentation.

According to [the Dockerfile best practices guide](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#add-or-copy), 
we should always prefer `COPY` over `ADD` unless we specifically need one of the two additional features of ADD.
Using `ADD` command automatically expands tar files and certain compressed formats, 
which can lead to unexpected files being written to the file system in our images.
If you are copying local files to your Docker image, always use COPY because it’s more explicit.

We will just use `COPY` instead of `ADD`. and investigate file permissions.
In our Dockerfile.
```Dockerfile
FROM golang:latest
COPY --chown=USER app/*.go /app/
EXPOSE 8090
CMD ls -al /app
```

`ADD` and `COPY` instuction sources are not limited to host filesystem and you can copy files from other image's filesystem.
We will utilise it in multistage build section.

---

#### 5.3.10 ENTRYPOINT Instruction

The best use for ENTRYPOINT is to set the image’s main command, 
allowing that image to be run as though it was that command (and then use CMD as the default flags).

Like`RUN`, `CMD` and `SHELL`, `ENTRYPOINT` has two form; exec and shell.

*Note* :If `CMD` is used to provide default arguments for the ENTRYPOINT instruction,
both the `CMD` and `ENTRYPOINT` instructions should be specified with the JSON array format. 
We will be using this in a lot of scenario, it is best to stick to that format.

Example:
```Dockerfile
FROM ubuntu:latest
ENTRYPOINT ["ls", "-al"]
CMD ["/etc/"]
```
This image will print all the files under `/etc/` from the container when you run :
````
docker build -t ubuntu-file-list
docker run ubuntu-file-list 
````
But if we run,
````
docker run ubuntu-file-list /bin/ 
````
It will print files under `/bin` directory.
`ENTRYPOINT` will also effect when you are using `docker exec` command.

----

#### 5.3.11 VOLUME Instruction 

The VOLUME instruction creates a mount point with the specified name and marks it as holding externally mounted volumes from native host or other containers. 
The value can be a JSON array, `VOLUME ["/var/log/"]`, 
or a plain string with multiple arguments, such as `VOLUME /var/log` or `VOLUME /var/log /var/db`.

For example,
```Dockerfile
FROM ubuntu
RUN mkdir /myvol
RUN echo "hello world" > /myvol/greeting
VOLUME /myvol
```

This Dockerfile results in an image that causes `docker run` to create a new mount point at `/myvol` and copy the greeting file into the newly created volume.
The host directory is declared at container run-time: The host directory (the mountpoint) is, by its nature, host-dependent. 
This is to preserve image portability, since a given host directory can’t be guaranteed to be available on all hosts. 
For this reason, you can’t mount a host directory from within the Dockerfile. 
The VOLUME instruction does not support specifying a host-dir parameter. 
You must specify the mountpoint when you create or run the container.

---

#### 5.3.12 USER Instruction 

The USER instruction sets the user name (or UID) and optionally the user group (or GID) 
to use when running the image and for any RUN, CMD and ENTRYPOINT instructions that follow it in the Dockerfile.
For example,
```Dockerfile
FROM ubuntu:20.04
RUN useradd -u 1001 app_user
USER app_user
CMD id
```
It's best practice to set user permissions to your files and set to that user while running your application.

---

#### 5.3.13 WORKDIR Instruction

The `WORKDIR` instruction sets the working directory for any `RUN`, `CMD`, `ENTRYPOINT`, `COPY` and `ADD` instructions that follow it in the Dockerfile. 
If the `WORKDIR` doesn’t exist, it will be created even if it's not used in any subsequent Dockerfile instruction.
You can have more than one `WORKDIR` instructions in your Dockerfile according to the needs of your successive instructions.
```Dockerfile
FROM ubuntu:20.04
WORKDIR /work
CMD pwd
```

---

#### 5.3.14 ONBUILD Instruction

The `ONBUILD` instruction adds to the image a trigger instruction to be executed at a later time, when the image is used as the base for another build. 

---

#### 5.3.15 STOPSIGNAL Instruction

The STOPSIGNAL instruction sets the system call signal that will be sent to the container to exit. 
This signal can be a signal name in the format `SIG<NAME>`. (e.g. SIGKILL)

---

#### 5.3.13 HEALTHCHECK Instruction

The HEALTHCHECK instruction tells Docker how to test a container to check that it is still working. 
This can detect cases such as a web server that is stuck in an infinite loop and unable to handle new connections, 
even though the server process is still running.
`HEALTHCHECK` instructions look like this;
```Dockerfile
HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost/health || exit 1
```
You can also disble any healthcheck inherited from the base image

---

#### 5.3.14 SHELL Instruction

The SHELL instruction allows the default shell used for the shell form of commands to be overridden. 

The SHELL instruction is particularly useful on Windows where there are two commonly used and quite different native shells: cmd and powershell

---

#### 5.3.15 .dockerignore File

Before the docker CLI sends the context to the docker daemon, it looks for a file named `.dockerignore` in the root directory of the context. 
If this file exists, the CLI modifies the context to exclude files and directories that match patterns in it. 
This helps to avoid unnecessarily sending large or sensitive files and directories to the daemon and potentially adding them to images using `ADD` or `COPY`.
Example `.dockerignore` file looks like;
````
*/temp*
temp?
*.md
````
---

### 5.4 Docker build 

The docker build command builds Docker images from a Dockerfile and a “context”. 

There's alot of thing you can pass to docker build commands.
https://docs.docker.com/engine/reference/commandline/build/

But we won't be using too many of them, in most cases, we are only interested in passing build context and Dockerfile.

A build’s context is the set of files located in the specified PATH or URL.
Build context is the PATH that specifies where to find the files for the “context” of the build on the Docker daemon.
Like we did with
````
docker build -t imagename .
````
We simply say "Docker, please build me a image from here."
Here, `.` is the context or path to Dockerfile and all of your relative paths used inside Dockerfile will start from this point.

---

### Image Layer Caches

In CI/CD world where release cycles are too frequent, you have to build a lot of images.
Docker helps building process by utilising image layer caching.
We know how images layer stack on top of each other as read-only layers. 
Docker will only rebuild when there are changes to each layer in your build process.
Docker will re-use exisiting layer from previous build process as long as cache-hit.

The process is simply.
Let's say you have 5 layered images that you built previously.
But you make a change at 4th layer of the image and then you rebuild it.
Then Docker will take previous 3 layers directly from previous build and re-write another layer on top as 4th layer.
For 5th layer, Docker will build new layer regardless of changes made to the layer itself because images inherited filesystems from their base layer.

So, it is very important to utilise image layer caching as they can reduce a lot of time building your image.

Good rule of thumb is to do "what you think it's gonna change most frequently" as LAST layer (as much as you can) of the image. (like your source code)

For example,
```Dockerfile
FROM node:latest
WORKDIR /app
COPY package*.json .
RUN npm install
COPY src .
CMD npm run
```
In this case, you copied only package.json to install node dependencies, not the source code itself, because it is less likely to change than source code itself.
Assume there's source code changes and if you were to copy whole app folder in first stage, 
you will be repeating npm install process regardless of whether you added new packages to package.json or not.
This matters since it will not only take your computing resource, it will take your time too.
You can also utilise this with the combination of .dockerignore utility to ignore tracking of specific files that your application doesn't need.

---

#### 5.4.1 Multistage builds

Another thing to notice when building images is, if your app uses build process, you can ship your final image without those build tools.
As we said before, `COPY` instructions can copy anything from any image that is available in your system.
We also learnt how we can use containers as complier/builder in chapter-4.
Multistage builds are kinda combination of these 2 logic.

We will build our application with 1 image. (builder image)
Then we copy binary code from that image and ship it without builder image.

Since build tools are usually huge they're not optimal for shipping. 
This greatly reduce your final image size and also security reason by not exposing your source code.

For example, we have simple go app here.

In our Dockerfile,
```Dockerfile
FROM golang:1.17-alpine AS build
WORKDIR /go-app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
RUN go build -o ./main .

FROM alpine:latest
WORKDIR /app
COPY --from=build /go-app/main .
EXPOSE 8080
CMD ["./main"]
```
We just copied final binary file with `COPY` instruction and leave the builder image behind.

Build it with 
````
docker build -t go-multi-stage .
````
and run with 
````
docker run -d \
--name go-app \
-p 8080:8080 \
go-multi-stage
````
If we inspect our image with 
````
docker images go-multi-stage
````
Output will look like this;
````
REPOSITORY       TAG       IMAGE ID       CREATED         SIZE
go-multi-stage   latest    a5259a38887f   2 minutes ago   11.7MB
````
It only has 11.7MB size even though we used `golang:1.17-alpine` image in first `FROM` instruction.
We can also check its size with
````
docker images golang:1.17-alpine
````
Output:
````
REPOSITORY   TAG           IMAGE ID       CREATED       SIZE
golang       1.17-alpine   d8bf44a3f6b4   2 weeks ago   315MB
````
If we do list all images with `docker images`,
````
REPOSITORY                                                          TAG               IMAGE ID       CREATED         SIZE
go-multi-stage                                                      latest            a5259a38887f   5 minutes ago   11.7MB
<none>                                                              <none>            23e3c634ddbb   5 minutes ago   321MB
````
We can also see extra image with <none> name and tag. 
This is the build image we did get rid of in final image.

As you can see, utilising Multi-stage build greatly reduce your final image size and therefore enhances image shipping processes.

---

Best practices refrences I usually look up to:

https://developers.redhat.com/articles/2021/11/11/best-practices-building-images-pass-red-hat-container-certification

https://docs.docker.com/develop/develop-images/dockerfile_best-practices/ 

https://sysdig.com/blog/dockerfile-best-practices/

If you are interested about microservices, you might be checking out 12 factor applications.

Building 12 factor apps with Docker: https://github.com/docker/labs/tree/master/12factor

Official document: https://12factor.net/

---
