package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// dsn := "admin:admin@tcp(localhost:3306)/pets"

type Repository struct {
	url string
	db  *gorm.DB
}

func NewRepository(url string) (Repository, error) {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return Repository{}, err
	}

	return Repository{url: url, db: db}, nil
}

func (r *Repository) Init(models []interface{}) error {
	return r.db.AutoMigrate(models...)
}

func (r *Repository) Save(data interface{}) error {

	result := r.db.Save(data)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
func (r *Repository) Get(id int, object interface{}) error {

	result := r.db.First(object, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *Repository) Delete(id int, object interface{}) error {
	result := r.db.Delete(object, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *Repository) GetFiltered(filter func(interface{}) bool) error {
	return nil
}
