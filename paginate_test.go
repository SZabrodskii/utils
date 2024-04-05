package main

import "testing"

func TestPaginate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	t.Run("should paginate", func(t *testing.T) {
		items := Paginate(arr, 3, 2)

		if len(items) != 2 {
			t.Errorf("expected to get 2 items, got %d", len(items))
		}

		if items[0] != 5 {
			t.Errorf("expected to get 5 as first item, got %d", items[0])
		}

		if items[1] != 6 {
			t.Errorf("expected to get 6 as second item, got %d", items[1])
		}
	})

	t.Run("should return empty array if page is out of range", func(t *testing.T) {
		items := Paginate(arr, 10, 2)

		if len(items) != 0 {
			t.Errorf("expected to get empty array, got %d", len(items))
		}
	})

	t.Run("should return last n items if page is out of range", func(t *testing.T) {
		items := Paginate(arr, 2, 5)

		if len(items) != 2 {
			t.Errorf("expected to get 2 items, got %d", len(items))
		}

		if items[0] != 6 {
			t.Errorf("expected to get 6 as first item, got %d", items[0])
		}

		if items[1] != 7 {
			t.Errorf("expected to get 7 as second item, got %d", items[1])
		}
	})
}
