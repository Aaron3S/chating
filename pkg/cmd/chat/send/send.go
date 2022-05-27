package send

import (
	"charting/pkg/api"
	"charting/pkg/client"
	"charting/pkg/profile"
	"fmt"
	"github.com/spf13/cobra"
)

var channelName string
var messageContent string

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send",
		Short: "send a message to channel",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := profile.View()
			if err != nil {
				fmt.Println(err)
			}
			ctx, err := profile.GetContext(p.CurrentContext)
			if err != nil {
				fmt.Println(err)
			}
			c, err := client.NewClient(&ctx)
			if err != nil {
				fmt.Println(err)
			}
			if err := c.SendMessage(channelName, &api.Message{Content: []byte(messageContent)}); err != nil {
				fmt.Println(err)
			}
			fmt.Println("ok")
		},
	}
	cmd.Flags().StringVar(&channelName, "channel", "", "")
	cmd.Flags().StringVarP(&messageContent, "message", "m", "", "")
	return cmd
}
