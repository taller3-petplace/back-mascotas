package db

type Storable interface {
	Save(data interface{}) error
	Get(id int, data interface{}) error
	Delete(id int, data interface{}) error
	GetFiltered(filter func(interface{}) bool) error
}
