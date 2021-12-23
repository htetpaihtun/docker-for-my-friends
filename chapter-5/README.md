At this point, we have learnt everything we need to build a real application. 

Let's just create a simple Go server that keeps track how many times you hit the end point.
In our `server.go`,

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
you can also use methond we used with bind mounts.
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

### 5.1 Containerization your application

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

### 5.2 Dockerfile 

A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image. 
Using `docker build`, users can create an automated build that executes several command-line instructions in succession.

Dockerfile is said to be a starting point of the image and it should specify all the requirements of the image/app.
This will include;
- the image filesystem (dependencies, volume mounts and application code itself)
- how to start running our application (entry points and run commands) 
- how the application should behave (healthcheck and restart policies)
- where to run our application (port binding, service exposing)
- some additional metadata (labels and maintainer)

---

#### 5.3 FROM Instruction 

The FROM instruction initializes a new build stage and sets the Base Image for subsequent instructions.
















