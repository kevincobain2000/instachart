#! /bin/bash

pm2 delete 3001
pm2 start --name 3001 "./main"