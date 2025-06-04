package server

import "html/template"

const tplRoot = "templates"

func (s *Server) ParseTemplates() {
	templateNames := []string{
		tplRoot + "/" + tplHome,
	}
	s.templates = template.Must(template.ParseFiles(templateNames...))
}
