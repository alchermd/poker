package server

import "html/template"

// Base template
const tplRoot = "templates"
const tplBase = "base.gohtml"

// Partial templates
const tplPartialRoot = tplRoot + "/partials"
const tplPartialNavbar = "navbar.gohtml"
const tplPartialLogo = "logo.gohtml"

func (s *Server) ParseTemplates() {
	tplPartials := []string{
		tplPartialRoot + "/" + tplPartialNavbar,
		tplPartialRoot + "/" + tplPartialLogo,
	}
	tpls := []string{
		tplRoot + "/" + tplBase,
		tplRoot + "/" + tplHome,
	}
	tpls = append(tplPartials, tpls...)
	s.templates = template.Must(template.ParseFiles(tpls...))
}
