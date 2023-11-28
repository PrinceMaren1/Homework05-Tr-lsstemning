# Homework05-Tr-lsstemning

To start the program:
## Starting a server an its replicas
- Run the first Server with a given port. "go run Server.go -port 1000"
- Run subsequent Servers with a given port and the ports of existing its replicas seperated by a comma. 
- "go run Server.go -port 1001 -replicas "1000"
- "go run Server.go -port 1002 -replicas "1000,1001"
- Note that the auction runs for 60 seconds from the moment that the first server is launched. 
## Starting Clients
- Now run the first Client, giving a port to any one of the running servers, aswell as a client id. "go run Client.go -port 1000 -id BillyTheBid"
- Repeat as many times as you like :)
- If you want to bid from a client, just type the number, that you want to bid.
- If you want to know the current state of the auction, just type "?".

## Quick start
- If you want to run a quick demo and you are on windows, you can run the demo.bat file.
- You still have to bid and such from the clients.