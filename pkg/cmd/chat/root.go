package chat

import (
	"charting/pkg/cmd/chat/conn"
	"charting/pkg/cmd/chat/profile"
	"charting/pkg/cmd/chat/receive"
	"charting/pkg/cmd/chat/send"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	root := &cobra.Command{
		Use:  "chat",
		Long: "chat is a terminal chat app client, based on grpc and golang",
	}
	root.AddCommand(send.NewCommand())
	root.AddCommand(receive.NewCommand())
	root.AddCommand(profile.NewCommand())
	root.AddCommand(conn.NewCommand())
	return root
}
