package beerds

import(
	"github.com/baruchlubinsky/beerapi/adapters"
	"appengine/datastore"
	"strconv"
)

type Table struct{
	singular, plural string
	database *Database
}

func (table *Table) Find(id string) (adapters.Model, error) {
	entity := make(datastore.PropertyList, 0)
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	key := datastore.NewKey(table.database.context, table.EntityName(), "", int64(intId), nil)
	err = datastore.Get(table.database.context, key, &entity)
	if err != nil {
		return nil, err
	}
	return &Model{
		entity: &entity,
		key: key,
		table: table,
	}, nil
}
	
// If query == nil, return entire contents of table.
func (table *Table) Search(query interface{}) (adapters.ModelSet)  {
	q := datastore.NewQuery(table.EntityName())
	if query != nil {
		// TODO
	}
	res := make(adapters.ModelSet, 0)
	results := q.Run(table.database.context)
	for true {
		record := make(datastore.PropertyList, 0)
		key, err := results.Next(&record)
		if err == datastore.Done {
			break
		}
		model := Model{
			key: key,
			entity: &record,
			table: table,
		}
		res.Add(&model)
	}
	return res
}
func (table *Table) NewRecord() (adapters.Model) {
	entity := make(datastore.PropertyList, 0)
	return &Model{
		entity: &entity,
		key: datastore.NewIncompleteKey(table.database.context, table.EntityName(), nil),
		table: table,
	}
}
func (table *Table) Delete(id string) (error) {
	k := datastore.NewKey(table.database.context, table.EntityName(), id, 0, nil)
	return datastore.Delete(table.database.context, k)
}

func (table *Table) EntityName() (string) {
	return toUpper(table.RecordName())
}

// The name of an individual record.
func (table *Table) RecordName() string {
	return table.singular
}

// The name of a collection of records.
func (table *Table) RecordSetName() string {
	return table.plural
}