#! /bin/bash


cd frontend
npm install
npm run build
cd ..

go build main.go
pm2 delete 3001
pm2 start --name 3001 "./main"

### or use this
#curl -sLk https://raw.githubusercontent.com/kevincobain2000/instachart/master/install.sh | sh
#pm2 delete 3001
#pm2 start --name 3001 "./instachart"