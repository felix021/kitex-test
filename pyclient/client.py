#!/usr/bin/python3
from __future__ import print_function

import logging

import grpc
import api_pb2
import api_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    print("Will try to send request ...")
    with grpc.insecure_channel('localhost:8888') as channel:
        stub = api_pb2_grpc.EchoStub(channel)
        response = stub.Echo(api_pb2.Request(message='hello world from pyclient'))
    print("client received: " + response.message)


if __name__ == '__main__':
    logging.basicConfig()
    run()