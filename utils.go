package main

func isInSlice(needle string, haystack []string) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}

	return false
}
