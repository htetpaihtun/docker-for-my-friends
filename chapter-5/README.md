At this point, we have learnt everything we need to build a real application.
But first, we need to master `docker cli`.

### 5.1 Hacking Docker CLI 

When you are new to docker nad just started using docker to do something, we should always do 2 things; 
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

But we don't want to manually start a container by manually doing ports mapping, volume mapping, attaching to network everytime we want to run.  
We want our container to actually start from pre-built docker image.

---

### 5.3 Dockerfile 

A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image. 
Using `docker build`, users can create an automated build that executes several command-line instructions in succession. 

Dockerfile is said to be a starting point of the image and it should specify all the requirements of the image/app.
This will include;
- the image filesystem (dependencies, volume mounts and application code itself)
- how to start running our application (entry points and run commands) 
- how the application should behave (healthcheck and restart policies)
- where to run our application (port binding, service exposing)
- some additional metadata (labels and maintainer)

Dockerfile gives birth to your application image and therefore, quality of your image/app varies greatly based on Dockerfile.
While exploring its contents, I will also try to explain and suggest best Dockerfile practices for performance, storage and security concerns.

Conatiners are supposed to be light-weighted and portable. 
So, while building docker images, we should try to reduce its size and layers as much as possible 
so that you can ship images around swiftly.
But there's argument saying sometimes rich-feature images/containers are better overall.
So, we also need to maintain necessary (or potentially important) features untouched while reducing it's image size.

