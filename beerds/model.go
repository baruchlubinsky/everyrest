package beerds

import (
	"appengine/datastore"
	"reflect"
)

type Model struct {
	entity *datastore.PropertyList
	key    *datastore.Key
	table  *Table
}

func (model *Model) GetId() string {
	return model.key.String()[len(model.table.EntityName())+2:]
}

func (model *Model) Attributes() map[string]interface{} {
	res := make(map[string]interface{})
	for _, property := range *model.entity {
		if property.Multiple {
			data, found := res[toLower(property.Name)]
			if found {
				res[toLower(property.Name)] = append(data.([]interface{}), property.Value)
			} else {
				res[toLower(property.Name)] = []interface{}{property.Value}
			}
		} else {
			res[toLower(property.Name)] = property.Value
		}
	}
	res["id"] = model.GetId()
	return res
}

func (model *Model) SetAttributes(data map[string]interface{}) {
	res := make(datastore.PropertyList, 0)
	relations := make(map[string]interface{})
	for _, property := range *model.entity {
		field, found := data[toLower(property.Name)]
		if found {
			if addProperty(&res, property.Name, field) || property.Multiple {
				relations[property.Name] = true
			}
			delete(data, toLower(property.Name))
		} else {
			if _, slice := relations[property.Name]; !slice {
				res = append(res, property)
			}
		}
	}
	for key, value := range data {
		addProperty(&res, toUpper(key), value)
	}
	model.entity = &res
}
func (model *Model) Save() error {
	key, err := datastore.Put(model.table.database.context, model.key, model.entity)
	if err != nil {
		return err
	}
	model.key = key
	return nil
}

func (model *Model) Delete() error {
	return datastore.Delete(model.table.database.context, model.key)
}

func addProperty(destination *datastore.PropertyList, name string, field interface{}) bool {
	t := reflect.TypeOf(field).Kind()
	if t == reflect.Slice {
		for _, elem := range field.([]interface{}) {
			*destination = append(*destination, datastore.Property{
				Name:     name,
				Value:    elem,
				NoIndex:  false,
				Multiple: true,
			})
		}
		return true
	} else {
		*destination = append(*destination, datastore.Property{
			Name:     name,
			Value:    field,
			NoIndex:  false,
			Multiple: false,
		})
		return false
	}
}
