package client

import (
	"bufio"
	api "charting/pkg/api"
	"charting/pkg/profile"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
)

type Client struct {
	context              *profile.Context
	channelServiceClient api.ChannelServiceClient
	messageServiceClient api.MessageServiceClient
}

func NewClient(context *profile.Context) (*Client, error) {
	conn, err := grpc.Dial(context.Server, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		context:              context,
		channelServiceClient: api.NewChannelServiceClient(conn),
		messageServiceClient: api.NewMessageServiceClient(conn),
	}, nil
}

func (c *Client) SendMessage(channel string, msg *api.Message) error {
	msg.UserName = c.context.User
	req := &api.SendMessageRequest{
		Message: msg,
		Channel: channel,
	}
	if channel == "" {
		req.Channel = c.context.Channel
	}

	_, err := c.messageServiceClient.Send(context.TODO(), req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ReceiveMessage(channel string) ([]*api.Message, error) {
	req := &api.ReceiveMessageRequest{
		Channel:  channel,
		UserName: c.context.User,
	}
	if channel == "" {
		req.Channel = c.context.Channel
	}

	resp, err := c.messageServiceClient.Receive(context.TODO(), req)
	if err != nil {
		return nil, err
	}
	return resp.Messages, nil
}

func (c *Client) Connect(stop <-chan struct{}, reader io.ReadCloser, writer io.WriteCloser, channel string) error {
	userName := c.context.User
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := c.messageServiceClient.Connect(ctx)
	if err != nil {
		return err
	}
	// 建立连接发送空消息
	err = conn.Send(&api.StreamMessage{
		Message: &api.Message{
			UserName: userName,
		},
		Channel: channel,
	})
	if err != nil {
		return err
	}
	_, err = conn.Recv()
	if err != nil {
		return err
	}
	// 开启一个消息接收循环
	errChan := make(chan error)
	go startReceiver(ctx, errChan, writer, conn)
	for {
		select {
		case <-stop:
			cancel()
			return nil
		case err = <-errChan:
			return err
		default:
			reader := bufio.NewReader(reader)
			buffer, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			if err := conn.Send(&api.StreamMessage{
				Message: &api.Message{
					Content:  []byte(buffer),
					UserName: userName,
				},
				Channel: channel,
			}); err != nil {
				return err
			}
		}
	}
}

func startReceiver(ctx context.Context, stop chan error, writer io.WriteCloser, conn api.MessageService_ConnectClient) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := conn.Recv()
			if err != nil {
				stop <- err
				return
			}
			if _, err = fmt.Fprintf(writer, "[%s]: %s", msg.Message.UserName, string(msg.Message.Content)); err != nil {
				stop <- err
				return
			}
		}
	}
}
