#!/bin/sh

BINARY='/usr/local/bin'

echo "Building ip"
go build ip.go

echo "Installing ip to $BINARY"
install -v ip $BINARY

echo "Removing the build"
rm ip
