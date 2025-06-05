package server

import "html/template"

const tplRoot = "templates"
const tplBase = "base.gohtml"

func (s *Server) ParseTemplates() {
	templateNames := []string{
		tplRoot + "/" + tplBase,
		tplRoot + "/" + tplHome,
	}
	s.templates = template.Must(template.ParseFiles(templateNames...))
}
