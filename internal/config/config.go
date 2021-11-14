package config

import (
	"github.com/AshRoy29/booking/internal/models"
	"github.com/alexedwards/scs/v2"
	"log"
	"text/template"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
