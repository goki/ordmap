// Copyright (c) 2022, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ordmap

import (
	"testing"
)

func TestOrdMap(t *testing.T) {
	om := New[string, int]()
	om.Add("key0", 0)
	om.Add("key1", 1)
	om.Add("key2", 2)

	if v, ok := om.ValByKeyTry("key1"); !ok || v != 1 {
		t.Error("ValByKey")
	}

	if i, ok := om.IdxByKeyTry("key2"); !ok || i != 2 {
		t.Error("IdxByKey")
	}

	if om.KeyByIdx(0) != "key0" {
		t.Error("KeyByIdx")
	}

	if om.ValByIdx(1) != 1 {
		t.Error("ValByIdx")
	}

	if om.Len() != 3 {
		t.Error("Len")
	}

	om.DeleteIdx(1, 2)
	// for i, kv := range om.Order {
	// 	fmt.Printf("i: %d\tkv: %v\n", i, kv)
	// }
	if om.ValByIdx(1) != 2 {
		t.Error("DeleteIdx")
	}
	if i, ok := om.IdxByKeyTry("key2"); !ok || i != 1 {
		t.Error("Delete IdxByKey")
	}

	// fmt.Printf("%v\n", om.Map)
	om.InsertAtIdx(0, "new0", 3)
	// fmt.Printf("%v\n", om.Map)
	if om.ValByIdx(0) != 3 {
		t.Error("InsertAtIdx ValByIdx 0")
	}
	if om.ValByIdx(1) != 0 {
		t.Error("InsertAtIdx ValByIdx 1")
	}
	if i, ok := om.IdxByKeyTry("key2"); !ok || i != 2 {
		t.Errorf("InsertAtIdx IdxByKey: %d != 2", i)
	}

	// constr

	nm := Make([]KeyVal[string, int]{{"one", 1}, {"two", 2}, {"three", 3}})

	if nm.ValByIdx(2) != 3 {
		t.Error("Make ValByIdx 2")
	}
	if nm.ValByIdx(1) != 2 {
		t.Error("Make ValByIdx 1")
	}
	if val, ok := nm.ValByKeyTry("three"); !ok || val != 3 {
		t.Error("Make ValByKey 3")
	}

}
