package filemaker

import (
	"github.com/jomla97/go-filemaker/pkg/session"
)

//New starts a database session
func New(host string, database string, username string, password string) (*session.Session, error) {
	return session.New(host, database, username, password)
}
