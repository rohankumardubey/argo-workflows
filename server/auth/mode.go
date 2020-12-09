package auth

import (
	"errors"
	"strings"

	"github.com/argoproj/argo/server/auth/sso"
)

type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"
)

func (m Modes) Add(value string) error {
	switch value {
	case "client", "server", "sso":
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true
		m[Server] = true
	default:
		return errors.New("invalid mode")
	}
	return nil
}

func (m Modes) GetMode(authorisation string) (Mode, bool) {
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, m[SSO]
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, m[Client]
	}
	return Server, m[Server]
}