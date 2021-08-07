#!/bin/bash

set -e

#if [[ -z $@ ]]; then
#	echo "usage: $0 port1 port2 ... portN"
#	echo "  will spawn a server running on each of these ports"
#	exit 2
#fi

if [[ -z $1 ]]; then
	echo "usage: $0 network-urls-file"
	exit 2
fi

url_file=$1

node_program=./onion-node
ports=$(sed -e 's/#.*$//' -e 's/.*://' $url_file)
for port in $ports; do
	$node_program $port &
	pids="$pids $!"
done
echo "node process ids: $pids"


sleep .5
killcmd='kill'
killme="nope not yet"
while [[ $killme != $killcmd ]]; do
	echo -n "type '$killcmd' to kill the nodes: "; read killme
done

kill $pids
