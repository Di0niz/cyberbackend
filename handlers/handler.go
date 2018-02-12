package handlers

import (
	"github.com/Di0niz/cyberbackend/config"
)

type (
	Handler struct {
		DB *config.DBPool
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
