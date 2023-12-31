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
func (r *Repository) GetFiltered(result interface{}, filters map[string]string, orderBy string, limit int, offset int) (int64, error) {
	// Obtener la cantidad total de filas sin limitación
	var totalCount int64
	queryCount := r.db.Model(result)
	queryCount = applyFilters(queryCount, filters)
	err := queryCount.Count(&totalCount).Error
	if err != nil {
		return 0, err
	}

	// Construir la consulta principal
	query := r.db.Model(result)
	query = applyFilters(query, filters)

	// Aplicar criterio de ordenamiento
	if orderBy != "" {
		query = query.Order(orderBy)
	}

	// Aplicar límite y desplazamiento
	query = query.Limit(limit).Offset(offset)

	// Ejecutar la consulta principal
	if err := query.Find(result).Error; err != nil {
		return 0, err
	}

	return totalCount, nil

}

// applyFilters aplica los filtros a la consulta
func applyFilters(query *gorm.DB, filters map[string]string) *gorm.DB {
	for campo, valor := range filters {
		query = query.Where(campo, valor)
	}
	return query
}
