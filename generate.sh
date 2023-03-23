#!/bin/bash

exec=kitex
exec=kitex-github@v0.4.4

$exec -module github.com/felix021/kitex-test -service kitex.felix021.server idl/api.thrift
