package filemaker

import (
	"github.com/jomla97/go-filemaker/pkg/errortypes"
	"github.com/jomla97/go-filemaker/pkg/record"
	"github.com/jomla97/go-filemaker/pkg/session"
)

//Record represents the record object returned by performed findcommands
type Record = record.Record

//FieldChange represents the object returned by record.SetField()
type FieldChange = record.FieldChange

//Session represents the session struct returned by filemaker.New()
type Session = session.Session

//----------- Errors -----------

//ErrorNotFound is returned by Session.PerformFind() when no records match the request
type ErrorNotFound = errortypes.ErrorNotFound
