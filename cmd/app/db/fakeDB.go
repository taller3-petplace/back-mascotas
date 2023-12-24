package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"petplace/back-mascotas/cmd/app/model"
)

const testDataFile = "cmd/app/db/testdata.json"

type FakeDB struct {
	data            map[int]model.Pet
	lastIdGenerated int
}

func NewFakeDB() FakeDB {
	return FakeDB{data: map[int]model.Pet{}, lastIdGenerated: 0}
}

func (fdb *FakeDB) Init() {

	pets, err := loadTestData(testDataFile)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic("Error trying to read test model: " + err.Error())
	}

	for _, pet := range pets {
		_ = fdb.Save(&pet)
	}
}

func (fdb *FakeDB) Save(pet *model.Pet) error {

	id := fdb.newID()
	pet.ID = id
	fdb.data[id] = *pet

	err := dumpToFile(testDataFile, fdb.data)
	if err != nil {
		fmt.Println("ERROR DUMP: " + err.Error())
	}

	return nil
}

func (fdb *FakeDB) Get(id int) (model.Pet, error) {

	item, ok := fdb.data[id]
	if !ok {
		return model.Pet{}, errors.New("not found")
	}

	return item, nil
}

func (fdb *FakeDB) GetByOwner(OwnerID int) ([]model.Pet, error) {

	var result []model.Pet
	for _, value := range fdb.data {
		if value.OwnerID == OwnerID {
			result = append(result, value)
		}
	}
	return result, nil
}

func (fdb *FakeDB) Delete(id int) {
	delete(fdb.data, id)

	err := dumpToFile(testDataFile, fdb.data)
	if err != nil {
		fmt.Println("ERROR DUMP: " + err.Error())
	}

}

func (fdb *FakeDB) newID() int {
	fdb.lastIdGenerated++
	return fdb.lastIdGenerated
}

func loadTestData(filename string) ([]model.Pet, error) {

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var petsList []model.Pet
	err = json.Unmarshal(fileContent, &petsList)
	if err != nil {
		return nil, err
	}

	return petsList, nil
}

func dumpToFile(filename string, petsMap map[int]model.Pet) error {

	var petList []model.Pet
	for _, pet := range petsMap {
		petList = append(petList, pet)
	}

	jsonData, err := json.MarshalIndent(petList, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
