# Kubernetes Tutorial

## Installation (Linux)

#### Minikube
https://minikube.sigs.k8s.io/docs/start/

curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64

sudo install minikube-linux-amd64 /usr/local/bin/minikube

#### Kubectl
https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/#install-kubectl-binary-with-curl-on-linux

`curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"`

`sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl`

-------------------------------

## Common Commands


##### Check nodes

`kubectl get nodes`

##### Check Service info

`kubectl get services`

##### Create a POD
`kubectl create deployment deployName --image=imageName `

##### check the deployed pods
`kubectl get deployment`
`kubectl get pod`

##### Remove/Delete the deployed app/pod
`kubectl delete deployment deployName`

##### edit kubernetes config file
`kubectl edit deployment`

##### Check Log of the application/pods/image
`kubectl logs deployName-someHash-someID` #check for kubectl get pod for deployNames


##### Apply edited kubernate config file
`kubectl apply -f /path/conf.yml`

##### Start a Minikube /create a VM for kube manager
`minikube start`

##### Delete a node
`kubectl delete -f k8s/`

##### Delete the Minikube / VM
`minikube delete`

##### Stop a Minikube / stop the process of the VM
`minikube stop`

