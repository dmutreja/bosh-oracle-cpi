#!/bin/bash

set -e 

export GOPATH=$PWD/cpi-src

make test
