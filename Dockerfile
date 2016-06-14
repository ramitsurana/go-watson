FROM node:slim

MAINTAINER Ramit Surana <ramitsurana@gmail.com>

RUN apt-get update && install wget curl git

#Installing Nodejs version 6
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash -
RUN apt-get install nodejs -y

#Installing NPM
RUN apt-get update
RUN apt-get install npm -y

#Installing yeoman
RUN apt-get update
RUN npm install -g generator-angular
RUN npm install --global yo

#Installing Watson
RUN npm install -g generator-watson

#Cleaning cache
RUN apt-get autoremove 
RUN apt-get clean

CMD ["/bin/bash"]
