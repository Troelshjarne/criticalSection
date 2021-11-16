# Critical Section/Mutual Exclusion
## Architecture
This implementation of the Critical Section problem, is solved using a client(s)-server architecture, making the problem rather trivial.  
A server retains access to the critical section, and doles it out to clients upon request, and clients can connect to the server or request access to the critical section at any time.

## Access requests
A client can, at any given time, decide that it wants to access the critical section and send a request to the server for permission. Upon sending a request, a client halts itself and awaits permission from the server.
The server continuously listens for access requests from clients and, upon receiving any, places them into a queue. Another thread takes the client at the front of the queue and, given that no one is currently in the critical section, grants permission to that client to enter the critical section. A client then has access to the critical section, until it elects to give up access.  
The server doles out access at a FIFO basis.
