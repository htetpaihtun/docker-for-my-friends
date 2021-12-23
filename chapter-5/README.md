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
`--filter` option works very great when you are using with "label" metadata on containers.
Remember, you can always visit [documentation page](https://docs.docker.com/) for more info on specify command's flags and options.
- Using `--format` is more advanced and more powerful in general. 
It pretty-prints containers or other objects using a Go template. 
With `--format` you can extract some piece of information out of particular object and display it gracefully.
We can also use both of them combined to make more readable outputs.
For example, if we want to see only container-id and their respective image of all the containers that exited.
````
docker ps --filter "status=exited" --format "table {{.ID}}\t{{.Image}}"
````
As usual, please visit [the documentation page's formatting and log output section](https://docs.docker.com/config/formatting/) for more combination and format.

You can also read about go-templates here : https://golangdocs.com/templates-in-golang.

`--format` can also be used together with `inspect` commands too. 

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

---

Now that we warm ourselves up with docker-cli, let's build something.
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
- INSTRUCTIONS are the ones that add extra layers since they manipulate with our image filesystem. 
(We learnt about docker images' filesystem nature in previous chapter). 
You will see them adding extra layer per instruction directly with `docker inspect --format "{{.RootFS.Layers}}"` command.
- METADATA are the ones that defines how our appilcation should be started and running and exited 
and also provides additional information about our images. 
These kind doesn't directly add layers to our images but you can see that took place with `docker history` command.

Let's dive further..

---

#### 5.3.1 FROM Instruction 

The FROM INSTRUCTION initializes a new build stage and sets the Base Image for subsequent instructions.
Every Dockerfile starts with at least one FROM instruction usually in very beginning. 


#### 5.3.2 ARG Instruction



#### 5.3.3 RUN Instruction 



#### 5.3.4 CMD Instruction



#### 5.3.5 LABEL Instruction 



#### 5.3.6 EXPOSE Instruction



#### 5.3.7 ENV Instruction



#### 5.3.8 ADD Instruction



#### 5.3.9 COPY Instruction



#### 5.3.10 ENTRYPOINT Instruction



#### 5.3.11 VOLUME Instruction 



#### 5.3.12 USER Instruction 



#### 5.3.13 WORKDIR Instruction




#### 5.3.14 ONBUILD Instruction



#### 5.3.15 STOPSIGNAL Instruction



#### 5.3.13 HEALTHCHECK Instruction



#### 5.3.14 SHELL Instruction




#### 5.3.15 .dockerignore File













