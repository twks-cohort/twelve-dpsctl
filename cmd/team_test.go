package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFixedValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/teams/dps1" {
			t.Errorf("Expected to request '/v1/teams/dps1', got: %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"name":"dps1"}`))
	}))
	defer server.Close()

	team, _ := GetTeamHandler(server.URL, "dps1")
	if team.Name != "dps1" {
		t.Errorf("Expected 'dps1', got %s", team)
	}
}

func TestGetTeams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/teams" {
			t.Errorf("Expected to request '/v1/teams/dps1', got: %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"name":"dps2"}]`))
	}))
	defer server.Close()

	var teams []Team
	responseBytes := GetTeamsData(server.URL)
	_ = json.Unmarshal(responseBytes, &teams)
	if teams[0].Name != "dps2" {
		t.Errorf("expected 'dps2', got %s", teams[0].Name)
	}
}
