# syntax=docker/dockerfile:1.7-labs
FROM golang:1.23.2-bookworm

WORKDIR /app
COPY --exclude=node_modules --exclude=static/js/* --exclude=src/contracts . .

RUN apt-get update -y && apt install -y nodejs npm
RUN npm install --global yarn

WORKDIR src/frontend
RUN yarn && yarn build-and-copy
WORKDIR /app
