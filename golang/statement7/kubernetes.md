# Kubernetes 
- Kubernetes is an open source container orchestration engine for automating deployment, scaling, and management of containerized applications.
- Kubernetes provides you with a framework to run distributed systems resiliently. It takes care of scaling, failover for your application and provides deployment patterns.
  
## Node 
- A node is a machine (physical or virtual) on which kubernetes is installed. A node is a worker machine and this is were containers will be launched by kubernetes.

## Cluster 
- A cluster is a set of nodes grouped together. This way even if one node fails you have your application still accessible from the other nodes. Moreover having multiple nodes helps in sharing load as well.


## Components of Kubernetes cluster 
- A Kubernetes cluster has 2 main parts :
1. Control Plane - Brain
2. Worker Nodes - Machines running app
- In production environments, the control plane usually runs across multiple computers and a cluster usually runs multiple nodes, providing fault-tolerance and high availability.

### Control Plane
- Manges the overall state of the cluster. It makes global decisions about the cluster (scheduling, detecting and responding to cluster events)
1. `kube-apiserver`
- The core component server that exposes the Kubernetes HTTP API.
- Every interaction with the cluster goes through the API server.
- It is the entrypoint of kubernetes.
- request → API Server → cluster updated
- It is stateless — Does store anything itself. It validates requests and persists state to etcd.
2. `etcd (Cluster database)`
- Consistent and highly-available key value store for all API server data.
- Stores cluster state 
- Information about pods, nodes, Deployment, ReplicaSet, StatefulSet, DaemonSet,ConfigMaps,Secrets
3. `kube-scheduler`
-  Watches for newly created Pods with no assigned node, and selects a node for them to run on.
-  Factors considered while deciding: resource requirements,data locality, constraints,deadlines,affinities.
4. `kube-controller-manager`
- Responsible for keeping the cluster in desired state.
- Single binary that runs multiple independent control loops (controllers) in the same process. Each controller is responsible for one specific resource type or concern.
- Continuously watches cluster state and makes changes to match the desired state.
- Kubernetes works on declarative configuration ( you describe what you want and controller makes it happen)
- Suppose we need three replicas , but currently they are two, the controller sees this mismatch and creates and additional pod.
- Watches objects via API server -> Compare desired vs actual state -> Take action to fix difference -> Repeat forever
- This process is called reconciliation loop.
- There are many different types of controllers.
5. `Cloud Controller Manager`
- component that embeds cloud-specific control logic.
- The cloud controller manager lets you link your cluster into your cloud provider's API, and separates out the components that interact with that cloud platform from components that only interact with your cluster.
- The cloud-controller-manager only runs controllers that are specific to your cloud provider. 
- Not required in learning environments/within own PC.
- just like the kube-controller-manager, it combines several logically independent control loops into a single binary that you run as a single process. 


### Worker nodes 
1. `kubelet`
- The kubelet is the manager on each worker node.
- Talks to API server, starts and stops pods,run health checks.
2. `kube-proxy `
- kube-proxy is a network proxy that runs on each node in your cluster, implementing part of the Kubernetes Service concept.
- kube-proxy maintains network rules on nodes. These network rules allow network communication to your Pods from network sessions inside or outside of your cluster.
3. `Container runtime`
- Responsible for running containers. (containerd, CRI-O,any implementation of kubernetes CRI)
- It is responsible for managing the execution and lifecycle of containers within the Kubernetes environment.

## kubectl 
- kubectl is the command-line tool used to interact with a Kubernetes cluster.
- The kube control tool is used to deploy and manage applications on a kubernetes cluster, to get cluster information, get the status of nodes in the cluster.

## Minikube 
- Minikube is a tool that runs a single-node Kubernetes cluster locally on your machine.

## kubeadm 
- real cluster setup
- kubeadm tool helps us setup a multi node cluster with master and workers on separate machines. 
- kubeadm bootstraps and configures a Kubernetes cluster.
  
## Pods
- Pods are the smallest deployable units of computing that you can create and manage in Kubernetes.
- The containers are encapsulated into a Kubernetes object known as Pods. A Pod is a single instance of an application. 
- A Pod is a group of one or more containers, with shared storage and network resources, and a specification for how to run the containers. 
- Pods in a Kubernetes cluster are used in two main ways:
1. Pods that run a single container - most common, in this case  the Pod is a wrapper around a single container.
2. Pods that run multiple containers that need to work together- Pod can encapsulate an application composed of multiple co-located containers that are tightly coupled and need to share resources. These co-located containers form a single cohesive unit. (advanced use case)
- To scale up you create new Pods and to scale down you delete Pods. You do not add additional containers to an existing Pod to scale your application.

## ReplicaSet and ReplicationController
- In case of a single Pod running applications , if the pod crashes ,users will no longer be able to access the application. 
- To solve this, we require more than one instance or Pod running at the same time. 
- The ReplicationController helps us run multiple instances of a single Pod in the kubernetes cluster
thus providing high availability.
- ReplicationController ensures that the specified number of Pod are running at all times(Even if it’s just 1 or 100)
- Ensures that a pod or a homogeneous set of pods is always up and available.

- ReplicaController is deprecated, and ReplicaSet is the newer technology.

