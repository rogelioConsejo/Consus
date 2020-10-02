package storage

import "testing"

func TestStoreAndUseItem(t *testing.T) {
	err, newItem := Item("TestItem",nil)
	var id uint
	if err == nil {
		err, id = newItem.Store()
	}

	var retrievedItem *item
	if err == nil {
		err, retrievedItem = RetrieveItem(id)
	}

	if err == nil && retrievedItem != nil {
		err = retrievedItem.Use()
	}

	//Finally
	if err != nil {
		t.Error(err.Error())
	}
}
