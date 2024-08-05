#!/bin/bash

# docker in docker in docker has a problem with mounts, hacking around it
echo "${FLAG}" > /usr/share/nginx/html/flag
nginx -g "daemon off;"