- A ReplicaSet's purpose is to maintain a stable set of replica Pods running at any given time. 
- Usually, you define a Deployment and let that Deployment manage ReplicaSets automatically.

## Deployments 
- A Deployment is a Kubernetes object that manages ReplicaSets and Pods, providing scaling, rolling updates,rollbacks,pause and resume rollout.
-  We can declare a new state of the pods by updating the PodTemplateSpec of the Deployment. A new ReplicaSet is created, and the Deployment gradually scales it up while scaling down the old ReplicaSet, ensuring Pods are replaced at a controlled rate. 
- Rolling Updates: Update application Pods gradually, one at a time, so users do not experience downtime during a new release.
- Rollback: Revert the Deployment to a previous stable version if the latest update causes errors or failures.
- Pause and Resume: Temporarily stop a Deployment rollout to make multiple changes, then resume so all updates are applied together safely.
- Scaling: Increase or decrease the number of running Pod replicas to handle more or less user traffic.

## StatefulSet 
- StatefulSet runs a group of Pods, and maintains a sticky identity for each of those Pods. This is useful for managing applications that need persistent storage or a stable, unique network identity.
- Each Pod keeps a fixed name, DNS identity, and persistent disk even after restart or rescheduling.
- They provide ordered deployment, scaling, and rolling updates, ensuring Pods start, update, and stop in a controlled sequence, which is important for databases.
- They are ideal for stateful applications like databases, Kafka, or Redis clusters that need consistent node identity and stored data across restarts.
- If an application is stateless and doesn’t need stable identity or ordered updates, a Deployment can be used.

## DaemonSet 
- A DaemonSet is a Kubernetes controller that ensures one copy of a Pod runs on every node (or selected nodes), automatically adding Pods when new nodes join and removing them when nodes leave.
- A DaemonSet defines Pods that provide node-local facilities. These might be fundamental to the operation of your cluster, such as a networking helper tool, or be part of an add-on.
- Typically use cases are : cluster storage daemon,logs collection daemon ,node monitoring daemon. All of these run on each node in the cluster.
- Deleting a DaemonSet will clean up the Pods it created.

## ConfigMaps
- A ConfigMap is a Kubernetes object used to store non-sensitive configuration data in key-value form, which Pods can use.
- Without ConfigMaps, configuration values must be hardcoded inside the application, so any config change requires code updates and redeployment, whereas ConfigMaps separate configuration from the app, allowing easy updates.

## Secret
- Secrets are Kubernetes objects used to store sensitive data like passwords, API tokens, or keys without putting them in application code or Pod specifications.
- Secrets are similar to ConfigMaps but meant specifically for confidential data.

## Services
- Service is a method for exposing a network application that is running as one or more Pods in your cluster.
- ClusterIP: Exposes a Service only inside the Kubernetes cluster, allowing internal Pods to communicate with each other.
- NodePort: Exposes a Service on a fixed port of every node’s IP address, allowing external users to access the application using NodeIP:Port.
- LoadBalancer: Exposes a Service externally by creating a cloud load balancer that distributes traffic to Pods, commonly used for production applications.

## Ingress
- Ingress is a Kubernetes object that controls how external traffic reaches Services, using rules based on domain name or URL path.

## Commands 
### Nodes
- `kubectl cluster-info`          # Show cluster info
- `kubectl get nodes`             # List all nodes
- `kubectl describe node <name>`  # Detailed node info
- `kubectl top nodes`            # real-time, CPU and memory usage for all nodes


### Pods
- `kubectl get pods`                     # List all pods
- `kubectl get pods -o wide`             # Show pods with node/IP details
- `kubectl describe pod <pod-name>`      # Detailed info about a pod
- `kubectl logs <pod-name>`              # View logs of a pod
- `kubectl exec -it <pod-name> -- COMMAND`  # Run a command inside a pod 
- `kubectl delete pod <pod-name>`        # Delete a pod


### ReplicaSet
- `kubectl get rs`                          # List ReplicaSets
- `kubectl describe rs <rs-name>`           # Detailed ReplicaSet info
- `kubectl scale rs <rs-name> --replicas=5` # Change number of replicas
- `kubectl delete rs <rs-name>`             # Delete a ReplicaSet


### Deployment
- `kubectl get deployments`                         # List deployments
- `kubectl describe deployment <name>`               # Detailed deployment info
- `kubectl apply -f app.yaml`                       # Create/update deployment from YAML
- `kubectl scale deployment <name> --replicas=5`     # Scale deployment
- `kubectl rollout restart deployment <name>`       # Restart all pods in a deployment
- `kubectl rollout status deployment <name>`        # Check rollout status
- `kubectl rollout undo deployment <name>`           # Rollback deployment
- `kubectl rollout history deployment <name>`           # Show rollout history of a deployment
- `kubectl delete deployment <name>`                # Delete deployment


### Service
- `kubectl get svc`                              # List services
- `kubectl describe svc <service-name>`          # Detailed service info
- `kubectl expose deployment <name> --port=80`   # Create service for a deployment
- `kubectl delete svc <service-name>`            # Delete a service


### General
- `kubectl apply -f file.yaml`        # Create or update resources defined in a YAML file
- `kubectl delete -f file.yaml`       # Delete resources defined in a YAML file
- `kubectl get all`                 # List most common resources 