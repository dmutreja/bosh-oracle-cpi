#!/bin/bash

set -e 

export GOPATH=$PWD/cpi-src
cd cpi-src
make test
