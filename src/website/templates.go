package main

import (
	"embed"
	"fmt"
	"html/template"
	"sync"
)

//go:embed templates/*
var templateFs embed.FS

var templates *template.Template
var templateCache = make(map[string]*template.Template)
var cacheMutex sync.Mutex

func GetTemplate(name string) (templ *template.Template, err error) {
	if templ = templateCache[name]; templ != nil {
		return
	}

	templ, err = template.ParseFS(templateFs, "templates/layout.html", fmt.Sprintf("templates/pages/%s/*.html", name))
	cacheMutex.Lock()
	templateCache[name] = templ
	cacheMutex.Unlock()
	return
}
