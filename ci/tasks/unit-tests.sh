#!/bin/bash

set -e 

export GOPATH=$PWD/cpi-src

echo $PWD
echo ls -lart 
make test
