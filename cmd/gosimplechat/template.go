package main

import (
	"github.com/utrack/go-simple-chat/assets"
	"html/template"
)

func getTemplate() (*template.Template, error) {
	buf, err := assets.Asset("assets/static/chat.tmpl")
	if err != nil {
		return nil, err
	}
	return template.New("chat").Parse(string(buf))

}
