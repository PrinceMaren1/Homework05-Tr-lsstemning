start go run Server/Server.go -port 1000
start go run Server/Server.go -port 1001 -replicas "1000"
start go run Server/Server.go -port 1002 -replicas "1000,1001"
REM start go run Server/Server.go -port 1003 -replicas "1000,1001,1002"
REM start go run Server/Server.go -port 1004 -replicas "1000,1001,1002,1003"

start go run Client/Client.go -server 1000 -id "BidGod"
start go run Client/Client.go -server 1001 -id "SortePowerRanger"
start go run Client/Client.go -server 1002 -id "WantToBuySpaces"
REM start go run Client/Client.go -server 1001 -id "TheAuctionator"
