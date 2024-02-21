#!/bin/bash

set -x

kitex-v0.9.0-rc7 -module github.com/felix021/dubbo-p2 -protocol Hessian2 -hessian2 java_extension api2.thrift
