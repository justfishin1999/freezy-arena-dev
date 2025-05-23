package field

// notifies websocket of changes in a/e stop outside of match
func (arena *Arena) NotifyStationTripStatus() {
	matchInProgress := arena.MatchState != PreMatch &&
		arena.MatchState != StartMatch &&
		arena.MatchState != PostMatch &&
		arena.MatchState != TimeoutActive &&
		arena.MatchState != PostTimeout

	for id, station := range arena.AllianceStations {
		arena.StationTripNotifier.NotifyWithMessage(map[string]any{
			"StationId":       id,
			"AStopTripped":    station.AStop,
			"EStopTripped":    station.EStop,
			"MatchInProgress": matchInProgress,
		})
	}
}
