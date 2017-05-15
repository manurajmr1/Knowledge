<h2>How to containerize a php web application with php, apache and mysql using Docker</h2>

<i>Some basic knowledge

DOCKER

Installation on debian


http://stackoverflow.com/questions/38728693/cannot-install-docker-on-debian-jessie or

Install from a package

From packages and list find latest .deb inside folder pool- https://docs.docker.com/engine/installation/linux/debian/#install-from-a-package 

Learn Docker - https://www.youtube.com/watch?v=YFl2mCHdv24

Docker compose - https://www.youtube.com/watch?v=Qw9zlE3t8Ko

https://us14.proxysite.com/process.php?d=YHYtpqAqwitm51Q5H02AHHdLxdxzcbz4N6wCMuzxuNjbjsvZpTuFOm72Jxq3YQ1U&b=1&f=norefer

DOcker Swarm - https://docs.docker.com/engine/swarm/#swarm-mode-key-concepts-and-tutorial

Dockerfile for Go - https://www.youtube.com/watch?v=-IoWzON4IRA, https://blog.golang.org/docker

For learning - https://katacoda.com 


<b>Some basic commands - </b>

docker ps(list docker containers) 

Steps for running a docker container and reuse it

Build  with name(ehr) - sudo docker build -t ehr .

(-p host:container)

Run with a custom name - sudo docker run -p 9001:80 -d  --name ehr -v /var/www/html/ehr/:/var/www/html/ ehr

After that start/stop using - sudo docker start ehr 

Then loginto container bash - sudo docker exec -it ehr bash

And install whatever u want…..apt-get update -y && apt-get install package-name -y before installing

</i>

<b>For composing & installing</b>

<i>reference - http://tech.osteel.me/posts/2017/01/15/how-to-use-docker-for-local-web-development-an-update.html</i>

sudo docker-compose up --build (each time a new package is added)

sudo docker-compose up
