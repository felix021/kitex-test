#!/bin/bash

exec=kitex
exec=kitex-github@v0.5.0

$exec -module github.com/felix021/kitex-test -service kitex.felix021.server idl/api.thrift
