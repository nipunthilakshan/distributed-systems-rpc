# distributed-systems-rpc
A simple distributed application for vegetable sales system using Go language and Remote Procedure Calls (RPC).

Server will maintain a file which keeps records of different available vegetables including price per kg and available amount of kg for each vegetable. The server has following functions.
1. Query the file and output names of all available vegetables.
2. Output the price per kg of a given vegetable.
3. Output the available amount of kg for a given vegetable.
4. Add new vegetable to the file with price per kg and among of kg.
5. Update the price or available amount of a given vegetable.
   
Accordingly, clients can use server functions to do the following tasks.
1. Receive a list of all available vegetables and display.
2. Get the price per kg of a given vegetable and display.
3. Get the available amount of kg of a given vegetable and display.
4. Send a new vegetable name to the server to be added to the server file.
5. Send new price or available amount for a given vegetable to be updated in the server file.

### Steps to follow
1. go run server.go
2. go run client.go

### References
- https://golang.org/doc/tutorial/getting-started
- https://www.tutorialspoint.com/go/index.htm
- https://pkg.go.dev/net/rpc
- https://medium.com/rungo/building-rpc-remote-procedure-call-network-in-go-5bfebe90f7e9