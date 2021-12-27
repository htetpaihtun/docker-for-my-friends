So far, we have learnt how to containerize our application and how to build Docker images and run containers and how they communicate between each other.

In this chapter, we will learn how to create multi-container applications with "Docker Compose".
In real world micro-services, you will be running many containers along with many volumes mounted and networks also.
So, we might want to organize them and manage them all together.
This is where "Docker Compose" comes in.

---

###  Docker Compose
 
Compose is a tool for defining and running multi-container Docker applications. 

Instead of gluing each microservice together with scripts and long docker commands, 
Docker Compose lets you describe an entire app in a single declarative conﬁguration ﬁle, and deploy it with a single command.

Once the app is deployed, you can manage its entire lifecycle with a simple set of commands. 
You can even store and manage the conﬁguration file in a version control system.

Compose has commands for managing the whole lifecycle of your application:
- Start, stop, and rebuild services
- View the status of running services
- Stream the log output of running services
- Run a one-off command on a service

Fig was a powerful tool, created by a company called Orchard.
Docker, Inc. accquired Orchard, the company and re-branded Fig as Docker Compose.

---

#### YAML

```YAML
%YAML 1.2
---
YAML: YAML Ain't Markup Language™

What It Is:
  YAML is a human-friendly data serialization language for all programming languages.
  YAML is popular mainly because it is human-readable and easy to understand. 
  It can also be used in conjunction with other programming languages.
  YAML is also a superset of JSON.
  
Use_case:
  YAML is commonly used for configuration files and 
  in applications, where data is being stored or transmitted.
  As they're very popular in cloud-native devops tools, 
  most popolar configuration management tools (ansible), 
  CICD tools (Gitlab-CI),
  and container orchestration tools (Kubernetes, Docker Swarm)
  as their main data format.
---
```

---

### Running Multi-containers Applications
 
We have been working with single container applications.
Also we have learn't how to make them talk to each other.
So, We should be fine building a fully functional application.

The work flow is like this;
- Build an image ( `docker build`, `docker pull`)
  - Write a Dockerfile
  - Build it
  - Tag the image
  - Push it to container registry

- Manage container lifecycle
  - Start containers (`docker run`)
    - passes commands
    - create networks 
    - volumes mounting
    - ports binding 
  - Monitor containers
    - check logs (`docker logs`)
    - debug by logging to containers (`docker exec`)
  - Stopping containers
  - Restarting containers
  - Removing containers

There's so much to do just to run a single container. (we haven't even talk about scaling yet)
In real world microservices applications, in small startups, you might want to have like 30-50 containers running.
We can't be managing all these steps manually.

We want to define our application as in self-contained and self-documenting style and also manage with simple commands.

Don't even start by saying "Hey, I can run my API server and database in one container."
It destroys the sole purpose of Docker.
Also because of 
- There’s a good chance you’d have to scale APIs and front-ends differently than databases.
- Separate containers let you version and update versions in isolation.

In real world, you might also need to scale by replicating many servies and perform rolling updates also.
In regards of updating containers, remember the immutable natue of containers; 
we don't update containers, we want to make new ones to replace.
Without orchestration tools, it would be very painful to manually perform rolling updates.
This is not accptable in CICD world where we want to release many updates on daily basis.
Everything that can be automated, should be automated.

Therefore, we will use container orchestration tools like "Docker Compose" and "Docker Stack". 
- It provides "Declarative nature",in which you will have complete definition of you application's state and architecture. 
Both developers and system admins can easily understand what is going on.
- You can also have it in your version control system. 
- You can easily setup multiple environments because making just one environment is easy now. (like dev and prod)
 
Using Docker Compose properly, alone satisfies many specifications of "12-factors app methodology for building software-as-a-service apps".
see : https://12factor.net and https://github.com/docker/labs/tree/master/12factor

To make things simple, with Docker Compose;
- You will need less commands to start your containers.
- You can have complete overall architecture of your application in one file.

---

#### Compose Files

Compose files defines how your application's overview architecture should looks like. 
It is self-documenting and serves as bridge between devlopment and operation sides.

Compose uses YAML ﬁles to deﬁne multi-service applications.
`docker-compose.yaml` files has 4 top-level keys:
- version
- services
- networks
- volumes
- secrets
other top-level keys such as `config `also exists

In Dockerfiles, we can't directly bind ports or mount volumes because Dockerfiles only build images.

Let's take a look at an example MERN stack application here. 
MERN application is the one that is made with MongoDatabase + ExpressJS + ReactJS + NodeJS.

