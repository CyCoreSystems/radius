#!/bin/sh

mkdir -p db
chmod 777 db
docker build -t freeradius freeradius
docker run --name=freeradius -v $PWD/db:/var/log/freeradius --rm --net=host -i -t freeradius 
