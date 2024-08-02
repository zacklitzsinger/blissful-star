package kvstore

import "errors"

type (
	KVStore interface {
		/* Get value from DB */
		Get(key string) string
		/* Set value */
		Set(key, value string) error
		/* Delete key-value pair */
		Delete(key string) error
	}

	kvStore struct {
		data map[string]string
	}
)

func CreateKV() KVStore {
	return &kvStore{data: make(map[string]string)}
}

func (m *kvStore) Get(key string) string {
	return m.data[key]
}

func (m *kvStore) Set(key, value string) error {
	if _, ok := m.data[key]; ok {
		return errors.New("key already exists")
	}
	m.data[key] = value
	return nil
}

func (m *kvStore) Delete(key string) error {
	if _, ok := m.data[key]; !ok {
		return errors.New("key does not exist")
	}
	delete(m.data, key)
	return nil
}
