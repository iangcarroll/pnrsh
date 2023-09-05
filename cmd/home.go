package main

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := Parse("home.html")
	t.Execute(w, struct {
		CommitHash string
	}{
		commitHash,
	})
}

func DeltaHomeHandler(w http.ResponseWriter, r *http.Request) {
	t := Parse("delta-home.html")
	t.Execute(w, struct {
		Error      bool
		CommitHash string
	}{
		r.URL.Query().Get("error") == "t",
		commitHash,
	})
}

func AeromexicoHomeHandler(w http.ResponseWriter, r *http.Request) {
	t := Parse("aeromexico-home.html")
	t.Execute(w, struct {
		Error      bool
		CommitHash string
	}{
		r.URL.Query().Get("error") == "t",
		commitHash,
	})
}

func AircanadaHomeHandler(w http.ResponseWriter, r *http.Request) {
	t := Parse("aircanada-home.html")
	t.Execute(w, struct {
		Error      bool
		CommitHash string
	}{
		r.URL.Query().Get("error") == "t",
		commitHash,
	})
}

func AirfranceklmHomeHandler(w http.ResponseWriter, r *http.Request) {
	t := Parse("airfranceklm-home.html")
	t.Execute(w, struct {
		Error      bool
		CommitHash string
	}{
		r.URL.Query().Get("error") == "t",
		commitHash,
	})
}

func UnitedHomeHandler(w http.ResponseWriter, r *http.Request) {
	t := Parse("united-home.html")
	t.Execute(w, struct {
		Error      bool
		CommitHash string
	}{
		r.URL.Query().Get("error") == "t",
		commitHash,
	})
}

func VirginHomeHandler(w http.ResponseWriter, r *http.Request) {
	t := Parse("virgin-home.html")
	t.Execute(w, struct {
		Error      bool
		CommitHash string
	}{
		r.URL.Query().Get("error") == "t",
		commitHash,
	})
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	t := Parse("help.html")
	t.Execute(w, struct {
		CommitHash string
	}{commitHash})
}
