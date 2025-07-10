#!/bin/bash

# start the database in the background
redis-server 1> log/redis 2>&1 & 

# start the backend in the foreground
go run main.go 

# kill the database
kill $!