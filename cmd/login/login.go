package login

import (
	"fmt"
	"os"

	"meroshare-cli/internal/login"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to MeroShare",
	Long: `Login to MeroShare using your credentials.
You will be prompted to enter your username and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if username == "" || password == "" {
			fmt.Println("Username and password are required.")
			os.Exit(1)
		}

		if err := login.Login(username, password); err != nil {
			fmt.Printf("Login failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Login successful!")
	},
}

func init() {
	loginCmd.Flags().StringP("username", "u", "", "MeroShare username")
	loginCmd.Flags().StringP("password", "p", "", "MeroShare password")
}
