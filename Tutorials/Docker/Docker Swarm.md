# Docker Swarm Tutorial

#### Reference
https://www.youtube.com/watch?v=3-7gZS4ePak

## Important things in docker-compose.yml
----------------
> No build feature for docker swarm means you need to use Image to run
`image: ngmedina14/ordering-system`

#### yaml version
the host port will only work if the yaml version is 3.5 or latest

```yaml
version: "3.5" 
````

#### Port configuration

##### You dont need to expose mysql port. they are living in a single instance/pc
```yaml
  mysql:
    image: mysql/mysql-server:8.0.26
    #ports:
    #  - "3306:3306"
```

##### Make the Web app the _Host Port_ and published port

```yaml
ports:
      - target: 8080
        published: 8080
        mode: host # the most important thing about docker swarm /Means Select 1 service 
``` 

##### Make sure to change Database Host to the Service Name

Use this command
`docker service ls`

```yaml
    environment: 
      - MYSQL_USER=root
      - MYSQL_ROOT_PASSWORD=NeilGwapo100%
      - MYSQL_ROOT_HOST=orderingsystem_mysql #Change this based on the mysql Service Name
      - MYSQL_DATABASE=ordering-system
```

##### Replication of Service/ Running Application to minimize downtime

```yaml
deploy:
      replicas: 2
      update_config:
        parallelism: 1 #Replace 1 task at a time
        delay: 10s
```

```yaml
version: "3.5"  # optional since v1.27.0
services:
  mysql:
    image: mysql/mysql-server:8.0.26
    #ports:
    #  - "3306:3306"
    environment: 
      MYSQL_USER: "root"
      MYSQL_ROOT_PASSWORD: "NeilGwapo100%"
      MYSQL_ROOT_HOST: "%"
      MYSQL_DATABASE: "ordering-system"
    volumes: 
      #- /var/lib/mysql:/var/lib/mysql #Docker Play Setup Only
      - MysqlBindVolume:/var/lib/mysql
    #restart: always
  web:
    # build: 
    #   context: . #Dockerfile
    # #   #dockerfile: Dockerfile.local
    #privileged: true
    image: ngmedina14/ordering-system
    deploy:
      replicas: 2
      update_config:
        parallelism: 1 #Replace 1 task at a time
        delay: 10s
    environment: 
      - MYSQL_USER=root
      - MYSQL_ROOT_PASSWORD=NeilGwapo100%
      - MYSQL_ROOT_HOST=orderingsystem_mysql
      - MYSQL_DATABASE=ordering-system
    ports:
      - target: 8080
        published: 8080
        mode: host

    volumes:
      - WebBindVolume:/go/src/github.com/OrderingSystem
    #restart: unless-stopped
    depends_on: 
      - mysql
volumes:
  WebBindVolume:
    driver_opts:
        type: none
        device: ${PWD} #/home/neil/go/src/github.com/ngmedina14/OrderingSystem
        o: bind
  MysqlBindVolume:
    driver_opts:
        type: none
        device: ${PWD}/data #/home/neil/go/src/github.com/ngmedina14/OrderingSystem
        o: bind
    
```

# Running Docker Swarm

## Join/ Create swarm init first
`docker swarm init #host pc`
###### if practicing swarm in playgroud use this
`docker swarm init --advertise-addr 127.0.0.1 #play with docker`

## Create a node
>Hint! replace the <nodeappname> to appname .. mysql host name is dependent to the appname<br>
> Ex. orderingsystem<br>
> then my  `- MYSQL_ROOT_HOST=_orderingsystem_mysql_`
  
###### Create node
`docker stack deploy -c docker-compose.yml <nodeappname>`

###### Service Status
`docker service ls`

-----------------------
  
###### Start Over Node / Remove Node
`docker stack rm <nodeappname>`

###### Start Over Swarm / Leave Swarm
`docker swarm leave -f`
  
###### Scale Application / Replica / Load balancer
`docker service scale <nodeapp_service>=<replica count>`
`docker service scale orderingsystem_web=3`

  
###### Run a docker Vizualizer For Visual Monitoring of Nodes
  
> Make sure you run this on the Swarm Manager VM/PC/Host

  
`docker service create --name=viz --publish=8000:8080/tcp --constraint=node.role==manager --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock dockersamples/visualizer`
  
-------------------
  
# Swarm
  
###### Monitor Swarm Network
`docker node ls`
  
###### Drain Swarm / Remove nodes in the VM/PC and place it to other VM/PC
  
`docker node update --availability=drain <VM/PC Name`
  
                                         
###### Active Swarm / Activate a previously Drain Swarm / Nodes can now run in the VM/PC

`docker node update --availability=active <VM/PC Name>` 

> Force the services or nodes go back and run to the VM/PC assigned originaly

  
`docker service update --force <nodeapp_service>`
                                         
---------------------  
  
# Service
 
###### Monitor Service Application / App Replica
`docker service ps <nodeapp_service>`
`docker service ps orderingsystem_web`
  
