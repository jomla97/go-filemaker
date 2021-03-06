package session

import "errors"

//Resume resumes a database session with the specified token
func Resume(host string, database string, username string, password string, token string) (*Session, error) {
	if host == "" {
		return nil, errors.New("No host specified")
	} else if database == "" {
		return nil, errors.New("No database specified")
	} else if username == "" {
		return nil, errors.New("No username specified")
	}

	//Determine protocol scheme
	var protocol = "https://"
	if len(host) >= 8 && host[0:8] == "https://" {
		protocol = ""
	}

	return &Session{
		Token:    token,
		Protocol: protocol,
		Host:     host,
		Database: database,
		Username: username,
		Password: password,
	}, nil
}
