So far we have learnt about how to build our multi-container applications from one 'docker-compose.yaml' file.
We will now try to do it one multiple nodes (machines) and scale even larger.
Many people tends to stay away from Docker Swarm and Docker Stack because they think those are for multi-nodes management only.
In facts, Docker Swarm mode is fully funtional even if you're using only one node and Docker Stack has additional features that Docker Compose doesn't have.
So, Docker Compose might be good for local test devlopments but Docker Stack can go production. 

---

### 7.1 Docker Swarm

Docker Swarm is two main things:
1. An enterprise-grade secure cluster of Docker hosts
2. An engine for orchestrating microservices apps

On the clustering front, Swarm groups one or more Docker nodes and lets you manage them as a cluster. 
Out-of-the-box, you get an encrypted distributed cluster store, encrypted networks, mutual TLS, secure cluster join
tokens, and a PKI that makes managing and rotating certiﬁcates a breeze. 
You can even non-disruptively add and remove nodes.

On the orchestration front, Swarm exposes a rich API that allows you to deploy and manage complex microservices apps with ease. 
You can deﬁne your apps in declarative manifest files and deploy them to the Swarm with native Docker commands. 
You can even perform rolling updates, rollbacks, and scaling operations.
Again, all with simple commands.


#### 7.1.1 Swarm Mode 

  




---


























