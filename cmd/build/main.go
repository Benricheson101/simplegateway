package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"text/template"
	"time"
)

// Build script inspired by https://github.com/bwmarrin/discordgo/blob/master/tools/cmd/eventhandlers/main.go

const EVENTS_TEMPLATE = `
// DO NOT EDIT
// Generated at {{timestamp}}

package gateway

var EventNames = []string {
	{{range .Events}}"{{camelToScreamingSnake .}}",
	{{end}}
}

func execHandlerFunc(gw *Gateway, handlers []interface{}, t string, payload interface{}) {
	switch t {
		{{range .Events }}case "{{camelToScreamingSnake .}}":
				for _, fn := range handlers {
					go fn.(func(*Gateway, *{{.}}))(gw, payload.(*{{.}}))
				}
		{{end}}
	}
}

func eventNameToPayload(t string) interface{} {
	switch t {
	{{range .Events}}case "{{camelToScreamingSnake .}}": return &{{.}}{}
	{{end}}default: return nil
	}
}
`

func main() {
	eventTemplate := template.Must(template.New("events").
		Funcs(template.FuncMap{
			"camelToScreamingSnake": camelToScreamingSnake,
			"timestamp":             timestamp,
		}).
		Parse(EVENTS_TEMPLATE))

	fs := token.NewFileSet()
	p, err := parser.ParseFile(fs, "./pkg/gateway/events.go", nil, parser.DeclarationErrors)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse events.go file: %v\n", err)
		os.Exit(1)
	}

	var events []string

	for name, obj := range p.Scope.Objects {
		if obj.Kind != ast.Typ {
			continue
		}

		events = append(events, name)
	}

	var buf bytes.Buffer
	eventTemplate.Execute(&buf, struct {
		Events []string
	}{
		Events: events,
	})

	fmtd, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to format generated code: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile("./pkg/gateway/eventhandler.go", fmtd, 0664)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file: %v\n", err)
		os.Exit(1)
	}
}

func camelToScreamingSnake(s string) string {
	var buf bytes.Buffer

	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(c)
		} else {
			buf.WriteRune(c + 'A' - 'a')
		}
	}

	return buf.String()
}

func timestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}
