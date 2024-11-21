package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check authentication status",
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, err := os.Stat(dirname + "/.toclido-cli"); err == nil {
			file, err := os.Open(dirname + "/.toclido-cli")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()
			var p Payload
			json.NewDecoder(file).Decode(&p)
			user := p.User
			fmt.Println("\nLogged in as " + user + "\n")
			return
		} else {
			fmt.Println("Not logged in")
		}
	},
}

func init() {
	authCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
