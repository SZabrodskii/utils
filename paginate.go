package main

func Paginate[T any](items []T, page int, pageSize int) []T {
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(items) {
		return []T{}
	}

	if end > len(items) {
		end = len(items)
	}

	return items[start:end]
}
