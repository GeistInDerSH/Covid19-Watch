#!/bin/bash

go build -o covidwatch main.go

touch output
( while true; do ./covidwatch > output; sleep $((60*60*24)); done) &

sleep 20

socat tcp-l:80,reuseaddr,fork system:"echo HTTP/1.0 200; echo Content-Type\: text/plain; echo; cat output"
