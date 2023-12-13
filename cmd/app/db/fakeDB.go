package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"petplace/back-mascotas/cmd/app/data"
)

const testDataFile = "testdata.json"

type FakeDB struct {
	data            map[int]data.Pet
	lastIdGenerated int
}

func NewFakeDB() FakeDB {
	return FakeDB{data: map[int]data.Pet{}, lastIdGenerated: 0}
}

func (fdb *FakeDB) Init() {

	pets, err := loadTestData(testDataFile)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic("Error trying to read test data: " + err.Error())
	}

	for _, pet := range pets {
		_ = fdb.Save(&pet)
	}
}

func (fdb *FakeDB) Save(pet *data.Pet) error {

	id := fdb.newID()
	pet.ID = id
	fdb.data[id] = *pet

	err := dumpToFile(testDataFile, fdb.data)
	if err != nil {
		fmt.Println("ERROR DUMP: " + err.Error())
	}

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

	err := dumpToFile(testDataFile, fdb.data)
	if err != nil {
		fmt.Println("ERROR DUMP: " + err.Error())
	}

}

func (fdb *FakeDB) newID() int {
	fdb.lastIdGenerated++
	return fdb.lastIdGenerated
}

func loadTestData(filename string) ([]data.Pet, error) {

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var petsList []data.Pet
	err = json.Unmarshal(fileContent, &petsList)
	if err != nil {
		return nil, err
	}

	return petsList, nil
}

func dumpToFile(filename string, petsMap map[int]data.Pet) error {

	var petList []data.Pet
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
