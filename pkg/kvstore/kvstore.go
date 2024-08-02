package kvstore

type (
	KVStore interface {
		Get(key string) string
		Set(key, value string)
		Unset(key string)
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

func (m *kvStore) Set(key, value string) {
	m.data[key] = value
}

func (m *kvStore) Unset(key string) {
	delete(m.data, key)
}
