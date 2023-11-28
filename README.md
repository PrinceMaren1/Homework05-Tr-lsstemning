# Homework05-Tr-lsstemning

To start the program:
## Starting a server an its replicas
- Run the first Server with a given port. "go run Server.go -port 1000"
- Run subsequent Servers with a given port and the ports of existing its replicas seperated by a comma. 
- "go run Server.go -port 1001 -replicas "1000"
- "go run Server.go -port 1002 -replicas "1000,1001"
## Starting Clients
- Now run the first Client, with a given port to one of the running servers, aswell as a client id. "go run Client.go -port 1000"
- Repeat as many times as you like :)

## Quick start
If you want to run a quick demo and you are on windows, you can run the demo.bat file.