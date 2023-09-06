package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
)

const (
	aeromexicoEnabled   = true
	aircanadaEnabled    = true
	airfranceklmEnabled = true
	deltaEnabled        = true
	unitedEnabled       = true
	virginEnabled       = true
)

var (
	commitHash = os.Getenv("HEROKU_SLUG_COMMIT")
)

func listenAddress() string {
	if port := os.Getenv("PORT"); port != "" {
		return "0.0.0.0:" + port
	}

	return "0.0.0.0:8080"
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/help", HelpHandler).Methods("GET")

	if aeromexicoEnabled {
		r.HandleFunc("/aeromexico", AeromexicoHomeHandler).Methods("GET")
		r.HandleFunc("/aeromexico", AeromexicoRetrieveHandler).Methods("POST")
	}

	if aircanadaEnabled {
		r.HandleFunc("/aircanada", AircanadaHomeHandler).Methods("GET")
		r.HandleFunc("/aircanada", AircanadaRetrieveHandler).Methods("POST")
	}

	if airfranceklmEnabled {
		r.HandleFunc("/airfranceklm", AirfranceklmHomeHandler).Methods("GET")
		r.HandleFunc("/airfranceklm", AirfranceklmRetrieveHandler).Methods("POST")
	}

	if deltaEnabled {
		r.HandleFunc("/delta", DeltaHomeHandler).Methods("GET")
		r.HandleFunc("/delta", DeltaRetrieveHandler).Methods("POST")
	}

	if unitedEnabled {
		r.HandleFunc("/united", UnitedHomeHandler).Methods("GET")
		r.HandleFunc("/united", UnitedRetrieveHandler).Methods("POST")
	}

	if virginEnabled {
		r.HandleFunc("/virgin", VirginHomeHandler).Methods("GET")
		r.HandleFunc("/virgin", VirginRetrieveHandler).Methods("POST")
	}

	srv := &http.Server{
		Handler: r,
		Addr:    listenAddress(),

		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	r.Use(secure.New(secure.Options{
		BrowserXssFilter:     true,
		ContentTypeNosniff:   true,
		FrameDeny:            true,
		STSSeconds:           31536000,
		STSIncludeSubdomains: true,
		STSPreload:           true,
	}).Handler)

	log.Println("Visit", "http://"+listenAddress(), "to use the app!")
	log.Fatal(srv.ListenAndServe())
}
