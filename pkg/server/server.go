package server

import (
	api "charting/pkg/api"
	"charting/pkg/channel"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
	Address      string
	Port         int
	channelStore *channel.Store
}

func (s *Server) ListChannels(context.Context, *api.ListChannelsRequest) (*api.ListChannelsResponse, error) {
	return nil, nil
}

func (s *Server) Send(ctx context.Context, req *api.SendMessageRequest) (*api.SendMessageResponse, error) {
	chName := req.Channel
	ch := s.channelStore.Get(chName)
	if ch == nil {
		return nil, fmt.Errorf("channel %s not found on server", chName)
	}
	ch.Write(req.Message)
	return &api.SendMessageResponse{Ok: true}, nil
}
func (s *Server) Receive(ctx context.Context, req *api.ReceiveMessageRequest) (*api.ReceiveMessageResponse, error) {
	chName := req.Channel
	ch := s.channelStore.Get(chName)
	if ch == nil {
		return nil, fmt.Errorf("channel %s not found on server", chName)
	}
	userName := req.UserName
	ms := ch.Read(userName)
	resp := &api.ReceiveMessageResponse{Ok: true, Messages: ms}
	return resp, nil
}

var errConnectionClosed = errors.New("connection already closed")
var errConnectionNotPrepare = errors.New("connection not opened")

func (s *Server) Connect(server api.MessageService_ConnectServer) error {
	msg, err := server.Recv()
	if err != nil {
		return err
	}
	chName := msg.Channel
	userName := msg.Message.UserName
	ch := s.channelStore.Get(chName)
	if ch == nil {
		return fmt.Errorf("channel %s not found on server", chName)
	}
	if err := server.SendMsg(&api.StreamMessage{}); err != nil {
		return err
	}
	stop := make(chan error)
	ctx, cancel := context.WithCancel(server.Context())
	go startReceiver(ctx, stop, ch, server)
	go startSender(ctx, stop, userName, ch, server)
	err = <-stop
	cancel()
	return err
}

func startSender(ctx context.Context, stop chan error, userName string, ch *channel.Channel, server api.MessageService_ConnectServer) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ms := ch.Read(userName)
			if len(ms) > 0 {
				for i := range ms {
					if err := server.Send(&api.StreamMessage{Message: ms[i], Channel: ch.Name}); err != nil {
						stop <- err
						return
					}
				}
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func startReceiver(ctx context.Context, stop chan error, ch *channel.Channel, server api.MessageService_ConnectServer) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := server.Recv()
			if err != nil {
				stop <- err
				return
			}
			ch.Write(msg.Message)
		}
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.Address, s.Port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	api.RegisterChannelServiceServer(srv, s)
	api.RegisterMessageServiceServer(srv, s)
	if err := srv.Serve(listen); err != nil {
		return err
	}
	return nil
}

func NewServer(address string, port int) *Server {
	cs := channel.NewStore()
	defaultCh := channel.NewChannel("default")
	cs.Put(defaultCh)
	return &Server{
		Address:      address,
		Port:         port,
		channelStore: cs,
	}
}
