package main

import (
	"net/http"

	aeromexico "github.com/pnrsh/pnrsh/pkg/aeromexico/pnr"
	delta "github.com/pnrsh/pnrsh/pkg/delta/pnr"
	united "github.com/pnrsh/pnrsh/pkg/united/pnr"
)

func DeltaRetrieveHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Add("Location", "/delta?error=t")
		w.WriteHeader(302)
		return
	}

	firstName := r.Form.Get("first_name")
	lastName := r.Form.Get("last_name")
	confirmationCode := r.Form.Get("confirmation_code")

	if len(confirmationCode) != 6 || len(firstName) == 0 || len(lastName) == 0 {
		w.Header().Add("Location", "/delta?error=t")
		w.WriteHeader(302)
		return
	}

	retrievedPNR, err := delta.Retrieve(delta.DeltaEndpoint, firstName, lastName, confirmationCode)
	if err != nil {
		w.Header().Add("Location", "/delta?error=t")
		w.WriteHeader(302)
		return
	}

	t := Parse("delta-show.html")

	t.Execute(w, struct {
		PNR              delta.PNR
		ConfirmationCode string
		CommitHash       string
	}{
		retrievedPNR,
		confirmationCode,
		commitHash,
	})
}

func AeromexicoRetrieveHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Add("Location", "/aeromexico?error=t")
		w.WriteHeader(302)
		return
	}

	lastName := r.Form.Get("last_name")
	confirmationCode := r.Form.Get("confirmation_code")

	if len(confirmationCode) != 6 || len(lastName) == 0 {
		w.Header().Add("Location", "/aeromexico?error=t")
		w.WriteHeader(302)
		return
	}

	retrievedPNR, err := aeromexico.Retrieve(lastName, confirmationCode)
	if err != nil {
		w.Header().Add("Location", "/aeromexico?error=t")
		w.WriteHeader(302)
		return
	}

	t := Parse("aeromexico-show.html")

	t.Execute(w, struct {
		PNR              aeromexico.PNR
		ConfirmationCode string
		CommitHash       string
	}{
		retrievedPNR,
		confirmationCode,
		commitHash,
	})
}

func UnitedRetrieveHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Add("Location", "/united?error=t")
		w.WriteHeader(302)
		return
	}

	lastName := r.Form.Get("last_name")
	confirmationCode := r.Form.Get("confirmation_code")

	if len(confirmationCode) != 6 || len(lastName) == 0 {
		w.Header().Add("Location", "/united?error=t")
		w.WriteHeader(302)
		return
	}

	retrievedPNR, err := united.Retrieve(lastName, confirmationCode)
	if err != nil {
		w.Header().Add("Location", "/united?error=t")
		w.WriteHeader(302)
		return
	}

	t := Parse("united-show.html")

	t.Execute(w, struct {
		PNR              united.PNR
		ConfirmationCode string
		CommitHash       string
	}{
		retrievedPNR,
		confirmationCode,
		commitHash,
	})
}

func VirginRetrieveHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Add("Location", "/virgin?error=t")
		w.WriteHeader(302)
		return
	}

	firstName := r.Form.Get("first_name")
	lastName := r.Form.Get("last_name")
	confirmationCode := r.Form.Get("confirmation_code")

	if len(confirmationCode) != 6 || len(firstName) == 0 || len(lastName) == 0 {
		w.Header().Add("Location", "/virgin?error=t")
		w.WriteHeader(302)
		return
	}

	retrievedPNR, err := delta.Retrieve(delta.VirginAtlanticEndpoint, firstName, lastName, confirmationCode)
	if err != nil {
		w.Header().Add("Location", "/virgin?error=t")
		w.WriteHeader(302)
		return
	}

	t := Parse("virgin-show.html")

	t.Execute(w, struct {
		PNR              delta.PNR
		ConfirmationCode string
		CommitHash       string
	}{
		retrievedPNR,
		confirmationCode,
		commitHash,
	})
}
