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
			t.Error(err)
		}

		value, exists := store[tc.key]
		if !exists {
			t.Errorf("put failed: expected %q to be in the store", tc.key)
		}

		if value != tc.value {
			t.Errorf("value mismatch: expected %q, but got %q", tc.value, value)
		}
	}

	clear(store)
}
