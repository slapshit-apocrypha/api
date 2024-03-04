package web

import (
	"encoding/json"
	"net/http"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
)

type djApplication struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	PreferredDay string `json:"preferred_day"`
	Experience   bool   `json:"experience"`
	Genre        string `json:"genre"`
	Pronouns     string `json:"pronouns"`
}

func (s *Server) postSendWebhook(w http.ResponseWriter, r *http.Request) {
	var b *djApplication
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		respond(w, r, http.StatusBadRequest, "unable to decode JSON body")
		return
	}

	e := createDJEmbed(b)
	data := webhook.ExecuteData{
		Username: "DJ Applications",
		Embeds:   []discord.Embed{e},
	}

	if err := s.webhookClient.Execute(data); err != nil {
		respond(w, r, http.StatusInternalServerError, "error executing webhook")
		return
	}

	respond(w, r, http.StatusOK, "webhook message created")
}
