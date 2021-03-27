#!/usr/bin/env bash

#dependencies: protoc protoc-gen-go protoc-gen-micro
# protoc-gen-go v1.3.5
protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. finalcache.proto