package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/google/uuid"
	pb "github.com/maurotory/chess-golang/api/proto"
	"github.com/maurotory/chess-golang/pkg/backend"
	"google.golang.org/grpc"
)

var port = flag.String("p", ":8082", "listen port")

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterSendMoveRequestServer(s, &server{})

	fmt.Println("Listenning on port 8082")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	pb.UnimplementedSendMoveRequestServer
	Games []backend.Game
}

func (s *server) Connect(ctx context.Context, msg *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	fmt.Println(len(s.Games))
	if len(s.Games) == 0 {
		id, err := uuid.Parse(msg.Id)
		if err != nil {
			return nil, fmt.Errorf("could not parse the uuid")
		}
		ch := make(chan backend.Move)
		player := backend.Player{
			Name:          "Player1",
			ChangeChannel: ch,
			Id:            id,
			IsWhite:       true,
		}
		newGame := backend.Game{WhitePlayer: &player, WhiteTurn: true}

		s.Games = append(s.Games, newGame)
		resp := &pb.ConnectResponse{Token: "token1", Colour: "white"}
		return resp, nil
	} else if len(s.Games) == 1 {
		id, err := uuid.Parse(msg.Id)
		if err != nil {
			return nil, fmt.Errorf("could not parse the uuid")
		}
		ch := make(chan backend.Move)
		player := backend.Player{
			Name:          "Player2",
			ChangeChannel: ch,
			Id:            id,
			IsWhite:       false,
		}
		s.Games[0].BlackPlayer = &player
		resp := &pb.ConnectResponse{Token: "token2", Colour: "black"}
		return resp, nil
	}
	return nil, fmt.Errorf("there are already 2 players playing")
}

func (s *server) ListenEvents(srv pb.SendMoveRequest_MoveServer, id uuid.UUID) error {
	var ch chan backend.Move
	if s.Games[0].WhitePlayer.Id == id {
		ch = s.Games[0].WhitePlayer.ChangeChannel
	} else if s.Games[0].BlackPlayer.Id == id {
		ch = s.Games[0].BlackPlayer.ChangeChannel
	} else {
		return fmt.Errorf("id not found")
	}
	for {
		select {
		case move := <-ch:
			fmt.Println("received from the channel")
			resp := pb.MoveAnswer{
				XInitPos:  int32(move.InitialX),
				YInitPos:  int32(move.InitialY),
				XFinalPos: int32(move.FinalX),
				YFinalPos: int32(move.FinalY),
				WhiteTurn: s.Games[0].WhiteTurn,
			}
			err := srv.Send(&resp)
			if err != nil {
				fmt.Println("error sending the data")
				return err
			}
		}

	}
	return nil
}

func (s *server) Move(srv pb.SendMoveRequest_MoveServer) error {
	// ctx := srv.Context()
	fmt.Println("stream openned in the server")
	var id string
	var started bool
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			return err
		}
		if started {
			move := backend.Move{
				InitialX: int(req.XInitPos),
				InitialY: int(req.YInitPos),
				FinalX:   int(req.XFinalPos),
				FinalY:   int(req.YFinalPos),
			}
			s.Games[0].WhiteTurn = !s.Games[0].WhiteTurn
			s.Games[0].WhitePlayer.ChangeChannel <- move
			s.Games[0].BlackPlayer.ChangeChannel <- move
		} else {
			uid, err := uuid.Parse(req.Id)
			if err != nil {
				return err
			}
			id = uid.String()
			fmt.Printf("Id initialized: %s\n", id)
			go s.ListenEvents(srv, uid)

			started = true
		}
	}
}
