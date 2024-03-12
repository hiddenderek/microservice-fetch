package database

// simulates a database

var dataStore = make(map[string]interface{})

func PostData[T any](key string, data T) (ret T) {
	dataStore[key] = data
	return any(dataStore[key]).(T)
}

func GetData[T any](key string) (T, bool) {
	ret, ok := (dataStore[key]).(T)
	return ret, ok
}
