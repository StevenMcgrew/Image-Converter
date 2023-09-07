package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"
)

/*
*************************************************
Wails boilerplate
*************************************************
*/
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

/*
*************************************************
Your code below here
*************************************************
*/
func render(path string, data any) string {
	if data != nil {
		var buf bytes.Buffer
		tmpl := template.Must(template.ParseFiles(path))
		tmpl.Execute(&buf, data)
		return buf.String()
	} else {
		bytes, err := os.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		}
		return string(bytes)
	}
}

func (a *App) GetHomePage() string {
	return render("html/pages/home.html", nil)
}
