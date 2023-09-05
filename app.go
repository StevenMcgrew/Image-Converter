package main

import (
	"bytes"
	"context"
	"html/template"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func renderFile(path string, data any) string {
	var buf bytes.Buffer
	tmpl := template.Must(template.ParseFiles(path))
	tmpl.Execute(&buf, data)
	return buf.String()
}

type Message struct {
	Text string
}

func (a *App) ReplaceLogo(message Message) string {
	return renderFile("htmlTemplates/message.html", message)
}
