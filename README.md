# Anika_Internship
This is a containerised URL shortener application. The frontend , backend and redis database run as different conatiners and communicate via kubernetes services. It is run locally on minikube 


Steps to run 
minikube start 

minikube status 

podman build -t localhost/url-shortener-backend:1.1 .

cd k8s/frontend

podman build -t localhost/url-shortener-frontend:1.9 .

podman save -o url-shortener-frontend.tar localhost/url-shortener-frontend:1.9
podman save -o url-shortener-backend.tar localhost/url-shortener-backend:1.1

minikube image load url-shortener-frontend.tar
minikube image load url-shortener-backend.tar

minikube addons enable ingress
kubectl apply -f k8s/redis-deployment.yaml
kubectl apply -f k8s/redis-service.yaml
kubectl apply -f k8s/backend-deployment.yaml
kubectl apply -f k8s/backend-service.yaml
kubectl apply -f k8s/frontend-deployment.yaml
kubectl apply -f k8s/frontend-service.yaml
kubectl apply -f k8s/frontend-ingress.yaml

On another terminal keep this running:
kubectl port-forward -n ingress-nginx svc/ingress-nginx-controller 8080:80

Create mapping 
sudo nano /etc/hosts
127.0.0.1 short.local

Open in browser
http://short.local:8080
