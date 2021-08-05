#!/bin/bash

set -e

# TODO: read portnumbers / urls of all the nodes from a file
#  have all the nodes also read from the urls file so the network is a fully connected graph
#    or do they even need that? if the client just has a list of all the nodes in the network, it can come up with its own path and then just have to send them to each other as general http URLs, right?

if [ -z $@ ]; then
	echo "usage: $0 port1 port2 ... portN"
	echo "  will spawn a server running on each of these ports"
	exit 2
fi

node_program=./onion-node
ports=$@
for port in $ports; do
	$node_program $port &
	pids="$pids $!"
done
echo "node process ids: $pids"


sleep .5
killcmd='kill'
killme="nope not yet"
while [[ $killme != $killcmd ]]; do
	echo -n "type '$killcmd' to kill the servers: "; read killme
done

kill $pids
