package mail

import (
	"embed"
	"html/template"
	"log"
)

//go:embed templates/*.html
var templateFiles embed.FS
var verifyCodeTmpl *template.Template

func init() {
	var err error
	verifyCodeTmpl, err = template.ParseFS(templateFiles, "templates/verify_email.html")
	if err != nil {
		log.Fatal(err)
	}
}
