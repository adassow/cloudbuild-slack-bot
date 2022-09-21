package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"google.golang.org/api/cloudbuild/v1"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load("environment.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/receive", slashCommandHandler)

	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":9090", nil)
}

func slashCommandHandler(w http.ResponseWriter, r *http.Request) {

	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(os.Getenv("SLACK_VERIFICATION_TOKEN")) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	response := fmt.Sprintf("Command not found %v", s.Command)
	switch s.Command {
	case "/build":
		response = fmt.Sprintf("Params not supported %v", s.Text)
		params := &slack.Msg{Text: s.Text}
		parsedParams := strings.Split(params.Text, " ")
		switch parsedParams[0] {
		case "trigger":
			cloudbuildService, err := cloudbuild.NewService(r.Context())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("StatusInternalServerError"))
				return
			}
			_, err = cloudbuildService.Projects.Triggers.Run(os.Getenv("PROJECT_ID"), os.Getenv("TRIGGER_ID"), &cloudbuild.RepoSource{BranchName: parsedParams[1]}).Do()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("StatusInternalServerError"))
				return
			}
			response = fmt.Sprintf("Build triggered for  %v", parsedParams[1])
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response))
		return
	}
}
