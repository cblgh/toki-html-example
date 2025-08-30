package main

import (
	myHTML "eyeneighteenn/html"
	"eyeneighteenn/translations"
	// toki deps
	"eyeneighteenn/tokibundle"
	"golang.org/x/text/language"

	// example deps
	"errors"
	"fmt"
	"flag"
	"time"
	"log"
	"net/http"
	"syscall"
	"html/template"
	
)
var views = []string{"index"}
func generateTemplates(translator tokibundle.Reader) (*template.Template, error) {
	rootTemplate := template.New("root")
	for _, view := range views {
		newTemplate, err := rootTemplate.ParseFS(myHTML.Templates, fmt.Sprintf("%s.html", view))
		if err != nil {
			return nil, fmt.Errorf("could not get files: %w", err)
		}
		rootTemplate = newTemplate
	}

	return rootTemplate, nil
}

// format all errors consistently, and provide context for the error using the string `msg`
func Eout(err error, msg string, args ...interface{}) error {
	if err != nil {
		// received an invocation of e.g. format:
		// Eout(err, "reading data for %s and %s", "database item", "weird user")
		if len(args) > 0 {
			return fmt.Errorf("%s (%w)", fmt.Sprintf(msg, args...), err)
		}
		return fmt.Errorf("%s (%w)", msg, err)
	}
	return nil
}

func Check(err error, msg string, args ...interface{}) {
	if len(args) > 0 {
		err = Eout(err, msg, args...)
	} else {
		err = Eout(err, msg)
	}
	if err != nil {
		log.Fatalln(err)
	}
}

type IndexData struct{
	Translator tokibundle.Reader
	Time time.Time
	NeutralName string
	Action string
}

type Handler struct {
	templates  *template.Template
	Translator tokibundle.Reader
}

func (h Handler) renderView(res http.ResponseWriter, viewName string, data IndexData) {
	view := fmt.Sprintf("%s.html", viewName)
	if err := h.templates.ExecuteTemplate(res, view, data); err != nil {
		if errors.Is(err, syscall.EPIPE) {
			fmt.Println("recovering from broken pipe")
			return
		} else {
			Check(err, "rendering %q view", view)
		}
	}
}

func (h Handler) indexRoute (res http.ResponseWriter, req *http.Request) {
	h.renderView(res, "index", IndexData{Translator: h.Translator, Time: time.Now(), NeutralName: "Frank", Action: h.Translator.String("Bake bread") })
}

// NOTE: since toki doesn't currently parse {{ .GoFunction "param" }} inside .html files, i have to expose all the keys
// somewhere in a go file. see my lo-fi hacky way of doing that in package `translations`

func main() {
	var port int
	var lang string
	flag.IntVar(&port, "port", 8292, "port")
	flag.StringVar(&lang, "lang", "en", "website language")
	flag.Parse()

	chosenLang := language.BritishEnglish

	if lang == "sv" {
		chosenLang = language.Swedish
	}

	reader, _ := tokibundle.Match(chosenLang)

	translations.Output(myHTML.Templates, views)

	templates := template.Must(generateTemplates(reader))
	handler := Handler{templates: templates, Translator: reader}
	http.HandleFunc("/", handler.indexRoute)

	portstr := fmt.Sprintf(":%d", port)
	fmt.Println("Listening on port: ", portstr)
	http.ListenAndServe(portstr, nil)
}
