package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/yosssi/gohtml"
)

// Data is Top level wrapper type to specifiy data variables
// which are injected in templates
type Data struct {
	User    string
	Heading string
	Links   string
	Content string
	Footer  string
}

// PageLink is link at top of page
type PageLink struct {
	FileName string
	Title    string
}

// Meta is for each Block
type Meta struct {
	Name    string
	Heading string
}

// Block is abstract type
type Block interface {
	getMeta() Meta
	getContent() string
}

// FooterLink is a link in page footer
type FooterLink struct {
	Title string
	URL   string
}

// ProjectsBlock is for Projects page
type ProjectsBlock struct {
	Meta
	Projects []Project
}

// Project is a single project
type Project struct {
	Title string
	Desc  string
	Link  string
}

func (pb *ProjectsBlock) getMeta() Meta {
	return pb.Meta
}

func (pb *ProjectsBlock) getContent() string {
	return compileTemplate("tmpl/projects_tmpl.html", struct {
		Projects []Project
	}{
		Projects: pb.Projects,
	})
}

// AboutBlock is for about page
type AboutBlock struct {
	Meta
	Paragraphs []string
}

func (ab *AboutBlock) getMeta() Meta {
	return ab.Meta
}

func (ab *AboutBlock) getContent() string {
	return compileTemplate("tmpl/about_tmpl.html", struct {
		Paragraphs []string
	}{
		Paragraphs: ab.Paragraphs,
	})
}

// InmemoryWriter holds compiled template strings
type InmemoryWriter struct {
	Value []byte
}

func (w *InmemoryWriter) Write(p []byte) (n int, err error) {
	w.Value = append(w.Value, p...)
	return len(p), nil
}

func (w *InmemoryWriter) toString() string {
	return string(w.Value)
}

const mainTmpl = "tmpl/main_tmpl.html"

func main() {
	// UserName, Blocks, FooterLinks is defined in content.go

	links := generateLinksTmpl(Blocks)
	footer := generateFooterTmpl(FooterLinks)

	for _, b := range Blocks {
		fname := fmt.Sprintf("public/%s.html", b.getMeta().Name)
		heading := b.getMeta().Heading
		content := b.getContent()

		html := compileTemplate(mainTmpl, Data{
			User:    UserName,
			Heading: heading,
			Links:   links,
			Content: content,
			Footer:  footer,
		})
		f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		gohtml.Condense = true // enable inline tags
		formatted := gohtml.Format(html)
		f.Write([]byte(formatted))
		f.Close()
	}
}

func generateLinksTmpl(blocks []Block) string {
	links := make([]PageLink, 0)
	for _, b := range blocks {
		links = append(links, PageLink{
			Title:    b.getMeta().Heading,
			FileName: fmt.Sprintf("%s.html", b.getMeta().Name),
		})
	}

	return compileTemplate("tmpl/nav_links_tmpl.html",
		struct {
			Links []PageLink
		}{
			Links: links,
		},
	)
}

func generateFooterTmpl(links []FooterLink) string {
	return compileTemplate("tmpl/footer_tmpl.html",
		struct {
			Links []FooterLink
		}{
			Links: links,
		},
	)
}

func compileTemplate(templateName string, data interface{}) string {
	w := InmemoryWriter{}
	tmpl := template.Must(template.ParseFiles(templateName))
	tmpl.Execute(&w, data)
	return w.toString()
}
