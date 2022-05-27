package receive

import (
	"charting/pkg/client"
	"charting/pkg/profile"
	"fmt"
	"github.com/spf13/cobra"
)

var channelName string

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "receive",
		Short: "receive a message from channel",
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
			ms, err := c.ReceiveMessage(channelName)
			if err != nil {
				fmt.Println(err)
			}
			if len(ms) == 0 {
				fmt.Println("no message to show")
			} else {
				for i := 0; i < len(ms); i++ {
					fmt.Println(fmt.Sprintf("[%s] -> [%d] %s", ms[i].UserName, ms[i].Id, string(ms[i].Content)))
				}
			}
		},
	}
	cmd.Flags().StringVar(&channelName, "channel", "", "")
	return cmd
}
