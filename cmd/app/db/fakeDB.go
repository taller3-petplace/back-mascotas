package db

import (
	"errors"
	"petplace/back-mascotas/cmd/app/data"
)

type FakeDB struct {
	data map[string]data.Pet
}

func NewFakeDB() FakeDB {
	return FakeDB{data: map[string]data.Pet{}}
}

func (fdb FakeDB) Save(id string, pet data.Pet) error {

	fdb.data[id] = pet

	return nil
}

func (fdb FakeDB) Get(id string) (data.Pet, error) {

	item, ok := fdb.data[id]
	if !ok {
		return data.Pet{}, errors.New("not found")
	}

	return item, nil
}

func (fdb FakeDB) Delete(id string) error {
	delete(fdb.data, id)
	return nil
}
