# Docker Tutorial

## Installation

##### LINUX

[Reference](https://docs.docker.com/engine/install/ubuntu/ "Linux Installation Guide")

###### Set up the repository
> Update the apt package index and install packages to allow apt to use a repository over HTTPS:

 ```
 sudo apt-get update
 sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
  ```
    
###### Add Docker’s official GPG key:

```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
```

###### Make sure to update apt

```
sudo apt-get update
```

###### Now Install

```
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

###### Validate Installation

```
sudo docker run hello-world
```

##### Install App Image

[Reference](https://hub.docker.com/search?q=&type=image "App Images")

> Choose your App Environtment <br>
> Example below

```
sudo docker pull golang:1.16.5
```


## Setup Docker to your app

> **Rules!!** <br>
> Dockerfile should be inside your app folder


#### Create Dockerfile

<img src="tree.png" alt="tree"  />

##### Inside Dockerfile

```
#declare image app environtment
FROM golang:1.16.5

#container directory setup
WORKDIR /container/directory/workdir/Application

#copy <source> to <container directory>
COPY . /container/directory/workdir/Application

#run the application in the container
CMD [ "./application" ]
```

## Build Docker

```
sudo docker build -t application-name .
```


## Run Docker

```
sudo docker run
```

## Check Image

```
sudo docker image ls
```

## Check Image Running

```
sudo docker ps
```

## Stop Image Running

```
sudo docker stop @-NAMES
```

## Docker hub Setup

> Follow in order


```
sudo docker login

sudo docker build -t dockerid/repository

sudo docker run --name yourname  dockerid/repository

sudo docker push dockerid/repository

```
