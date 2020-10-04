package storage

import (
	"errors"
	"fmt"
	"github.com/rogelioConsejo/Consus/storage/persistence"
	"github.com/rogelioConsejo/Consus/storage/storable"
	"testing"
)

func TestStoreAndCheckOutItem(t *testing.T) {
	var label storable.Label
	label = make(storable.Label)

	const testKey = "testKey"
	const testValue = "testValue"

	label[testKey] = testValue

	item := storable.Storable(label)

	//Store
	storage := Storage(persistence.Mock())
	err, id := storage.Store(item)

	var retrievedItem storable.Type
	if err == nil {
		fmt.Printf("saved: {%s} \n", id)
		//CheckOut
		err, retrievedItem = storage.CheckOut(id)
	} else {
		err = errors.New(fmt.Sprintf("error trying Store(): \n%s", err.Error()))
	}


	if err == nil && retrievedItem != nil{
		fmt.Printf("{%s} retrieved\n", id)
		//Read label of retrieved
		if retrievedItem.Label()[testKey] == testValue {
			println(retrievedItem.Label()[testKey])
			println("ok")
		} else {
			err = errors.New("value retrieved incorrectly")
		}
	} else if err != nil {
		err = errors.New(fmt.Sprintf("error trying CheckOut(): \n%s", err.Error()))
	}

	if err == nil {
		//CheckOut previously Used (should fail and return error)
		err, notRetrievedItem := storage.CheckOut(id)
		if err == nil || notRetrievedItem != nil{
			err = errors.New("error: must not be able to CheckOut() a storable that has already been CheckOut()'d")
		} else {
			fmt.Printf("{%s} already retrieved: %s\n", id, err.Error())
		}
	}

	if err != nil {
		t.Error(err.Error())
	}
}