**DISCLAIMER**: I yeeted it from [this video](https://www.youtube.com/watch?v=0B2raYYH2fE) as "refrence".
So, shout out to them. 
Also visit them, they have a lot of cools videos: https://www.youtube.com/channel/UC4MdpjzjPuop_qWNAvR23JA
For simplicity sake, I renamed it, cut some parts and added to my repo directly.
Original repo: https://github.com/sidpalas/devops-directive/tree/master/2020-08-31-docker-compose

You can also contribute with your own multi-container application with Docker compose.
I am more them happy to have them here.
The only limitations are
- have to work by using single `docker-compose up` command.
- simple enough to be called example docker-compose apps.
I will try to review as much as I can.

Anyway, let's take a look at its contents.

````
├── client
│   ├── Dockerfile
│   ├── package.json
│   ├── public
│   ...
│   ├── src
│   └── yarn.lock
├── docker-compose.yml
...
└── server
    ...
    ├── Dockerfile
    ├── index.js
    ...
    ├── package.json
    └── yarn.lock
````
We can see, there's two contaierized applications, "client" and "server" with their respective Dockerfiles.
What we are really interested in here is `docker-compose.yaml" file. Let's explore it.

```YAML
version: "3"
services:
  react-app:
    image: react-app
    stdin_open: true
    ports: 
      - "3000:3000"
    networks:
      - mern-app
  api-server:
    image: api-server
    ports:
      - "5000:5000"
    networks:
      - mern-app
    depends_on:
      - mongo
  mongo:
    image: mongo:3.6.19-xenial
    ports:
      - "27017:27017"
    networks:
      - mern-app
    volumes:
      - mongo-data:/data/db
networks:
  mern-app:
    driver: bridge
volumes:
  mongo-data:
    driver: local
```
Docker-compose files have more keywords than we had in our Dockerfiles. 
For simplicity sake, I won't be writing all of them like I did for Dockerfile in previous chapter.

You can always refer to: https://docs.docker.com/compose/compose-file/.

The higher the version, the better it is. 
So I will be exploring most used fields from [version 3](https://docs.docker.com/compose/compose-file/compose-file-v3/)

Topmost `version` is usually the ﬁrst line at the root of the ﬁle. this deﬁnes the version of the
Compose ﬁle format (basically the API). 
You should normally use the latest version.

It’s important to note that the versions key does not deﬁne the version of Docker Compose or the Docker Engine. 
For information regarding compatibility between versions of the Docker Engine, Docker Compose, and
the Compose file format, visit: https://docs.docker.com/compose/compose-file/compose-versioning

---

#### Service Configuration References

In `services` section, we define the list of services (or containers) we want to run as part of our application.

Let's explore `services section`;

```YAML
services:
  react-app:
    image: react-app
    stdin_open: true
    ports: 
      - "3000:3000"
    networks:
      - mern-app
  api-server:
    image: api-server
    ports:
      - "5000:5000"
    networks:
      - mern-app
    depends_on:
      - mongo
  mongo:
    image: mongo:3.6.19-xenial
    ports:
      - "27017:27017"
    networks:
      - mern-app
    volumes:
      - mongo-data:/data/db
```

In `services` section,

-  the name of each service ("react-app","api-server","mongo").
 
- `image` specifies the image to start the container from. 
Can either be a repository/tag or a partial image ID.
Docker will search from dockerhub if you doesn't have it locally.

- `ports` binds host ports and container(service) ports. (like we did in `docker run ... -p ...`)

- `networks` the service should be attached to. (like we did in `docker ... run --network ...`)

- `volumes` the service is using. (like we did in `docker run ... -v ...`) 

- `build` defines configuration options that are applied at build time.
 example: 
```YAML
version: "3.9"
services:
  webapp:
    build: ./dir
```
- `labels` adds metadata to the resulting image. You can use either an array or a dictionary.
example:
```YAML
build:
  context: .
  labels:
    - "com.example.description=Accounting webapp"
    - "com.example.department=Finance"
    - "com.example.label-with-empty-value"
```
- `command` to overwrite default commands.
```
command: echo "this can overwrite `CMD` we used in Dockerfile"
```
- `depends_on` express dependency between services. 
example:
```YAML
version: "3.9"
services:
  web:
    build: .
    depends_on:
      - db
      - redis
  redis:
    image: redis
  db:
    image: postgres
```
This say web won't start before redis and postgres, if you start web specifically, it will start redis and postgres beforehand.
When stopping, web will stop first.

- `deploy` specifies configuration related to the deployment and running of services. 
This only takes effect when deploying to a swarm with docker stack deploy, and is ignored by docker-compose up and docker-compose run.
We will talk about this later.

- `environment` adds environment variables. You can use either an array or a dictionary.

- `healthcheck` configures a check that’s run to determine whether or not containers for this service are “healthy”. 
This will overwrite healthcheck we defined in Dockerfile.

- `logging` defines logging configuration for the service.

- `restart` can overwrites restart policies defined in Dockerfile. (restart option is ignored when deploying a stack in swarm mode)

- `secrets` grants access to secrets on a per-service basis using the per-service secrets configuration.

- `volumes` mount host paths or named volumes, specified as sub-options to a service.
If you mount a host path as part of a definition for a single service, and there is no need to define it in the top level volumes key.

---

#### Network Configuration References

In `networks` section, we define networks we want to use, 
if it exists Docker will use it, 
if it doesn't, Docker will create it for you and default driver is "Bridge".

Our example `docker-compose.yaml`:
```YAML
networks:
  mern-app:
    driver: bridge
    attachable: true
```
- `driver` specifies which driver should be used for this network.

-  `attachable` only used when the driver is set to overlay. 
If set to true, then standalone containers can attach to this network, in addition to services. 
If a standalone container attaches to an overlay network, it can communicate with services and standalone containers 
that are also attached to the overlay network from other Docker daemons.

---

#### Volume Configuration References

In `volumes` section, we mount our volumes and host filesystem.

Our example `docker-compose.yaml`:
```YAML
volumes:
  mongo-data:
    driver: local
    external: true
```
An entry under the top-level volumes key can be empty, 
in which case it uses the default driver configured by the Engine (in most cases, this is the local driver).

- `driver` specifies which volume driver should be used for this volume.

- `external` : if set to true, specifies that this volume has been created outside of Compose. 
docker-compose up does not attempt to create it, and raises an error if it doesn’t exist.

---

#### Secret Configuration References

Docker also has special object called `secrets` to define our sensitive information.
The top-level secrets declaration defines or references secrets that can be granted to the services in this stack. 
The source of the secret is either file or external.

Example: 
```YAML
secrets: 
  db_password:
     file: db_password.txt
  db_root_password:
     file: db_root_password.txt
```

---

#### Docker Swarm 


























