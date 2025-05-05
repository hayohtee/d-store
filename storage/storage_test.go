package storage

import "testing"

var tt = []struct {
	key   string
	value string
}{
	{
		key:   "demo-key1",
		value: "demo-value1",
	},
	{
		key:   "demo-key-2",
		value: "demo-value-2",
	},
	{
		key:   "demo-key-3",
		value: "demo-value-3",
	},
	{
		key:   "demo-key-4",
		value: "demo-value-4",
	},
}

func TestPut(t *testing.T) {
	for _, tc := range tt {
		if err := Put(tc.key, tc.value); err != nil {
			t.Fatal(err)
		}

		value, exist := store[tc.key]
		if !exist {
			t.Errorf("put failed: expected %q to be in the store", tc.key)
		}

		if value != tc.value {
			t.Errorf("value mismatch: expected %q, but got %q", tc.value, value)
		}
	}
	clear(store)
}

func TestGet(t *testing.T) {
	for _, tc := range tt {
		if err := Put(tc.key, tc.value); err != nil {
			t.Error(err)
		}

		value, err := Get(tc.key)
		if err != nil {
			t.Fatal(err)
		}

		if value != tc.value {
			t.Errorf("value mismatch: expected %q, but got %q", tc.value, value)
		}
	}
	clear(store)
}

func TestDelete(t *testing.T) {
	for _, tc := range tt {
		if err := Put(tc.key, tc.value); err != nil {
			t.Fatal(err)
		}
		
		if err := Delete(tc.key); err != nil {
			t.Fatal(err)
		}
		
		_, exists := store[tc.key]
		if exists {
			t.Errorf("delete failed: %q key should not exist", tc.key)
		}
	}
	clear(store)
}