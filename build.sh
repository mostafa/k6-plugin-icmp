#!/bin/bash

# This fetches the k6-plugin-icmp from GitHub and
# builds it in your $GOPATH/src/github.com/mostafa/k6-plugin-icmp

go build -buildmode=plugin -ldflags="-s -w" -o icmp.so github.com/mostafa/k6-plugin-icmp
