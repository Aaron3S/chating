package profile

import (
	"charting/pkg/cmd/chat/profile/create"
	"charting/pkg/cmd/chat/profile/show"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "manage your users and servers",
	}
	cmd.AddCommand(create.NewCommand())
	cmd.AddCommand(show.NewCommand())
	return cmd
}
