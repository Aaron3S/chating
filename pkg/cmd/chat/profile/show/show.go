package show

import (
	"charting/pkg/profile"
	"fmt"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show some resource",
	}
	cmd.AddCommand(newShowServersCommand())
	cmd.AddCommand(newShowUsersCommand())
	cmd.AddCommand(newShowContextCommand())
	cmd.AddCommand(newShowCurrentContextCommand())
	return cmd
}

func newShowServersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "show all servers",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := profile.View()
			if err != nil {
				fmt.Println(err)
			}
			for i := range p.Servers {
				n := fmt.Sprintf("%s", p.Servers[i].Url)
				fmt.Println(n)
			}
		},
	}
	return cmd
}
func newShowUsersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "show all users",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := profile.View()
			if err != nil {
				fmt.Println(err)
			}
			for i := range p.Users {
				n := fmt.Sprintf("%s  %s", p.Users[i].Name, p.Users[i].Email)
				fmt.Println(n)
			}
		},
	}
	return cmd
}

func newShowContextCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "context",
		Short: "show all contexts",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := profile.View()
			if err != nil {
				fmt.Println(err)
			}
			for i := range p.Contexts {
				n := fmt.Sprintf("%s  %s  %s", p.Contexts[i].Name, p.Contexts[i].User, p.Contexts[i].Server)
				if p.Contexts[i].Name == p.CurrentContext {
					n = fmt.Sprintf("* %s", n)
				}
				fmt.Println(n)
			}
		},
	}
	return cmd
}

func newShowCurrentContextCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "current-context",
		Short: "show current-context",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := profile.View()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(p.CurrentContext)
		},
	}
	return cmd
}
