#### In this chapter, you will learn background of Docker.Inc and your brief answers to your very first questions about Docker starting with why Docker, what Docker, how Docker and more.

### 2.1 Why Docker
> Docker is here and there’s no point hiding. In fact, if you want the best jobs working on the best technologies,
you need to know Docker and containers.
If you want to thrive in the modern cloud-first world, you need to know Docker.

A quotation from Nigel Poulton's [Docker Deep Dive](https://www.amazon.com/Docker-Deep-Dive-Nigel-Poulton/dp/1521822808)

Containers are barely minimum computing space to run your application inside.

For software developers, it is now easier to manage your application since docker packages it for you.
They can now have the exact same computing environment as the production servers.
So, with Docker, the famous "It works on my machine" problem is no more.
You can run your containerized application anywhere without worrying about extra configuration and dependencies 
as your containerized application will have all the thing it needs to run.

For system admins, your computing resources are now easier to manage since applications come in as little self-containing workspaces 
(also because docker containers are much more easier to control than VM ifself).
You just need to manage them from high level control pane and not to mess with dependencies, installation and stuff.

Nowsaday, in cloud-native enviorments, almost every application runs with containers since they are small, mobile, easy to set up, easy to manage.

### 2.2 The Docker, Inc. 
Docker, Inc. is a San Francisco based technology company founded by French-born American developer and entrepreneur Solomon Hykes.
The company started out as PaaS provider called dotCloud.
DotCloud is a product built on Linux containers and the team built an in-house tool nick-named "Docker" and that was the origin of what we know as Docker today. 
Later, they abandoned PaaS side and the team worked hard to bring that technology into mainstream IT and spoiler-alert "they successed!". 

### 2.3 The Docker Technology
Docker is software that creates, manages and even orchestrates containers. 

Docker is originated from open-source project called "Moby" created by the Docker,Inc.
Linux container technologies such as LXC existed at that time before Docker and are major influence to Docker. 
At first, Docker directly utilises LXC containers underneath and makes them much more easier to manage and use.
In later years, Docker got rid of LXC and continously improved the overall architecture and delivered to us better and better everytime. 

Docker bascially makes us just say the magic word and makes containers for us out of thin air.                
So what does Docker do under the hood.                                
There are 3 main components which make part of the Docker:
  1. The Container runtime
  2. The Docker daemon
  3. The Orchrestrator

   <!--tmp image-->
![docker architecture overview](https://www.docker.com/blog/wp-content/uploads/2018/11/DockerEngineDiagram-1.png)

_Figure 2.3.1 docker architecture overview_
          

### 2.3.1 The Container Runtime 
The runtime operates at the lowest level and is responsible for managing containers such as starting and stopping. 
This may raise some questions about what is container runtime at its core.
We will get to them in Docker Engine section in next Chapter. 
For now, you can assume them as little rooms in OS that can run your application.
So, Docker runtime is the one making room(space/environment) for your code to run.
Docker implements this runtime in multi-layered architecture which includes: 
containerd(high-level runtime) and runc(low-level runtime).

### 2.3.2 The Docker Daemon

The Docker daemon is the service that runs on your host operating system that control lower level runtime.
This is the core of the Docker engine and the place where we ask to create containers for our needs. 
It receives our instructions to Docker via docker-cli and acts as the brain of Docker. 
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

  _Figure 2.3.2 docker daemon_

### 2.3.3 The Orchestrator (Docker swarm) 

Docker also provides with the service where we can make many nodes to work together on top of Docker.
Docker swarm eases that cooperation between many machines via network connection and it provides internal scheduling of the containers.

### *Personal take*

The most frequently asked question and also the most frequently explained topic about the Docker is "Docker vs VMs". 

The answer to that question is that it is not an 'either or' question.      
After all, you might need something to run your docker containers on and your host machine or single server is not always enough 
for that, the same way older monolith applications needed more server partitions to fully optimised hardware resources. 
And VMs are excellent for this purpose since they can do better hardware virtualization.
Also having your container engine runs on VMs ensure better security since they reduce attack surface to your host system.

As software developers, having access to test your code on complete clean and new enviornment is very advantageous 
in a way you can ensure your code will run in same behaviour everywhere. 
You can also make your own dev environment completely in docker and it is also extremely portable! 
So, Docker helps me build very reliable and powerful testing environments.

Also, contaienrs are the fundamental building block of modern microservices. 
They serve as basic computing blocks for cloud-native applications. 
As container orchestration technologies like Kubernetes are becoming more powerful and popular, containers become a must for scalable application development.

Since part of the Docker(Docker-CE) is an open-source, you can find learning resources very easily.

### *Free Learning resources*

- A beginner lab to learn fundamentals: https://dockerlabs.collabnix.com/
- Interactive learning playground: https://www.docker.com/play-with-docker/
- Interactive web-based lab: https://www.katacoda.com/courses/docker/

