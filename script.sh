#!/bin/bash

apt-get update
apt-get install -y golang
apt-get install -y postgresql postgresql-contrib

service postgresql start
sudo -u postgres psql -c "CREATE USER postgres WITH PASSWORD '123456';"
sudo -u postgres createdb ejerciciogo
psql -h localhost:5433 -u postgres -p 123456 ejerciciogo < tribalgo/queries.sql

go mod tidy

go build -o tribalgo/main tribalgo/*.go

cd tribalgo
./main