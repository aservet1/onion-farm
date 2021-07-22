# onion-farm
onion network on one server with mocroservices (just server processes running, not docker container or anything) as the nodes

<img src='onion-server.png'/>

Here's a really good picture describing what I have in mind

So we have one server that routes messages, but when they exit the server we don't know who they came from, because in the routing the message goes through a bunch of onion nodes. Each node will be a process on the machine, running as a (probably Java) server program
