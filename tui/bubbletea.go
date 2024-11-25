package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/muesli/termenv"
)

func bubbleTeaMiddleware() wish.Middleware {
	newProgram := func(m tea.Model, opts ...tea.ProgramOption) *tea.Program {
		p := tea.NewProgram(m, opts...)
		return p
	}

	teaHandler := func(s ssh.Session) *tea.Program {
		_, _, active := s.Pty()
		if !active {
			wish.Fatalln(s, "no active terminal, skipping")
			return nil
		}

		hash := sha256.Sum256(s.PublicKey().Marshal())
		fingerprint := hex.EncodeToString(hash[:])
		payload := fmt.Sprintf(`{"fingerprint":"%s"}`, fingerprint)

		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://localhost:3000/api/ssh-keys/token", strings.NewReader(payload))

		if err != nil {
			wish.Println(s, "Some Error Occured")
			fmt.Println(err)
			s.Close()
			return nil
		}
		req.Header.Add("Content-Type", "application/json")

		wish.Println(s, "Verifying SSH Keys with the server")
		res, err := client.Do(req)

		if err != nil {
			wish.Println(s, "Some Error Occured")
			fmt.Println(err)
			s.Close()
			return nil
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		if strings.Contains(string(body), "Invalid public key") {
			wish.Println(s, "Invalid SSH Key")
			s.Close()
			return nil
		}

		m := NewModel()
		json.Unmarshal(body, &m)
		return newProgram(m, append(bubbletea.MakeOptions(s), tea.WithAltScreen())...)
	}
	return bubbletea.MiddlewareWithProgramHandler(teaHandler, termenv.ANSI256)
}
