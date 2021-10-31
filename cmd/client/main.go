package main

import (
	//"context"
	"context"
	"fmt"
	"log"

	//"log"
	"os"
	"time"

	//"github.com/google/uuid"
	"github.com/google/uuid"
	pb "github.com/maurotory/chess-golang/api/proto"
	"github.com/maurotory/chess-golang/pkg/frontend"
	"google.golang.org/grpc"
	//	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is a chess game")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	g, err := frontend.InitBoard()
	if err != nil {
		return fmt.Errorf("could not create Game: %v", err)
	}

	playerID := uuid.New()

	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connected to localhost:8082")

	client := pb.NewSendMoveRequestClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Connect(ctx, &pb.ConnectRequest{Id: playerID.String(), Name: "Mauri", Password: "Password"})
	if err != nil {
		log.Fatalf("could not receive the response: %v", err)
	}
	g.PlayerID = playerID
	if resp.Colour == "white" {
		g.PlayerWhite = true
	} else if resp.Colour == "black" {
		g.PlayerWhite = false
	} else {
		return fmt.Errorf("not a correct color response")
	}

	g.Render()

	stream, err := client.Move(context.Background())
	if err != nil {
		return err
	}
	id := g.PlayerID.String()
	req := pb.MoveRequest{
		XInitPos:  0,
		YInitPos:  0,
		XFinalPos: 0,
		YFinalPos: 0,
		Id:        id,
	}
	stream.SendMsg(&req)
	done, err := g.RunBoard(stream)
	if err != nil {
		return err
	}
	select {
	case <-done:
		return nil
	}

	if err := g.FinishBoard(); err != nil {
		return fmt.Errorf("could not finish the Game correctly: %v", err)
	}

	return nil
}
