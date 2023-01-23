package config

import (
	"fmt"
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

func (config *AppConfig) AppConfigToString() string {
	return fmt.Sprintf("UseCache: %v, InfoLog: %v, InProduction: %v", config.UseCache, config.InfoLog, config.InProduction)
}
