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

Fig was a powerful tool, created by a company called Orchard.
Docker, Inc. accquired Orchard, the company and re-branded Fig as Docker Compose.

---

#### YAML

```YAML
%YAML 1.2
---
YAML: YAML Ain't Markup Language™

What It Is:
  YAML is a human-friendly data serialization
  language for all programming languages.
```

----

#### Compose Files

`docker-compose.yaml` is 

