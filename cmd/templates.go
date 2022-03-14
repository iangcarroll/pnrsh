package main

import (
	"embed"
	"io"
	"log"
	"text/template"
)

type safeTemplate struct {
	t *template.Template
}

// Safely execute a template and panic on failure.
func (s *safeTemplate) Execute(wr io.Writer, data interface{}) {
	if err := s.t.Execute(wr, data); err != nil {
		log.Fatal(err)
	}
}

var (
	//go:embed templates/*
	templates embed.FS
)

// Safely parse a template and panic on failure.
func Parse(view string) *safeTemplate {
	t, err := template.ParseFS(
		templates,
		"templates/"+view,
		"templates/base.html",
	)

	if err != nil {
		log.Fatal(err)
	}

	return &safeTemplate{t}
}
