package web

import (
	"fmt"
	"log"
	"net/http"
)

// ServeAnnouncerInterface spins up a second HTTP server on the given port
// that only serves the announcer display and its static assets.
func (w *Web) ServeAnnouncerInterface(port int) {
	mux := http.NewServeMux()

	// 1) Static assets (css/js/images) from /static/…
	mux.Handle("/static/",
		http.StripPrefix("/static/", addNoCacheHeader(http.FileServer(http.Dir("static/")))),
	)

	// 2) Root path → announcer HTML
	mux.HandleFunc("/", w.announcerDisplayHandler)

	// 3) Websocket endpoint for live updates
	mux.HandleFunc("GET /displays/announcer/match_load", w.announcerDisplayMatchLoadHandler)
	mux.HandleFunc("GET /displays/announcer/score_posted", w.announcerDisplayScorePostedHandler)
	mux.HandleFunc("GET /displays/announcer/websocket", w.announcerDisplayWebsocketHandler)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Announcer interface listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
