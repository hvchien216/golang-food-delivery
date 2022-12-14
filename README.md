
<h1 align="center">Welcome to Food Delivery [Golang] 👋</h1>

<p align="left">
<a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a>
<a href="https://www.mysql.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/mysql/mysql-original-wordmark.svg" alt="mysql" width="40" height="40"/> </a>
<a href="https://www.docker.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/docker/docker-original-wordmark.svg" alt="docker" width="40" height="40"/> </a>
<a href="https://www.nginx.com" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/nginx/nginx-original.svg" alt="nginx" width="40" height="40"/> </a>
</p>

## `✨ Prerequisites` ✈️️

- Go >= 1.19.2

## `🚀 Command` 

### `≈Build≈`
```bash
# build cross platform
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

```

### `≈Docker≈`
```bash
# Database
$ docker run --name mysql --privileged=true \
    -e MYSQL_ROOT_PASSWORD="1234" \
    -e MYSQL_USER="food_delivery" \
    -e MYSQL_PASSWORD="1234" \
    -e MYSQL_DATABASE="food_delivery" \
    -p 3306:3306 bitnami/mysql:5.7

#docker build
$ docker build -t food-delivery-image

# create a network
$ docker network create fd-delivery

# docker network connect
$ docker network connect fd-delivery mysql

# run 
$ docker run -d --name food-delivery \ 
    -e DBConnectionStr="food_delivery:1234@tcp(mysql:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local" \ 
    --network=fd-delivery \ 
    -p 3001:8080 \
    food-delivery-app
    
```

### `≈Install 🐳🐳 DOCKER 🐳🐳 Engine on Ubuntu≈` [Reference](https://docs.docker.com/engine/install/ubuntu/)

```bash
$ sudo apt-get update

$ sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

$ sudo mkdir -p /etc/apt/keyrings
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

$ echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
  
$ sudo apt-get update

$ sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin


```

### `≈Nginx as a Reverse Proxy≈`
<p align="left"><a href="https://www.nginx.com" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/nginx/nginx-original.svg" alt="nginx" width="40" height="40"/> </a></p>

```bash
$ docker run -d -p 80:80 -p 443:443 \
    --network=fd-delivery --name nginx-proxy \
    -e ENABLE_IPV6=true \
    --privileged=true \
    -v ~/nginx/vhost.d:/etc/nginx/vhost.d \
    -v ~/nginx-certs:/etc/nginx/certs:ro \
    -v ~/nginx-conf:/etc/nginx/conf.d \
    -v ~/nginx-logs:/var/log/nginx \
    -v /usr/share/nginx/html \
    -v /var/run/docker.sock:/tmp/docker.socker:ro \
    --label nginx_proxy jwilder/nginx-proxy
    
$ docker run -d --network=fd-delivery \
    -v ~/nginx/vhost.d:/etc/nginx/vhost.d \
    -v ~/nginx-certs:/etc/nginx/certs:rw \
    -v /var/run/docker.sock:/tmp/docker.socker:ro \
    --volumes-from nginx-proxy \
    --privileged=true \
    jrcs/letsencrypt-nginx-proxy-companion
```

## `🚀 Author`
👤 **Chien Anbs**
- Github: [@hvchien216](https://github.com/hvchien216)


### `Other...`
<p align="left">
<a href="https://letsencrypt.org/" target="_blank" rel="noreferrer"> <img src="https://cdn.iconscout.com/icon/free/png-256/letsencrypt-3521543-2944961.png" alt="cloudflare" width="40" height="40"/> </a>
<a href="https://www.cloudflare.com/" target="_blank" rel="noreferrer"> <img src="https://cdn.iconscout.com/icon/free/png-256/cloudflare-2752221-2285038.png" alt="cloudflare" width="40" height="40"/> </a>
<a href="https://www.namecheap.com/domains/#pricing" target="_blank" rel="noreferrer"> <img src="https://cdn.iconscout.com/icon/free/png-256/namecheap-283654.png" alt="namecheap" width="40" height="40"/> </a>
</p>