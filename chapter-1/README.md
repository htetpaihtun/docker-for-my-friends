#### Hello. In this chapter we will start our first step into application containerization journey. I will try to explain how containers came in, in the first place.

**Disclaimer:**
Since I only got into IT industry in early 2016-2017, I had no clue about what the industry was like before.
So, below statements are my impression on earlier days and what I've read about earlier days and therefore not my own experience. 
Anyway, these are what I was told. 


**Let's get started with a story uhm.. I mean history ye.. uhmm... ok whatever you'd like to call.**

Back to the earlier days in software development industry, they could only run one application per server as Operating Systems (mostly Linux, Windows) weren't mature enough to implement way to reliably run multiple applications on same server.

If a company demand a new application, they had to buy extra hardwares for server first and then installed OS stuff and did configuration on them.

It was horrible to do them plus both requirement and resource management on per server basis and application basis.

Then, Virtual Machines technologies came in and now they could host many applications on a single server.

VMs can separate a server into little blocks of servers with their own OS and hardware space (virutally) according to your needs.

Now everyone's happy.

Things got better but could they be even more BETTER.?

VMs are not perfect, they require their own OS (along with the cost to maintain them), their own hardware limitations,they are slow to boot, lack of portability and more..

Here's come our main hero,containers. 

The container are similar to VMs but they do more of software virtualization than hardware virtualization and therefore don't need their on OS and instead share their host OS as well as hardwares.

Guess what. They are faster and portable too!

They first came in form of Linux containers and later Windows and other adopted it as well too.

People and companies were so happy about containers and they invested a lot of resource into containerization technologies.

The major technologies that leads to modern containers as nowsaday were kernel namespaces, control groups, union filesystems. 

These wouldn't be possible with such great community support.

Thanks to the indivials and organizations participated and of course Docker itself. 

What a time to be alive. 

#### In later chapter, we will go further into Docker and what makes them so special.
