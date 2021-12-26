So far, we have learnt how to containerize our application and how to build Docker images and run containers and how they communicate between each other.

In this chapter, we will learn how to create multi-container applications with `docker-compose`
In real world micro-services, you will be running many containers along with many volumes mounted and networks also.
So, we might want to organize them and manage them all together.
This is where "Docker Compose" comes in.

---

###  Running multiple containers with Compose
 
Compose is a tool for defining and running multi-container Docker applications. 
With Compose, you use a YAML file to configure your application’s services. 
Then, with a single command, you create and start all the services from your configuration. 

Instead of gluing each microservice together with scripts and long docker commands, 
Docker Compose lets you describe an entire app in a single declarative conﬁguration ﬁle, and deploy it with a single command.


---




