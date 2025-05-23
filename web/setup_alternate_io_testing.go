package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (web *Web) sendEStopState(state bool) error {
	url := "http://localhost:8080/freezy/eStopState"

	// Build payload with channels 1-12 all set to given state
	payload := make([]map[string]interface{}, 13)
	for i := 0; i < 12; i++ {
		payload[i] = map[string]interface{}{
			"channel": i + 1,
			"state":   state,
		}
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	return nil
}
func (web *Web) altIOEStopHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		State bool `json:"state"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := web.sendEStopState(payload.State); err != nil {
		http.Error(w, "Failed to send E-Stop state", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
