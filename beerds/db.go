package beerds

import (
	"appengine"
	"strings"
	"net/http"
)

type Database struct{
	namespace string
	context appengine.Context
}

func (db *Database) Table(name string) (*Table, error) {
	return &Table {
		plural: name,
		singular: name[0:len(name) - 1],
		database: db,
	}, nil
}

func (db *Database) SetContext(req *http.Request) {
	context := appengine.NewContext(req)
	db.context = context
}

func toLower(in string) string {
	return strings.ToLower(string(in[0])) + in[1:]
}

func toUpper(in string) string {
	return strings.ToUpper(string(in[0])) + in[1:]
}