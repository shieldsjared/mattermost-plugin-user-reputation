package main

import (
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	switch path := r.URL.Path; path {
	case "/api/v1/reputation":
		p.handleGetUserReputation(w, r)
		return
	default:
		http.NotFound(w, r)
	}
}

func (p *Plugin) handleGetUserReputation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	userId := r.URL.Query().Get("user_id")

	if userId == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	reputations, err := p.GetUserReputations(userId)

	if err != nil {
		p.API.LogError("Error fetching user reputations", err)
		http.Error(w, "Error fetching user reputations", http.StatusInternalServerError)
		return
	}

	json, jsonErr := json.Marshal(reputations)
	if jsonErr != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(json)
	if err != nil {
		p.API.LogError("failed to write http response", err.Error())
	}
}

