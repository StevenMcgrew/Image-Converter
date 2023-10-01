package main

import (
	"bytes"
	"context"
	"encoding/json"
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

const STORE_PATH = "store.json"

type UserPreferences struct {
	FolderPaths         []string `json:"folderPaths"`
	ConvertTo           string   `json:"convertTo"`
	ResizeWidth         int      `json:"resizeWidth"`
	ResizeHeight        int      `json:"resizeHeight"`
	MaintainAspectRatio bool     `json:"maintainAspectRatio"`
	FillUnusedSpace     bool     `json:"fillUnusedSpace"`
}

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

func getUserPreferences(path string) UserPreferences {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	userPreferences := UserPreferences{}
	json.Unmarshal(bytes, &userPreferences)
	// figure out how to check if we get some json data,
	// if no json data, handle in some way
	return userPreferences
}

func (a *App) SaveUserPreferences(userPreferences UserPreferences) {
	file, err := os.Create(STORE_PATH)
	if err != nil {
		fmt.Println("ERROR on os.Create:", err)
		return
	}
	defer file.Close()

	bytes, err := json.MarshalIndent(userPreferences, "", "\t")
	if err != nil {
		fmt.Println("ERROR on json.Marshal: ", err)
		return
	}
	fmt.Fprint(file, string(bytes))
}

func (a *App) GetHomePage() string {
	userPref := getUserPreferences(STORE_PATH)
	return render("html/pages/home.html", userPref)
}
