package filemaker

import "github.com/jomla97/filemaker/internal/connection"

//Connect starts a database session
func Connect(host string, database string, username string, password string) (*connection.Connection, error) {
	return connection.New(host, database, username, password)
}
