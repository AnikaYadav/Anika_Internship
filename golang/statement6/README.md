# Trie gRPC Kubernetes Project

## Overview

This project implements a **Trie-based word store** exposed via **gRPC**, containerized with **Podman**, and deployed on **Kubernetes (Minikube)**.

It supports the following operations:

* **Add** a word
* **Remove** a word
* **Check** if a word exists
* **List** all words

The project demonstrates:

* Go + gRPC + Protocol Buffers
* Interactive client inside Kubernetes

## Commands
- podman build -t localhost/trie-server:v1 -f server/Dockerfile .
- podman build -t localhost/trie-client:v1 -f client/Dockerfile .
- podman save -o trie-server-v1.tar localhost/trie-server:v1
- podman save -o trie-client-v1.tar localhost/trie-client:v1
- minikube image load trie-server-v1.tar
- minikube image load trie-client-v1.tar
- minikube image ls | grep trie   # should exist
- kubectl exec -it trie-client -- /app/client # start the client   
