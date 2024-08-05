#!/bin/bash

echo "#@{{flag1}}@#" > /usr/share/nginx/html/flag
nginx -g "daemon off;"
