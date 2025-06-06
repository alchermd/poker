package server

import "html/template"

type templateCache map[string]*template.Template

// Base template
const tplRoot = "templates"
const tplBase = "base.gohtml"

// Partial templates
const tplPartialRoot = tplRoot + "/partials"
const tplPartialNavbar = "navbar.gohtml"
const tplPartialLogo = "logo.gohtml"

// ParseTemplates loads the templates and caches them in-memory.
func (s *Server) ParseTemplates() {
	// Partial templates that are loaded for all templates.
	tplPartials := []string{
		tplPartialRoot + "/" + tplPartialNavbar,
		tplPartialRoot + "/" + tplPartialLogo,
	}
	// List of templates to load and cache.
	tpls := []string{
		tplHome,
		tplCreateGame,
	}
	// For each template, we combine the partials and the template (+ base) and
	//
	for _, tpl := range tpls {
		tplWithBase := []string{
			tplRoot + "/" + tplBase,
			tplRoot + "/" + tpl,
		}
		tpls = append(tplPartials, tplWithBase...)
		// Each template is key-accessible by its template name.
		// Ex: templates/home.gohtml -> home.gohtml
		s.templates[tpl] = template.Must(template.ParseFiles(tpls...))
	}
}