Dockerfile contains INSTRUCTIONS and METADATA for our images.
- INSTRUCTIONS are the ones that add extra layers on top of other images since they manipulate with our image filesystem. 
(We learnt about docker images' filesystem nature in previous chapter). 
You will see them adding extra layer per instruction directly with `docker inspect --format "{{.RootFS.Layers}}"` command.
- METADATA are the ones that defines how our appilcation should be started and running and exited 
and also provides additional information about our images. 
These kind doesn't directly add layers to our images but you can see that took place with `docker history` command.

Let's dive further..

---

#### 5.3.1 FROM Instruction 

The FROM INSTRUCTION initializes a new build stage and sets the Base Image for subsequent instructions.
Every Dockerfile starts with at least one FROM instruction usually at the start as based image.

The **base image** for your application provides the Linux libraries required by the application. 
The base image that you choose affects the versatility, security, and efficiency of your container.

You can use `FROM` instrcution with any image you like as long as it is pullable via Container Registry or locally available.

You will always want to use just possibly smallest image just enough to run out application. 

Let's create a Docker image with base ubuntu image. 
In `Dockerfile`, 
````
FROM ubuntu:20.04 
````
When using `FROM` instruction, a good practice is to include the exact image tag you want. 
By default, Docker will pull image from Dockerhub with 'latest` tag.
If you want to specify other registry, you can use something like this :
````
FROM registry.access.redhat.com/ubi8/ubi:8.1
````
Defining exact image tag and registry provides your app more stablity.
When using 'latest' tag, if provider releases another version with latest tag, you will have to pull that image and 
this is not optimal for image layer caching which we will talk about in later sections.
Also, you should always visit respective image's Docker hub page or the official documentation page. 

We will talk about `docker build` in details, in next section but first, let's build with `docker build` command,
````
docker build -t my-image . 
````
- `-t` defines image tag in format 'image:tag', we can obmit 'tag' part if we want (default is latest).
- `.` is location to your Dockerfile.

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

You can see `FROM scratch` with says 'let's build from scratch'.
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

**Reminder**: Creating new images will not create whole new image stack, 
but rathers writes on top of the existing image (that image can also have multiple layers on its own stack).
In other words, it will not inheritates it but rather depends on it. 
and tag it as whole new image.

#### 5.3.2 ARG Instruction

`ARG` instructions are used together with `FROM` instructions.
`FROM` instructions support variables that are declared by any `ARG` instructions that occur before the first FROM.
For example, you can do something like this in our `Dockerfile`:
````
ARG VERSION=20.04
FROM ubuntu:$VERSION
````
Note: An `ARG` declared before a `FROM` is outside of a build stage, so it can’t be used in any instruction after a `FROM` instruction.

#### 5.3.3 RUN Instruction 

The RUN instruction will execute any commands in a new layer on top of the current image and commit the results. 
The resulting committed image will be used for the next step in the Dockerfile.
This is one of the INSTRUCTION types that adds extra layers , since we manipulate the base image's filesystem and enviornment.
Let's include something like this in our `Dockerfile`:
````
ARG VERSION=20.04
FROM ubuntu:$version
ARG VERSION
RUN echo $version > image_version
````
- `echo $version > image_version` command writes `$version` to `image_version` file.
- `ARG VERSION` is added to demonstrate how `ARG defined before `FROM` instruction can't further be used in later instruction.

Build it with `docker build -t my-image .` 
Run with `docker run -it my-image cat image_version`.
Output : `latest`

`RUN` has two from 
- Shell form : the form we used in previous example. Format : `RUN <command>`. 
In shell form, the command is run in a shell, with `/bin/sh -c` on Linux based contianers.
- Exec form : format: `RUN ["executable", "param1", "param2"]`.
The exec form makes it possible to avoid shell string munging, and to RUN commands using a base image that does not contain the specified shell executable.

So, following `RUN` instruction in shell form (say we use Linux-based containers); 
````
RUN echo $PWD
````
Would be the same as this in exec form; 
````
RUN [["/bin/bash", "-c", "echo", $PWD ]
````

The RUN instruction will execute any commands in a new layer on top of the current image and commit the results. 
The resulting committed image will be used for the next step in the Dockerfile.

Layering RUN instructions and generating commits conforms to the core concepts of Docker where commits are cheap and containers can be created from any point in an image’s history, much like source control. To simply, it add extra layer on stack and you can track and use it later. 

Since `CMD` adds extra layers to your image, you might want to reduce `CMD` instructions as much as possible.
You might want to do something like this most of the time;
````
RUN apk update ; apk upgrade  
````
or 
````
RUN apk update \ 
apk upgrade
````
---

#### 5.3.4 CMD Instruction

`CMD` instructions are similar to `RUN` but the difference is they don't immediately run during build stage and doesn't add extra image layer on top.
The main purpose of `CMD` is to provide defaults for an executing container. 
These defaults can include an executable, or they can omit the executable, in which case you must specify an `ENTRYPOINT` instruction as well.			
Like `RUN` instructions, `CMD` instructions also have shell form and exec form.

Let's try using one in our `Dockerfile`;
````
ARG VERSION=latest
FROM ubuntu:$VERSION
ARG VERSION
RUN echo $VERSION > image_version
CMD cat image_version
````
Let's build `docker build -t my-image .` and run `docker run -it --rm my-iamge`
Output : `latest` 


#### 5.3.5 LABEL Instruction 

The LABEL instruction adds metadata to an image as a key-value pairs.

For examples:
````
LABEL "maintainer"="htetpainghtun" \
	"app_name"="test app" \ 
	"cool"="true" \
	"description"="Cool test app"
````
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
To actually publish the port when running the container, 
use the -p flag on docker run to publish and map one or more ports, 
or the -P flag to publish all exposed ports and map them to high-order ports.

It can be something like this.
````
EXPOSE 80/tcp
EXPOSE 80/udp
````
And then in `docker run` command;

````
docker run -p 80:80/tcp -p 80:80/udp my-image
````
---

#### 5.3.7 ENV Instruction

The ENV instruction sets the environment variable <key> to the value <value>. 
This value will be in the environment for all subsequent instructions in the build stage and can be replaced inline in many as well. 
The value will be interpreted for other environment variables, so quote characters will be removed if they are not escaped.
	
Example:
````
ARG VERSION=20.04
FROM ubuntu:$version
ENV FOO="BAR"
CMD echo $FOO
````
Difference between `ARG` and `ENV` is that 
`ARG` **will not persist** in its final image.

---
	
#### 5.3.8 ADD Instruction



	
---
	
#### 5.3.9 COPY Instruction



#### 5.3.10 ENTRYPOINT Instruction


*Note* :If `CMD` is used to provide default arguments for the ENTRYPOINT instruction, 
both the `CMD` and `ENTRYPOINT` instructions should be specified with the JSON array format. 
We will be using this in a lot of scenario, it is best to stick to that format.

#### 5.3.11 VOLUME Instruction 



#### 5.3.12 USER Instruction 



#### 5.3.13 WORKDIR Instruction




#### 5.3.14 ONBUILD Instruction



#### 5.3.15 STOPSIGNAL Instruction



#### 5.3.13 HEALTHCHECK Instruction



#### 5.3.14 SHELL Instruction




#### 5.3.15 .dockerignore File













