So far we have learnt about how to build our multi-container applications from one 'docker-compose.yaml' file.
We will now try to do it one multiple nodes (machines) and scale even larger.
Many people tends to stay away from Docker Swarm and Docker Stack because they think those are for multi-nodes management only.
In facts, Docker Swarm mode is fully funtional even if you're using only one node and Docker Stack has additional features that Docker Compose doesn't have.
So, Docker Compose might be good for local test developments but Docker Stack can go production. 

---

### 7.1 Docker Swarm

Docker Swarm provides cluster management and orchestration features, embedded in the Docker Engine, 
which are built using [swarmkit](https://github.com/docker/swarmkit). 
Swarmkit is a separate project which implements Docker’s orchestration layer and is used directly within Docker.

A swarm consists of multiple Docker hosts which run in swarm mode and act as managers and workers. 
A given Docker host can be a manager, a worker, or perform both roles. 
When you create a service, you define its optimal state 
(number of replicas, network and storage resources available to it, ports the service exposes to the outside world, and more). 
Docker works to maintain that desired state. 
For instance, if a worker node becomes unavailable, Docker schedules that node’s tasks on other nodes. 
A task is a running container which is part of a swarm service and managed by a swarm manager, as opposed to a standalone container.

One of the key advantages of swarm services over standalone containers is that you can modify a service’s configuration, 
including the networks and volumes it is connected to, without the need to manually restart the service. 
Docker will update the configuration, stop the service tasks with the out of date configuration, and create new ones matching the desired configuration.

Docker Swarm is two main things:
- An enterprise-grade secure cluster of Docker hosts
- An engine for orchestrating microservices apps

On the clustering front, Swarm groups one or more Docker nodes and lets you manage them as a cluster. 
Out-of-the-box, you get an encrypted distributed cluster store, encrypted networks, mutual TLS, secure cluster join
tokens, and a PKI that makes managing and rotating certiﬁcates a breeze. 
You can even non-disruptively add and remove nodes.

On the orchestration front, Swarm exposes a rich API that allows you to deploy and manage complex microservices apps with ease. 
You can deﬁne your apps in declarative manifest files and deploy them to the Swarm with native Docker commands. 
You can even perform rolling updates, rollbacks, and scaling operations.
Again, all with simple commands.

---

#### 7.1.1 Docker Swarm Features

In this section, we will learn about the features Docker Swarm provides us.

- Cluster management integrated with Docker Engine. (Docker Swarm is easy to setup and manage, once you have learnt docker, than other tools like Kubernetes)
- Decentralized design.
- Declarative service model: Docker Engine uses a declarative approach to let you define the desired state of the various services in your application stack.
- Scaling: For each service, you can declare the number of tasks you want to run. 
When you scale up or down, the swarm manager automatically adapts by adding or removing tasks to maintain the desired state.
- Desired state reconciliation: 
The swarm manager node constantly monitors the cluster state and reconciles any differences between the actual state and your expressed desired state.
- Multi-host networking: You can specify an overlay network for your services.
- Service discovery: Swarm manager nodes assign each service in the swarm a unique DNS name and load balances running containers.
- Load balancing: You can expose the ports for services to an external load balancer. 
Internally, the swarm lets you specify how to distribute service containers between nodes
- Secure by default: Each node in the swarm enforces TLS mutual authentication and encryption to secure communications between itself and all other nodes.
- Rolling updates: At rollout time you can apply service updates to nodes incrementally. 

---

#### 7.1.2 Swarm Mode 

On clustering front, Docker Swarm can easily initiate or join machines to the cluster.
Nodes can be anything from Virtual machines to Raspberry PIs to laptops to cloud instances to on-premise servers.
The only requirement is to have Docker installed and can communicate over network.

Nodes are conﬁgured as managers or workers. 
- Managers look aafter the control plane of the cluster, meaning things like the state of the cluster and dispatching tasks to workers. 
- Workers accept tasks from managers and execute them

Leader is the manager that initiates the Swarm.

The conﬁguration and state of the swarm is held in a distributed etcd database located on all managers. 
It’s kept in memory and is extremely up-to-date.

Something that’s game changing on the clustering front is the approach to security. 
TLS is so tightly integrated that it’s impossible to build a swarm without it.

Swarm uses TLS to encrypt communications, authenticate nodes, and authorize roles.
Automatic key rotation is also thrown in as the icing on the cake. 
And the best part… it all happens so smoothly that you don’t even know it’s there.

![swarm-diagram](https://user-images.githubusercontent.com/47061262/147570102-5fc79711-b82e-48d0-9f4a-65d4b569be5b.png)
*Figure 7.1.1 Docker Swarm Overview

Docker Desktop for Mac and Windows only supports a single Docker node. 
Alternatively, you can try Play with Docker at https://labs.play-with-docker.com.

To create a swarm cluster, we will first initiate a swarm mode from one node.
Running docker `swarm init` on a Docker host in single-engine mode will switch that node into swarm mode,
create a new swarm, and make the node the ﬁrst manager of the swarm.
````
docker swarm init
````
You will be prompted with information about cluster and how to add nodes to your cluster.

You can see available nodes from manager with `docker node ls`.  

You can run previously saved output from `docker swarm init` command and then run it on another machine to add nodes.
Beware of advertise addresss field, you will be using the same network that's available to other nodes in clutser.

You can regenerate token with `docker swarm join-token manager` or `docker swarm join-token worker`.

You can also lock/unlock the swarm.

Using `docker swarm leave` on leader node will delete the cluster.

Having many nodes as manager secures high availability of your cluser.

On the topic of HA, the following two best practices apply:
1. Deploy an odd number of managers. (incase you're having network partition problems)
2. Don’t deploy too many managers. (3 or 5 is recommended)

Swarm clusters have a ton of built-in security that’s conﬁgured out-of-the-box with sensible defaults — 
CA settings, join tokens, mutual TLS, encrypted cluster store, encrypted networks, cryptographic node ID’s and more.

---

### 7.1.3 Docker Services

On the application orchestration front, the atomic unit of scheduling on a swarm is the service. 
This is a new object in the API, introduced along with swarm, and is a higher level construct that wraps some advanced features around containers. 

A service is the definition of the tasks to execute on the manager or worker nodes. 
It is the central structure of the swarm system and the primary root of user interaction with the swarm.

`docker service` command is only available in Swarm mode. 
We declared services section in docker compose but if we list with `docker service ls` won't should you anything.
Docker Compose will only treat services as regular containers with maybe some more metadata on it.
Docker Swarm utilises services to its potential.
That includes scaling, rollback, update and other features. 

Services let us specify most of the familiar container options, such as name, port mappings, attaching to networks, and images. 
But they add important cloud-native features, including desired state and automatic reconciliation.
For example, swarm services allow us to declaratively deﬁne a desired state for an application that we can apply
to the swarm and let the swarm take care of deploying it and managing it.

You can create services in one of two ways:
- Imperatively on the command line with `docker service`.
In imperative methond, you have subcommand `docker service` to manage services at swarm level. 
You can do basic monitoring and CRUD things and scaling as well.
- Declaratively with a stack file. 
You can directly use compose files and feed it to Docker Stack but Docker Stack has many more fields that do not work with Comopose.

We will see them in action in later section.

But, you will always prefer declarative way.

---

### 7.1.4 Docker Secrets and Configs

Another things you could do with Docker Stack and Docker Swarm includes secrets management and configuration management.
They are like volumes but have its own manageable area on Docker Swarm. 
You will storing database password or API key in docker secrets and 
configuration like nginx configs and database configs in Docker configs.
And you will be able to mangae them with `docker secret` and `docker config` commands.

Docker won't provide any encryption regarding your secrets, you will need to manage them yourself.

---

### 7.2 Docker Stack 

Deploying and managing cloud-native microservices applications comprising lots of small integrated services at scale is hard.
Fortunately, Docker Stack is here to help. 
They simplify application management by providing; 
desired state, rolling updates, simple, scaling operations, health checks, and more! All wrapped in a nice declarative model

We will follow instructions from this super simple example from [Docker Docs]().

The process is simple. 
Deﬁne the desired state of your app in a Compose file, then deploy and manage it with the docker stack command.

The Compose file includes the entire stack of microservices that make up the app. 
It also includes all of the volumes, networks, secrets, and other infrastructure the app needs. 
In the same way that you can use Docker Compose to define and run containers, you can define and run Swarm service stacks.
The docker stack deploy command is used to deploy the entire app from the single file.

Wait. This is Docker Compose. 

But Docker Compose won't guarantee your current state meets the desired state meet. It only up and down things at once. 
You write compose file, you declare things, you bring them up, that's it. 
It can't actually scale or achieve self-healing.

The real power comes from Docker Stack, which is the ability to scale and automate things with simple declarative desired state. 

To accomplish all of this, stacks build on top of Docker Swarm, meaning you get all of the security and advanced features that come with Swarm.
In a nutshell, Docker Compose is great for application development and testing. 
Docker Stacks are great for scale and production.

To see them in action, let's follow this tutorial and play around.

https://docs.docker.com/engine/swarm/stack-deploy/

This tutorial doesn't demonstrate how to scale services, so we will do ourselves.
````
docker service scale stackdemo_web=3
````
Output will look like this;
````
stackdemo_web scaled to 3
overall progress: 3 out of 3 tasks 
1/3: running   
2/3: running   
3/3: running   
verify: Service converged 
````
Verify by running;
````
docker stack services stackdemo
````

**Play around with compose-files**


---

That's all you need to acutally build production-grade containerized applications.
In next chapter, we will try to build our features-rich example application with all the components we have learnt through.

---

