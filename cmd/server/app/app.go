package app

import (
	"charting/pkg/server"
	"fmt"
	"github.com/spf13/cobra"
)

var bindAddress string
var bindPort int

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "chat-server",
		Run: func(cmd *cobra.Command, args []string) {
			s := server.NewServer(bindAddress, bindPort)
			if err := s.Run(); err != nil {
				fmt.Println(err)
			}
		},
	}
	cmd.Flags().StringVar(&bindAddress, "bind-address", "0.0.0.0", "")
	cmd.Flags().IntVar(&bindPort, "bind-port", 8080, "")
	return cmd
}
