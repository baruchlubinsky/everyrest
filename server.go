package main

import (
	"beerds"
	"github.com/baruchlubinsky/beerapi/api"
	"log"
	"net/http"
	"strings"
)

var Db *beerds.Database

func init() {
	Db = &beerds.Database{}
	http.HandleFunc("/", beer)
}

func beer(response http.ResponseWriter, request *http.Request) {
	header := response.Header()
	header.Set("Content-Type", "application/json")
	// CORS
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "POST, PUT, DELETE, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, content-type, Accept, X-AUTH-TOKEN, X-API-VERSION")
	// Check for table
	table, err := tableFor(request)
	if err != nil {
		response.Write([]byte("Error while creating appengine context:\n" + err.Error()))
		response.WriteHeader(500)
		return
	}
	switch request.Method {
	case "POST":
		api.Post(table, response, request)
	case "GET":
		api.Get(table, response, request)
	case "PUT":
		api.Put(table, response, request)
	case "DELETE":
		api.Delete(table, response, request)
	case "OPTIONS":
		response.WriteHeader(200)
	default:
		response.WriteHeader(400)
	}
}

func tableFor(request *http.Request) (*beerds.Table, error) {
	Db.SetContext(request)
	args := strings.Split(strings.Trim(request.URL.Path, "/"), "/")
	if len(args) == 0 {
		return nil, ServerError("Must provide a resource name.")
	}
	name := args[0]
	table, err := Db.Table(name)
	return table, err
}

type ServerError string

func (s ServerError) Error() string {
	return string(s)
}
