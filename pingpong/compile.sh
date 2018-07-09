#!/bin/bash

protoc pingpong.proto --go_out=plugins=grpc:.
