package tests

import (
	"log"
	"testing"
	"time"

	"github.com/informeai/dbinfos/adapters"
	"github.com/informeai/dbinfos/config"
)

func TestNewDBInfoMongoAdapter(t *testing.T) {
	configMongo := config.DBInfoMongoConfigure{}
	mongoInfoDB := adapters.NewDBInfoMongoDB(configMongo)
	if mongoInfoDB == nil {
		t.Errorf("TestNewDBInfoMongoAdapter: expect(!nil) - got(%v)\n", mongoInfoDB)
	}
	log.Printf("mongoInfoDB: %+v\n", mongoInfoDB)
}

func TestDBInfoMongoAdapterConnect(t *testing.T) {
	configMongo := config.DBInfoMongoConfigure{
		Host:       "localhost",
		Username:   "root",
		Password:   "secret",
		Database:   "informeai",
		Port:       "27017",
		Collection: "dbinfos",
	}
	mongoInfoDB := adapters.NewDBInfoMongoDB(configMongo)
	if mongoInfoDB == nil {
		t.Errorf("TestDBInfoMongoAdapterConnect: expect(!nil) - got(%v)\n", mongoInfoDB)
	}
	if err := mongoInfoDB.Connect(); err != nil {
		t.Errorf("TestDBInfoMongoAdapterConnect: expect(nil) - got(%s)\n", err.Error())
	}
	if err := mongoInfoDB.Disconnect(); err != nil {
		t.Errorf("TestDBInfoMongoAdapterConnect: expect(nil) - got(%s)\n", err.Error())
	}
}

func TestDBInfoMongoAdapterSave(t *testing.T) {
	configMongo := config.DBInfoMongoConfigure{
		Host:       "localhost",
		Username:   "root",
		Password:   "secret",
		Database:   "informeai",
		Port:       "27017",
		Collection: "dbinfos",
	}
	mongoInfoDB := adapters.NewDBInfoMongoDB(configMongo)
	if mongoInfoDB == nil {
		t.Errorf("TestDBInfoMongoAdapterSave: expect(!nil) - got(%v)\n", mongoInfoDB)
	}
	if err := mongoInfoDB.Connect(); err != nil {
		t.Errorf("TestDBInfoMongoAdapterSave: expect(nil) - got(%s)\n", err.Error())
	}
	stct := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "informeai",
		Age:  23,
	}
	data := make(map[string]any)
	data["content"] = "tester"
	data["message"] = "dbinfo message"
	data["datetime"] = time.Now().String()
	data["user"] = stct
	if err := mongoInfoDB.Save("dbinfo.test", data); err != nil {
		t.Errorf("TestDBInfoMongoAdapterSave: expect(nil) - got(%s)\n", err.Error())
	}
	if err := mongoInfoDB.Disconnect(); err != nil {
		t.Errorf("TestDBInfoMongoAdapterSave: expect(nil) - got(%s)\n", err.Error())
	}
}
