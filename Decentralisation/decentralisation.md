===================================================================================================================
4.2 
Kubernetes and Docker are two of the major players in container orchestration.


Swarm 
======
Docker Swarm is an open-source container orchestration platform and is the native clustering engine for and by Docker. It consists of Docker Engines deployed on multiple nodes.  So swarm consists of multiple docker nodes. 

There are two kind of nodes - Manager node and worker Nodes. Both has its own role and together perfors the comntainer orchestration. 
Manager nodes perform orchestration and cluster management. Worker nodes receive and execute tasks from the manager nodes.




Kubernetes
===========
Kubernetes is an open source container orchestration platform developed by Google. It is production ready, enterprise grade and has auto heal capabilities. Auto heal capabilities allows kubernetes to manage auto-scale, auto-replicate and auto-restart features.  It is modular in design and can be utilized for any architecture deployment. 
Thus, Kubernetes automates the deployment, scaling, maintenance, scheduling and operation of multiple application containers across clusters of nodes. 

With Kubernetes, containers run in pods. A pod is a basic unit that hosts one or multiple containers, which share resources and are located on the same physical or virtual machine. For each pod, Kubernetes finds a machine with enough compute capacity and launches the associated containers. A node agent, called a Kubelet, manages pods, their containers and their images. Kubelets also automatically restart a container if it fails.




hyperledger fabric on single host
-----------------------------------
Hyperledger Fabric rely on docker based architecture. Each of its component has its onw image and docker run them on its own container. Each container is isolated env and thus they do not have any visibility of other containers. To make them interact with each other, we define the network  on that node and let all container join this network. Thus, single network provides them the oppurtunity to be discoverable by other containers..  These are the configuration details captured in the yaml file. 



hyperledger fabric on multiple hosts - Docker swarm
---------------------------------------------------
When containers are deployed on multiple hosts then these container can't communicate with each other. We need to define a single network spreding multiple hosts and let all the container spread on these hosts join this global network fot them to be discoverable. Ideally, we need a container orchstartion service which can be utilised to oversee these configuration. 
As both Kubernetes and Swarm are container orchestartion platorm, we can deploy hyperledger fabric on multiple hosts and use these two to connect all the containes on these hosts. 


Since all the hyperledger fabric images are docker images, its best to use docker swarm to manage the container orchestration. 


hyperledger fabric on multiple hosts - Kubernetes
-------------------------------------------------
Kubernetes is ideal for the following reasons:

It's easy to achieve high availability (HA) with Kubernetes.
Fabric is built into container images, and Kubernetes is useful in orchestration, scaling, and managing containers.
Kubernetes supports multi-tenancy, making it possible to develop and test blockchain applications.



links 
=============

Fabric and Swarm 
-----------------
https://github.com/swarmpit/swarmpit (Lightweight Docker Swarm management UI )
https://github.com/tubackkhoa/hyperledger-swarm
https://github.com/docker/swarm (Swarm: a Docker-native clustering system)
https://github.com/suddutt1/fabricswarm1 (First Docker Swarm With 3 machines and balance transfer)




Kubernetes and Swarm
---------------------
https://github.com/kubernetes/kompose
https://github.com/kubernetes/kubernetes
https://github.com/kubernetes/client-go
    - Go client for Kubernetes.
https://github.com/kubernetes/dashboard 
    - General-purpose web UI for Kubernetes clusters
https://github.com/kubernetes/kops
    - Kubernetes Operations (kops) - Production Grade K8s Installation, Upgrades, and Management



===================================================================================================================
4.3 IBM cloud and Hyperledger Fabric 

- IBM cloud provides  'IBM Cloud Kubernetes Service API'. It combines Docker and Kubernetes to deliver powerful tools to automate the deployment, operation, scaling, and monitoring of containerized applications over a cluster of independent compute hosts by using the Kubernetes APIs. 
- It also provides 'IBM Blockchain Platform', which is software-as-a-service offering from IBM Cloud. It enables network members to quickly get started developing and easily move to a collaborative environment. It is the only fully integrated blockchain platform designed to accelerate the development, governing, and operation of a multi-institution business network. 
===================================================================================================================
4.4

Hyperledger Cello
==================
Hyperledger Cello is a blockchain provision and operation system, which helps manage blockchain networks in an efficient way. It is one of the Hyperledger projects hosted by The Linux Foundation.
 
 
Aim
===
Hyperledger Cello aim is to bring the on-demand “as-a-service” deployment model to the blockchain ecosystem to reduce the effort required for creating, managing and terminating blockchains. Hyperledger Cello was initially contributed by IBM, with sponsors from Soramitsu, Huawei and Intel.
 
Using Cello, everyone can easily:
    Build up a Blockchain as a Service (BaaS) platform quickly from scratch.
    Provision customizable Blockchains instantly, e.g., a Hyperledger fabric network v1.0.
    Support for customized blockchains request (e.g., size, consensus) — currently, there is support for Hyperledger Fabric
    Maintain a pool of running blockchain networks on top of baremetals, Virtual Clouds (e.g., virtual machines, vsphere Clouds), Container clusters (e.g., Docker, Swarm, Kubernetes).
    Check the system status, adjust the chain numbers, scale resources... through dashboards.
 
 
More Info
=========
GitHub - https://github.com/hyperledger/cello
 
Main doc - https://hyperledger-cello.readthedocs.io/en/latest/
user dashboard for Blockchain - https://github.com/hyperledger/cello/tree/master/user-dashboard
https://hyperledger-cello.readthedocs.io/en/latest/dashboard/
 
 
Architecture design  https://hyperledger-cello.readthedocs.io/en/latest/arch/


Support 
=======
We can touch base with Hyperledger members to get needed support on Hyperldger Cello. Below are the core memebers that together developed this project. 
IBM 
Soramitsu
Huawei
Intel 

We can filter out all the members from the site - https://www.hyperledger.org/resources/vendor-directory

===================================================================================================================