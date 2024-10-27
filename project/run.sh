#!/usr/bin/env bash
# Created by Petr Lozhkin

 docker build -t server:latest . && \
 docker run -p 3002:3002 server:latest
