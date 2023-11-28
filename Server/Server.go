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
var replicas = flag.String("replicas", "", "ports of server replicas") // example use -replicas "1000,1001,1002"
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

	replicaPorts := strings.Split(*replicas, ",")

	go launchServer()

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
		for state.TimeStamp > 0 {
			time.Sleep(time.Duration(1) * time.Second)
			state.TimeStamp = state.TimeStamp - 1
		}
		if !state.IsCompleted {
			log.Printf("Auction completed. Updating replicas.")
			state.IsCompleted = true
			UpdateReplicas()
		}
	}()

	for {

	}
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

func connect(dialPort int64) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	fmt.Printf("Server %v: Attemps to dial to replica node on port %v\n", *port, dialPort)

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(fmt.Sprintf(":%v", dialPort), opts...)

	if err != nil {
		fmt.Printf("Failed to Dial : %v\n", err)
		return
	}

	server := gRPC.NewServerConnectionClient(conn)
	serverReplicas[dialPort] = server
	log.Printf("Connecting to replica server at: %v\n", dialPort)
	fmt.Println("The connection is: ", conn.GetState().String())
}

func UpdateReplicas() {
	for key := range serverReplicas {
		freshestState, err := serverReplicas[key].UpdateAuctionState(context.Background(), state)

		if err != nil {
			delete(serverReplicas, key)
		} else {
			state = freshestState
		}
	}
}

func (s *Server) GetAlternateServer(ctx context.Context, msg *gRPC.Empty) (*gRPC.Connection, error) {
	var port int64 = 0
	log.Println("Fetching alternate port for client")

	for len(serverReplicas) == 0 {

	}

	for key := range serverReplicas {
		port = key
		break
	}

	log.Printf("Sending value to client: %v\n", port)
	return &gRPC.Connection{ServerId: port}, nil
}

func (s *Server) ConnectToReplica(ctx context.Context, msg *gRPC.Connection) (*gRPC.Empty, error) {
	connect(msg.ServerId)
	return &gRPC.Empty{}, nil
}

func (s *Server) GetAuctionState(ctx context.Context, msg *gRPC.ClientInfo) (*gRPC.AuctionState, error) {
	log.Printf("Sending auction state to client %v", msg.ClientId)
	return state, nil
}

func (s *Server) UpdateAuctionState(ctx context.Context, msg *gRPC.AuctionState) (*gRPC.AuctionState, error) {
	// Checking which state has the highest bid an update accordingly.
	if msg.HighestBid > state.HighestBid && !state.IsCompleted {
		log.Printf("Updating highest bid to align with replica")
		state.HighestBid = msg.HighestBid
		state.BidderId = msg.BidderId
	}
	if !state.IsCompleted {
		if msg.IsCompleted {
			log.Printf("Closing auction to align with replica")
		}
		state.IsCompleted = msg.IsCompleted
	}
	// Set Auction time to the lowest of the two states
	state.TimeStamp = min(msg.TimeStamp, state.TimeStamp)
	// Update completed state
	// Returning updated states
	return state, nil
}

func (s *Server) Bid(ctx context.Context, msg *gRPC.ClientBid) (*gRPC.Acknowledgement, error) {
	log.Printf("Recived bid of %v from client %v\n", msg.Amount, msg.ClientId)
	var bidSucceded int64 = 0

	if (msg.Amount > state.HighestBid) && !state.IsCompleted {
		log.Printf("Bid seems good unless replicas has a higher bid. Updating state in replicas.")
		state.HighestBid = msg.Amount
		state.BidderId = msg.ClientId

		// Update state of other replicas. Waits for all of them to respond before responding to client
		UpdateReplicas()
		// The state update goes both ways, so we confirm that the bidder is still highest after confirming with other replicas
		if state.BidderId == msg.ClientId && state.HighestBid == msg.Amount {
			bidSucceded = 1
		}
	}

	if bidSucceded == 1 {
		log.Printf("Sending BID SUCCESS response to client.")
	} else {
		log.Printf("Sending BID DENIED response to client.")
	}

	return &gRPC.Acknowledgement{Ack: bidSucceded}, nil
}
