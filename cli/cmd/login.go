package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Toclido",
	Run:   login,
}

func init() {
	authCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Payload struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")

}

func login(cmd *cobra.Command, args []string) {
	// check if user is already logged in
	fmt.Println("Open the following URL in your browser:")
	fmt.Println("")
	fmt.Println("\thttp://localhost:3001/auth/cli/login")
	fmt.Println("")

	server := &http.Server{Addr: ":1606"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == "POST" {
			w.Header().Set("Content-Type", "application/json")
			var p Payload
			json.NewDecoder(r.Body).Decode(&p)
			// save p.Token to local
			dirname, err := os.UserHomeDir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			file, err := os.Create(dirname + "/.toclido-cli")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()
			jsonstr, err := json.Marshal(p)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			_, err = file.WriteString(string(jsonstr))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// close the server
			defer server.Close()
			// display auth complete
			fmt.Println("Authentication Successful.\nLogged in as " + p.User)
		}
	})

	server.ListenAndServe()
}
