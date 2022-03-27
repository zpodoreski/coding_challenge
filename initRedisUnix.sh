#!/bin/bash
sudo apt-get install docker-engine -y

sudo service docker start

sudo docker pull redis

sudo docker run --name my-redis -p 6379:6379 -d redis



