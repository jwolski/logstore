package main

import "strconv"

// Looks up key in map and convert its string value to an int. If key does not
// exist, then default value is returned.
func intOrDefault(m map[string]string, key string, def int32) int32 {
	value, ok := m[key]
	if !ok {
		return def
	}

	if ret, err := strconv.Atoi(value); err != nil {
		return int32(ret)
	}

	return def
}
