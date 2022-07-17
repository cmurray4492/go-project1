package config

import (
	"html/template"
	"log"
)

// AppConfig holds application config data points
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
