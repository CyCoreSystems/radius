From ubuntu:trusty
MAINTAINER Elliott Ye

# Forked from https://github.com/catatnight/docker-freeradius/commit/b1dee571cfd792fa8acf054d461d5691e685938f

# Set noninteractive mode for apt-get
ENV DEBIAN_FRONTEND noninteractive

# Upgrade base system packages
RUN apt-get update

RUN apt-get -y install make # required for freeradius-config

RUN apt-get -y install software-properties-common

RUN add-apt-repository ppa:freeradius/stable-3.0

RUN apt-get update

RUN apt-get -y install freeradius 

# Add files
ADD install.sh /opt/install.sh
RUN chmod +x /opt/install.sh

COPY default /etc/freeradius/sites-enabled/

# Run
CMD /opt/install.sh;/usr/sbin/freeradius -f -Xxx
