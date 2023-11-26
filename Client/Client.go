package main

import (
	gRPC "Homework05/Proto"
	"context"
	"flag"
	"fmt"
	"log"
	//"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var id = flag.String("id", "Client", "This clients Id / name")    // Used when communicating with the server(s)
var serverPort = flag.Int64("server", 1500, "The port for the initial server to use")
var alternatePort int64
var server gRPC.ServerConnectionClient
var ServerConn *grpc.ClientConn


func main() {
	flag.Parse()

	// Setup logging
	f, err := os.OpenFile(fmt.Sprintf("log_%v", *id), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Printf("Starting client")

	ConnectToServer(*serverPort)

	for {
		var input string
		fmt.Scan(&input)
		if input == "?" {
			PrintAuctionState()
		} else {
			bid, err := strconv.Atoi(input)

			if (err != nil) {
				fmt.Println("Could not parse input. Please write an integer to submit a bid, or a question mark to query current highest bid")
			}

			Bid(int64(bid))
		}
	}
}

func ConnectToServer(port int64) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	fmt.Printf("Client %v: Attemps to dial on port 9000\n", port)

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(fmt.Sprintf(":%v", port), opts...)
	if err != nil {
		fmt.Printf("Failed to Dial : %v\n", err)
		return
	}

	server = gRPC.NewServerConnectionClient(conn)
	ServerConn = conn
	fmt.Println("The connection is: ", conn.GetState().String())

	alternateConnection, _ := server.GetAlternateServer(context.Background(), &gRPC.Empty{})
	alternatePort = alternateConnection.ServerId
}

func Bid(bidAmount int64){
	result, err := server.Bid(context.Background(), &gRPC.ClientBid{Amount: bidAmount, ClientId: *id})
	
	if err != nil {
		ConnectToServer(alternatePort)
	} 
	
	if result.Ack == 1 { 
		fmt.Printf("Bid accepted")
	} else if result.Ack == 0 {
		fmt.Printf("Bid to low or the action has ended")
	}
}

func PrintAuctionState(){
	state, err := server.GetAuctionState(context.Background(), &gRPC.Empty{})

	if err != nil {
		ConnectToServer(alternatePort)
	} 

	if (state.IsCompleted) {
		fmt.Printf("The auction is completed. %v won with a bid of %v\n", state.BidderId, state.HighestBid)
	} else {
		fmt.Printf("The current highest bidder is %v, with a bid of %v\n", state.BidderId, state.HighestBid)
	}
}