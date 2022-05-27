package app

import (
	"charting/pkg/cmd/chat"
	"github.com/spf13/cobra"
)


func CreateRootCmd() *cobra.Command {
	return chat.NewCommand()
}
