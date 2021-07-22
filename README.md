# onion-farm
onion network on one server with container microservices as the nodes

<img src='onion-server.png'/>

Here's a really good picture describing what I have in mind

So we have one server that routes messages, but when they exit the server we don't know who they came from, because in the routing the message goes through a bunch of onion nodes, which we can implement as microservice containers.

I'm wondering if they don't even have to be containers that work as APIs, they could also just be in-program nodes in an in-memory graph. I'd like to do the containers though because communicating through those would be the same way that one would communicate through a real-world network where all of the nodes are real servers scattered across the world.
