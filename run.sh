#!/bin/sh

set -e
go get
go build -x -o ./test/daemon/testdaemon ./test/daemon
go build -x -o runsit
./runsit --config_dir=config $@
