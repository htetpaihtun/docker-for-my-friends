#### Hello.! In this chapter, we will start our first step into application containerization journey. I will try to explain how containers came in, in the first place.

**Disclaimer:**
Since I only got into IT industry in early 2016-2017, I had no clue about what the industry was like before.
So, below statements are my impression on earlier days and what I've read about earlier days and therefore not my own experience. 
Anyway, these are what I was told. 

Back to the earlier days in software development industry, they could only run one application per server 
as Operating Systems weren't mature enough to implement way to reliably run multiple applications on same server.
If a company demand a new application, they had to buy extra hardwares for server first and then installed OS stuff and did configuration on them.
It was horrible to do them plus both requirement and resource management on per server basis and application basis.

Then, Virtual Machines technologies came in and now they could host many applications on a single server.
VMs can separate a server into little blocks of servers with their own OS and hardware space (virutally) according to your needs.

Now everyone's happy.

Things got better but could they be even more BETTER.?

VMs are not perfect, they require their own OS (along with the cost to maintain them), 
their own hardware limitations,they are slow to boot, lack of portability and more..

Here comes our main hero,containers. 

The container are similar to VMs but they do more of software virtualization than hardware virtualization and therefore don't need their on OS and instead share their host OS as well as hardwares.

Guess what. They are faster and portable too!

They first came in form of Linux containers and later Windows and other adopted it as well too.

People and companies were so happy about containers and they invested a lot of resource into containerization technologies.
The major technologies that leads to modern containers as nowsaday were kernel namespaces, control groups, union filesystems. 

These wouldn't be possible with such great community support.         
Thanks to the indivials and organizations participated and of course Docker itself. 

![type1 hypervisors and type2 hypervisors](http://www.techplayon.com/wp-content/uploads/2020/08/hypervisor.png)

_Figure 1.1 Type 1 Hypervisor and Type 2 hypervisor_

![difference between docker and vm](https://vmarena.com/wp-content/uploads/2018/08/DOCK02.png)

_Figure 1.2 Docker comparison over VMs_

#### In later chapter, we will go further into Docker and what makes them so special.


[Chapter-2 Docker Overview](https://github.com/htetpaihtun/docker-for-my-friends/tree/main/chapter-2#in-this-chapter-you-will-learn-background-of-dockerinc-and-your-brief-answers-to-your-very-first-questions-about-docker-starting-with-why-docker-what-docker-how-docker-and-more)
