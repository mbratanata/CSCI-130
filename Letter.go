package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	const letter = `
Dear {{.Type}} {{.LName}},
{{if .Attend}}
Thank you so much for attending the fundraiser. It is a real pleasure to meet you at the event.{{else}}
I am sorry that you could not attend the fundraiser.{{end}}
{{if .Donate}}
Thank you for your big consideration to donate. I greatly appreciate it.
{{end}}
I would like to remind you of our upcoming events:
{{range .Events}} {{.}}
{{end}}
Best wishes,
Mr. Bratanata
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Type, LName string
		Attend   bool
		Donate	bool
		Events string
	}
	var recipients = []Recipient{
		{"Mr.", "Smith", true},
		{"Mrs.", "Brown", false},
		{"Ms.", "Darla", false},
	}

	var upcomingEvents = []string{"Event A", "Event B", "Event C"}

	// STEP 1 & STEP 2
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	//STEP 3
	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}