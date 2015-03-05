package beerds

import (
	"appengine"
	"net/http"
	"strings"
)

type Database struct {
	namespace string
	context   appengine.Context
}

func (db *Database) Table(name string) (*Table, error) {
	return &Table{
		plural:   name,
		singular: name[0 : len(name)-1],
		database: db,
	}, nil
}

func (db *Database) SetContext(req *http.Request) {
	context := appengine.NewContext(req)
	//namespaceContext, err := appengine.Namespace(context, urlToNamespace(req.URL.Host))
	namespaceContext, err := appengine.Namespace(context, "0.0.0.0-4200")
	if err != nil {
		panic(err.Error())
	}
	db.context = namespaceContext
}

func (db *Database) GetContext() appengine.Context {
	return db.context
}

func toLower(in string) string {
	return strings.ToLower(string(in[0])) + in[1:]
}

func toUpper(in string) string {
	return strings.ToUpper(string(in[0])) + in[1:]
}

func urlToNamespace(url string) string {
	return strings.Replace(url, ":", "-", -1)
}
