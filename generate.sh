#!/bin/bash
protoc protos/*.proto --go_out=protos/  -I protos/  --proto_path=protos/
