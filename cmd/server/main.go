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
	playerID := uuid.New()
	for i, g := range s.Games {
		if g.BlackPlayer == nil {
			ch := make(chan backend.Move)
			player := backend.Player{
				Name:          msg.Name,
				ChangeChannel: ch,
				Id:            playerID,
				IsWhite:       false,
			}
			s.Games[i].BlackPlayer = &player
			resp := &pb.ConnectResponse{Token: playerID.String(), Colour: "black"}
			return resp, nil
		}
	}
	ch := make(chan backend.Move)
	player := backend.Player{
		Name:          msg.Name,
		ChangeChannel: ch,
		Id:            playerID,
		IsWhite:       true,
	}
	newGame := backend.Game{WhitePlayer: &player, WhiteTurn: true}

	s.Games = append(s.Games, newGame)
	resp := &pb.ConnectResponse{Token: playerID.String(), Colour: "white"}
	return resp, nil
}

func (s *server) ListenEvents(srv pb.SendMoveRequest_MoveServer, id uuid.UUID) error {
	var ch chan backend.Move
	var gameID int
	//There is a nil panic error here in case the uuid does not match with any player FIX THIS
	for i, game := range s.Games {
		if game.WhitePlayer.Id == id {
			ch = game.WhitePlayer.ChangeChannel
			gameID = i
		} else {
			if game.BlackPlayer.Id == id {
				ch = game.BlackPlayer.ChangeChannel
				gameID = i
			}
		}
	}
	for {
		select {
		case move := <-ch:
			resp := pb.MoveAnswer{
				XInitPos:  int32(move.InitialX),
				YInitPos:  int32(move.InitialY),
				XFinalPos: int32(move.FinalX),
				YFinalPos: int32(move.FinalY),
				WhiteTurn: s.Games[gameID].WhiteTurn,
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
		var gameID int
		for i, game := range s.Games {
			id, err := uuid.Parse(req.Id)
			if err != nil {
				return err
			}
			if game.WhitePlayer.Id == id {
				gameID = i
			} else {
				if game.BlackPlayer.Id == id {
					gameID = i
				}
			}
		}
		if started {
			move := backend.Move{
				InitialX: int(req.XInitPos),
				InitialY: int(req.YInitPos),
				FinalX:   int(req.XFinalPos),
				FinalY:   int(req.YFinalPos),
			}
			s.Games[gameID].WhiteTurn = !s.Games[gameID].WhiteTurn
			s.Games[gameID].WhitePlayer.ChangeChannel <- move
			s.Games[gameID].BlackPlayer.ChangeChannel <- move
		} else {
			uid, err := uuid.Parse(req.Id)
			if err != nil {
				return err
			}
			go s.ListenEvents(srv, uid)

			started = true
		}
	}
}
