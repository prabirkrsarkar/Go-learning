package main

import (
	"fmt"
	"log"
	"strings"
	"text/template"
)

type EmailData struct {
	RecipientName string
	Subject       string
	Body          string
	Attachments   []string
	SenderName    string
	UnreadCount   int
}

func main() {
	fmt.Println("----------Email Template----------")

	// The dot operator means get the current data context, current object or values. Similar to 'this' in other languages,
	// It is used to access fields within the EmailData struct.
	emailTemplate := `
		To: {{.RecipientName}}
		Subject: {{.Subject}}
		Body: {{.Body}}

		{{if .Attachments}}
		Attachments: {{range .Attachments}}{{.}}, 
		{{end}} 
		{{end}}
		
		{{if gt .UnreadCount 0}}
		You have {{.UnreadCount}} unread messages.
		{{else}} 
		You have no messages.
		{{end}}
		
		Thanks
		{{.SenderName}}
    `
	// Create a new template object and parse the email template string
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	data := EmailData{
		RecipientName: "John Doe",
		Subject:       "Welcome!",
		Body:          "Thanks for joining our platform.",
		Attachments:   []string{"file1.txt", "file2.jpg"},
		SenderName:    "Prabir Sarkar",
		UnreadCount:   5,
	}

	// Execute func - func (t *Template) Execute(wr io.Writer, data any) error {...}
	// and io.Writer is an interface that defines a method Write(p []byte) (n int, err error).
	// type Writer interface { Write(p []byte) (n int, err error)
	// This means that any type that implements this Write()method can be used as an io.Writer.
	// os.Stdout is the standard output stream, which implements io.Writer, so we can use it to print the output directly to the console.
	// Other examples of io.Writer include os.File for writing to files, bytes.Buffer for writing to an in-memory buffer,
	// and http.ResponseWriter for writing HTTP responses in web applications.
	// Also, the string.Builder type implements io.Writer, allowing you to build strings efficiently by writing to a buffer.

	// err = tmpl.Execute(os.Stdout, data)

	// If you want to capture the output of the template execution into a string instead of printing it directly to the console,
	// you can use a strings.Builder as an in-memory buffer.
	// The strings.Builder type implements the io.Writer interface, allowing you to write the output of the template execution to it and then convert it to a string.
	var output strings.Builder

	// Because tmpl.Execute() requires an io.Writer interface, and strings.Builder type implements io.Writer
	// only through its pointer receiver methods, we need to pass a pointer to the strings.Builder instance
	// when calling tmpl.Execute().
	err = tmpl.Execute(&output, data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(strings.ToUpper(output.String()))

}
