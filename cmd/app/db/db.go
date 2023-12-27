package db

type StorableItem struct {
	ID   int
	Data interface{}
}

type Storable interface {
	NewID() int
	Save(table string, data interface{}) (*StorableItem, error)
	Get(table string, id int) (*StorableItem, error)
	Delete(table string, id int)
	GetFiltered(table string, filter func(StorableItem) bool) ([]StorableItem, error)
}
