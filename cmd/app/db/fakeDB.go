package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const testDataFile = "cmd/app/db/testdata.json"

type FakeDB struct {
	data            map[string]map[int]StorableItem
	lastIdGenerated int
}

func NewFakeDB() FakeDB {
	return FakeDB{data: map[string]map[int]StorableItem{}, lastIdGenerated: 0}
}

func (fdb *FakeDB) Init() {

	db, err := loadTestData(testDataFile)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic("Error trying to read test model: " + err.Error())
	}

	if db != nil {
		fdb.data = db
		lastID := 0
		for _, table := range fdb.data {
			for _, item := range table {
				lastID = max(lastID, item.ID)
			}
		}
		fdb.lastIdGenerated = lastID
	}
}

func (fdb *FakeDB) Save(table string, data interface{}) (*StorableItem, error) {

	fdb.checkTable(table)

	item := fdb.getStorableItem(data)
	fdb.data[table][item.ID] = item

	err := dumpToFile(testDataFile, fdb.data)
	if err != nil {
		fmt.Println("ERROR DUMP: " + err.Error())
	}

	return &item, nil
}

func (fdb *FakeDB) Get(table string, id int) (*StorableItem, error) {

	fdb.checkTable(table)

	item, ok := fdb.data[table][id]
	if !ok {
		return nil, errors.New("not found")
	}

	return &item, nil
}

func (fdb *FakeDB) GetFiltered(table string, filter func(item StorableItem) bool) ([]StorableItem, error) {

	var result []StorableItem
	for _, item := range fdb.data[table] {

		if filter(item) {
			result = append(result, item)
		}
	}
	return result, nil
}

func (fdb *FakeDB) Delete(table string, id int) {

	fdb.checkTable(table)

	delete(fdb.data[table], id)

	err := dumpToFile(testDataFile, fdb.data)
	if err != nil {
		fmt.Println("ERROR DUMP: " + err.Error())
	}

}

func (fdb *FakeDB) NewID() int {
	fdb.lastIdGenerated++
	return fdb.lastIdGenerated
}

func (fdb *FakeDB) checkTable(table string) {
	_, ok := fdb.data[table]
	if !ok {
		fdb.data[table] = map[int]StorableItem{}
	}
}

func loadTestData(filename string) (map[string]map[int]StorableItem, error) {

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data map[string]map[int]StorableItem
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func dumpToFile(filename string, data map[string]map[int]StorableItem) error {

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fdb *FakeDB) getStorableItem(data interface{}) StorableItem {

	switch v := data.(type) {
	case StorableItem:
		return v
	}

	return StorableItem{
		ID:   fdb.NewID(),
		Data: data,
	}
}
