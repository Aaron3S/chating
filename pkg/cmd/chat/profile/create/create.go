package create

import (
	"charting/pkg/profile"
	"fmt"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create some resource",
	}
	cmd.AddCommand(newCreateUserCommand())
	cmd.AddCommand(newCreateServerCommand())
	cmd.AddCommand(newCreateContextCommand())
	return cmd
}

var userName string
var userEmail string

func newCreateUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "create a user",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCreateUser(); err != nil {
				fmt.Println(err)
			}
		},
	}
	cmd.Flags().StringVar(&userName, "name", "", "")
	cmd.Flags().StringVar(&userEmail, "email", "", "")
	return cmd
}

var serverUrl string

func newCreateServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "add an server to profile",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCreateServer(); err != nil {
				fmt.Println(err)
			}
		},
	}
	cmd.Flags().StringVar(&serverUrl, "server", "", "")
	return cmd
}

var contextServer string
var contextUser string
var contextChannel string

func newCreateContextCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "context",
		Short: "create a chat context",
		Run: func(cmd *cobra.Command, args []string) {
			if contextUser == "" || contextServer == "" {
				fmt.Println("user and server can not be none")
				return
			}
			contextName := fmt.Sprintf("%s@%s", contextUser, contextServer)
			if contextChannel == "" {
				contextChannel = "default"
			}
			// TODO: save context to profile
			ctx := profile.Context{
				Name:    contextName,
				User:    contextUser,
				Server:  contextServer,
				Channel: contextChannel,
			}
			if err := profile.AddContext(&ctx); err != nil {
				fmt.Println(err)
			}
			fmt.Println(fmt.Sprintf("Context %s created", contextName))
		},
	}
	cmd.Flags().StringVar(&contextUser, "user", "", "")
	cmd.Flags().StringVar(&contextServer, "server", "", "")
	cmd.Flags().StringVar(&contextChannel, "channel", "", "")
	return cmd
}

func runCreateServer() error {
	serverName := ""
	if err := profile.AddServer(&profile.Server{Url: serverUrl}); err != nil {
		fmt.Println(err)
	}
	//TODO: add server to profile
	fmt.Println(fmt.Sprintf("Server %s created", serverName))
	return nil
}

func runCreateUser() error {
	if userName == "" {
		fmt.Println("Please input an name, eg: Bob")
		for {
			n, err := fmt.Scanln(&userName)
			if err != nil {
				return err
			}
			if n == 0 {
				fmt.Println("Please input an invalid name, eg: Bob")
				continue
			}
			//TODO: validate user name
			break
		}
	}
	if userEmail == "" {
		fmt.Println("Please input an email, eg: apple@huawei.com")
		for {
			n, err := fmt.Scanln(&userEmail)
			if err != nil {
				return err
			}
			if n == 0 {
				fmt.Println("Please input an invalid name, eg: apple@huawei.com")
				continue
			}
			//TODO: validate user name
			break
		}
	}
	//TODO: add user to profile
	if err := profile.AddUser(&profile.User{Name: userName, Email: userEmail}); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("User %s created", userName))
	return nil
}
