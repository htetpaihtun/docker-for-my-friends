So far, we have learnt how to containerize our application and how to build Docker images and run containers and how they communicate between each other.

In this chapter, we will learn how to create multi-container applications with "Docker Compose".
In real world micro-services, you will be running many containers along with many volumes mounted and networks also.
So, we might want to organize them and manage them all together.
This is where "Docker Compose" comes in.

---

###  Docker Compose
 
Compose is a tool for defining and running multi-container Docker applications. 
With Compose, you use a YAML file to configure your application’s services. 
Then, with a single command, you create and start all the services from your configuration. 

Instead of gluing each microservice together with scripts and long docker commands, 
Docker Compose lets you describe an entire app in a single declarative conﬁguration ﬁle, and deploy it with a single command.

Once the app is deployed, you can manage its entire lifecycle with a simple set of commands. 
You can even storeand manage the conﬁguration file in a version control system.

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
```

----

#### Compose Files

Compose files defines how your application's overview architecture should looks like. 
It is self-documenting and serves as bridge between devlopment and operation sides.

Compose uses YAML ﬁles to deﬁne multi-service applications.
`docker-compose,yaml` files has 4 top-level keys:
- version
- services
- networks
- volumes

In Dockerfiles, we can't directly bind ports or mount volumes because Dockerfiles only build images.
In Compose files, we can do them as they run container directly but the images you defined in Compose files must be built first.(You can't use Dockerfiles)

Using Docker Compose properly, alone satisfies many specifications of "12-factors app methodology for building software-as-a-service apps".

---

### Running Multi-containers Applications

example 



---

#### Docker Swarm 

