#!/bin/bash

python3 -m grpc_tools.protoc -I../idl --python_out=. --pyi_out=. --grpc_python_out=. ../idl/api.proto
