package conn

import (
	"charting/pkg/client"
	"charting/pkg/cmd/chat/signal"
	"charting/pkg/profile"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var channelName string
var userName string

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "conn",
		Short: "connect to a channel",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := profile.View()
			if err != nil {
				fmt.Println(err)
			}
			ctx, err := profile.GetContext(p.CurrentContext)
			if err != nil {
				fmt.Println(err)
			}
			if channelName != "" {
				ctx.Channel = channelName
			}
			if userName != "" {
				ctx.User = userName
			}

			c, err := client.NewClient(&ctx)
			if err != nil {
				fmt.Println(err)
			}

			stop := signal.SetUpSignalHandler()
			if err := c.Connect(stop, os.Stdin, os.Stdout, ctx.Channel); err != nil {
				fmt.Println(err)
			}
			fmt.Println("ok")
		},
	}
	cmd.Flags().StringVar(&channelName, "channel", "", "")
	cmd.Flags().StringVar(&userName, "user", "", "")
	return cmd
}
