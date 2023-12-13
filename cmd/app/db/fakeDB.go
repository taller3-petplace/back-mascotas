package db

import (
	"errors"
	"petplace/back-mascotas/cmd/app/data"
)

type FakeDB struct {
	data            map[int]data.Pet
	lastIdGenerated int
}

func NewFakeDB() FakeDB {
	return FakeDB{data: map[int]data.Pet{}, lastIdGenerated: 0}
}

func (fdb *FakeDB) Save(pet *data.Pet) error {

	id := fdb.newID()
	pet.ID = id
	fdb.data[id] = *pet

	return nil
}

func (fdb *FakeDB) Get(id int) (data.Pet, error) {

	item, ok := fdb.data[id]
	if !ok {
		return data.Pet{}, errors.New("not found")
	}

	return item, nil
}

func (fdb *FakeDB) GetByOwner(OwnerID int) ([]data.Pet, error) {

	var result []data.Pet
	for _, value := range fdb.data {
		if value.OwnerID == OwnerID {
			result = append(result, value)
		}
	}
	return result, nil
}

func (fdb *FakeDB) Delete(id int) {
	delete(fdb.data, id)
}

func (fdb *FakeDB) newID() int {
	fdb.lastIdGenerated++
	return fdb.lastIdGenerated
}
