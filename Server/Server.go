package main

import (
	gRPC "Homework05/Proto"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	gRPC.UnimplementedServerConnectionServer
}

var port = flag.Int64("port", 1000, "port to use for this")
var replicas = flag.String("Replicas", "", "ports of server replicas") // example use -Replicas "1000,1001,1002"
var serverReplicas map[int64]gRPC.ServerConnectionClient = make(map[int64]gRPC.ServerConnectionClient)
var state = &gRPC.AuctionState{}


func main() {
	flag.Parse()

	// Setup logging
	f, err := os.OpenFile(fmt.Sprintf("log_%v", *port), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Starting server")
	state.TimeStamp = 30 // Time counts down. When it hits 0, the auction ends

	replicaPorts := strings.Split(*replicas,",")

	for i := 0; i < len(replicaPorts); i++ {
		if replicaPorts[i] == "" {
			continue
		}
		p, err := strconv.Atoi(replicaPorts[i])

		if err != nil {
			log.Println("PANIC")
		}

		connect(int64(p))
		_, _ = serverReplicas[int64(p)].ConnectToReplica(context.Background(), &gRPC.Connection{ServerId: *port})
	}

	go func() {
		for (state.TimeStamp > 0) {
			time.Sleep(time.Duration(1) * time.Second)
			state.TimeStamp = state.TimeStamp - 1
		}
		if (!state.IsCompleted) {
			state.IsCompleted = true
			UpdateReplicas()
		}
	}()

	launchServer()
}

func launchServer() {
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Printf("Failed to listen on port %v: %v\n", *port, err)
		return
	}
	grpcServer := grpc.NewServer()

	server := &Server{}

	gRPC.RegisterServerConnectionServer(grpcServer, server)

	log.Println("Started listening for incoming messages")
	if err := grpcServer.Serve(list); err != nil {
		log.Printf("Failed to serve gRPC server over port %v %v\n", port, err)
	}
}

func connect(dialPort int64){
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	fmt.Printf("Server %v: Attemps to dial to replica node on port %v", *port, dialPort)

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(fmt.Sprintf(":%v", dialPort), opts...)

	if err != nil {
		fmt.Printf("Failed to Dial : %v\n", err)
		return
	}

	server := gRPC.NewServerConnectionClient(conn)
	serverReplicas[dialPort] = server
	fmt.Println("The connection is: ", conn.GetState().String())
}

func UpdateReplicas() {
	for i := 0; i < len(serverReplicas); i++ {
		freshestState, err := serverReplicas[int64(i)].UpdateAuctionState(context.Background(), state)
		
		if (err != nil) {
			delete(serverReplicas,int64(i))
		}

		state = freshestState
	}
}

func (s *Server) GetAlternateServer(ctx context.Context, msg *gRPC.Empty) (*gRPC.Connection, error) {
	var port int64 = 0
	for key := range serverReplicas {
		port = key
		break
	}
	return &gRPC.Connection{ServerId: port}, nil
}

func (s *Server) ConnectToReplica(ctx context.Context, msg *gRPC.Connection) (*gRPC.Empty, error) {
	connect(msg.ServerId)
	return &gRPC.Empty{}, nil
}

func (s *Server) GetAuctionState(ctx context.Context, msg *gRPC.Empty) (*gRPC.AuctionState, error) {
	return state, nil
}

func (s *Server) UpdateAuctionState(ctx context.Context, msg *gRPC.AuctionState) (*gRPC.AuctionState, error) {
	// Checking which state has the highest bid an update accordingly.
	if msg.HighestBid > state.HighestBid && !state.IsCompleted{
		state.HighestBid = msg.HighestBid
		state.BidderId = msg.BidderId
	}
	if !state.IsCompleted {
		state.IsCompleted = msg.IsCompleted
	}
	// Set Action time to the lowest of the	two states
	state.TimeStamp = min(msg.TimeStamp, state.TimeStamp)
	// Update completed state
	// Returning updated states
	return state, nil
}

func (s *Server) Bid(ctx context.Context, msg *gRPC.ClientBid) (*gRPC.Acknowledgement, error) {
	var bidSucceded int64 = 0

	if ((msg.Amount > state.HighestBid) && !state.IsCompleted) {
		state.HighestBid = msg.Amount
		state.BidderId = msg.ClientId

		// Update state of other replicas. Waits for all of them to respond before responding to client
		UpdateReplicas()
		// The state update goes both ways, so we confirm that the bidder is still highest after confirming with other replicas
		if (state.BidderId == msg.ClientId && state.HighestBid == msg.Amount) {
			bidSucceded = 1
		}
	}

	return &gRPC.Acknowledgement{Ack: bidSucceded}, nil
}