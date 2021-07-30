#!/bin/bash

set -e

if [ -z $@ ]; then
	echo "usage: $0 port1 port2 ... portN"
	echo "  will spawn a server running on each of these ports"
	exit 2
fi

server_program=./server
ports=$@
for port in $ports; do
	$server_program $port &
done

sleep .5
killme="nope not yet"
while [[ $killme != "kill" ]]; do
	echo -n "type \"kill\" to kill the servers: "; read killme
done

killall $server_program
