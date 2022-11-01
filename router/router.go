package router

import "net/http"

type ledger interface {
	Search(w http.ResponseWriter, r *http.Request)
	Upload(w http.ResponseWriter, r *http.Request)
}

func New(l ledger) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static/")))
	mux.HandleFunc("/q", l.Search)
	mux.HandleFunc("/u", l.Upload)
	return mux
}
