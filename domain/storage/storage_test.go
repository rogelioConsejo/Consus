package storage

import (
	"errors"
	"fmt"
	"github.com/rogelioConsejo/Consus/domain/storable"
	"testing"
)

func TestStoreAndUseItem(t *testing.T) {
	var label storable.Label
	label = make(storable.Label)

	const testKey = "testKey"
	const testValue = "testValue"

	label[testKey] = testValue

	item := storable.Storable(label)

	//Store
	storage := Storage()
	err, id := storage.Store(item)

	var retrievedItem storable.Type
	if err == nil {
		//Use
		err, retrievedItem = storage.Use(id)
	} else {
		err = errors.New(fmt.Sprintf("error trying Store(): \n%s", err.Error()))
	}


	if err == nil && retrievedItem != nil{
		//Read label of retrieved
		if retrievedItem.Label()[testKey] == testValue {
			println(retrievedItem.Label()[testKey])
			println("ok")
		} else {
			err = errors.New("value retrieved incorrectly")
		}
	} else if err != nil {
		err = errors.New(fmt.Sprintf("error trying Use(): \n%s", err.Error()))
	}

	if err == nil {
		//Use previously Used (should fail and return error)
		err, notRetrievedItem := storage.Use(id)
		if err == nil || notRetrievedItem != nil{
			err = errors.New("error: must not be able to Use() a storable that has already been Use()'d")
		}
	}

	if err != nil {
		t.Error(err.Error())
	}
}
