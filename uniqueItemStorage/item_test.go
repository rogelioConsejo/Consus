package uniqueItemStorage

import "testing"

func TestStoreAndUseItem(t *testing.T) {
	err, newItem := Item("TestItem",nil)
	var id uint

	if err == nil {
		err, id = newItem.Store()
	} else {
		t.Error(err.Error())
	}

	var retrievedItem *item
	if err == nil {
		err, retrievedItem = RetrieveItem(id)
	} else {
		t.Error(err.Error())
	}

	if err == nil && retrievedItem != nil {
		err = retrievedItem.label.SetName("SameItemNewName")
	} else if err != nil {
		t.Error(err.Error())
	}

	if err == nil && retrievedItem != nil {
		err = retrievedItem.Use()
	}

	//Finally
	if err != nil {
		t.Error(err.Error())
	}
}
