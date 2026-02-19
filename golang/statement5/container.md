# Containers

## What is a Container ?
- A container is a standard unit of software that packages up code and all its dependencies into a single bundle so the application runs quickly and reliably from one computing environment to another.
- A container image is a lightweight, standalone, executable package of software that includes everything needed to run an application.
  
### Containers typically include the following :
-  `Application code` - Actual code files (Go binary/Node.js app/Python sripts)
-  `Runtime Environment` - Software needed to run the code (Go runtime/Node.js runtime/ Python Interpreter)
-  `Libraries and Dependencies` - All packages required by the application (npm packages/Go modules)
-  `Configuration Files` - Settings used to run the application (.env files/ config.json)
-  `System Tools` - Utilities required by the application (bash/ apt packages)
-  `Filesystem` - A small isolated filesystem that includes `/app` folder,binaries,libraries,configs
-  
### Containers do not include :
- `Full operating system` - Containers share the host OS Kernel, which makes them lightweight

## Advantages of Containers 
- `Portability` Containers package all dependencies, enabling applications to run consistently across various environments. Solves the 'it works on my machine' problem.
- `Scalability` Containers can be easily scaled up or down, enabling efficient handling of dynamic workloads. 
- `Microservices` Containers are just the right fit for deploying microservices, where each service runs in its own container.
- `Faster Deployment` Due to their lightweight nature, containers can be deployed quickly, enhancing development and CI/CD processes.
- `Consistency` Containers ensure that the application runs the same way across development, testing, and production environments.
- `Efficiency` Containers use fewer resources than VMs, allowing for more apps on the same server.


## Containers vs Virtual Machines

| Feature        | Containers                       | Virtual Machines                       |
| -------------- | -------------------------------- | -------------------------------------- |
| What they are  | Package app + dependencies       | Full computer inside a computer        |
| OS             | Share host OS kernel             | Each VM has its own OS                 |
| Size           | Small (MBs)                      | Large (GBs)                            |
| Startup Time   | Seconds                          | Minutes                                |
| Speed          | Faster                           | Slower than containers                 |
| Resource Usage | Use less CPU & RAM               | Use more CPU & RAM                     |
| Isolation      | Process-level isolation          | Full OS-level isolation                |
| Portability    | Very portable                    | Portable but heavier                   |
| Example Tools  | Docker, Podman                   | VMware, VirtualBox                     |
| Best Use       | Microservices, CI/CD, cloud apps | Running different OS, strong isolation |


## Container files 
A Containerfile (same as Dockerfile) is a text file that contains instructions to build a container image.

### Image vs Container
- A container is a running instance of an image.  
- An image is a blueprint used to create many containers.  
- The relationship between images and containers is similar to classes and objects in programming.


## Common Containerfile Instructions

| Instruction | Description                                                 |
| ----------- | ----------------------------------------------------------- |
| FROM        | Specifies the base image to start building from.            |
| WORKDIR     | Sets the working directory inside the container.            |
| COPY        | Copies files and folders from host to container.            |
| RUN         | Executes commands while building the image.                 |
| ENV         | Sets environment variables inside the container.            |
| EXPOSE      | Documents which port the container will listen on.          |
| ENTRYPOINT  | Sets the main executable that always runs in the container. |
| CMD         | Sets default command/arguments for the container.           |
| LABEL       | Adds metadata  to the image.                                |
| STOPSIGNAL  | Defines the signal used to stop the container gracefully.   |


## Containerflow
- Container flow is the step-by-step process of how a container is created, run, and used.
  
1. Create Containerfile that defines how to build image
2. Build Container Image (using Docker/Podman)
3. Run Container
4. Test Container ( checks logs and functionality)
5. Push image to registry
6. Deploy container ( pull image and run the container in production using Kubernetes/Cloud services)

## Conatinertools
- Container tools are software used to build, run, manage, and deploy containers.

1. `Container engines` - Used to build and run containers (Docker, Podman)
2. `Container Orchestration tools` - Used to manage many containers (Kubernetes, Docker Swarm, OpenShift)
3. `Container Registries` - Used to store container images (Docker Hub, GitHub Container Registry)


## Commands 

### Images 
- `podman images` - List all images 
- `podman rmi <image name>` - Delete an image (Delete all dependent containers to remove image)
- `podman image prune` - Remove unused images 
- `podman build -t <image_name>:<version> .` - Build image 
- `podman build -t <image_name>:<version> . --no-cache` - Build image without cache

### Containers
`podman pull <image_name>` _ Pull image from registry
`podman run <image_name>` - Run a container
`podman run -d <image_name>` - Run container in background
`podman run -p <host_port>:<container_port> <image_name> ` - Run container with port mapping
`podman run --name <container_name> <image_name>` - Run container with name
`podman ps` - List running containers
`podman ps -a`  - List all containers (including stopped)
`podman stop <container_id_or_name>` - Stop a container
`podman start <container_id_or_name> ` - Start a stopped container
`podman restart <container_id_or_name>` - Restart a container
`podman rm <container_id_or_name>` - Remove a container
`podman logs <container_id_or_name>` - View container logs
`podman logs -f <container_id_or_name>` - Shows real time logs 



