#!/bin/bash

set -e 

export GOPATH=$PWD/cpi-src

echo $0
echo `dirname $0`
echo $PWD
ls -lart 
make test
