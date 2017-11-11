#!/bin/bash

set -e 

export GOPATH=$PWD/cpi-src
cd cpi-src/github.com/oracle/bosh-oracle-cpi
make test
