package db

type Storable interface {
	Save(data interface{}) error
	Get(id int, data interface{}) error
	Delete(id int, data interface{}) error
	GetFiltered(result interface{}, filter map[string]string, orderBy string, limit int, offset int) (int64, error)
}
