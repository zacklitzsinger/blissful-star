package main

import "errors"

type (
	MockDB interface {
		Get(key string) string
		Create(key, value string) error
		Update(key, value string) error
		Delete(key string) error
	}

	mockDB struct {
		db map[string]string
	}
)

/* Connect to DB */
func Connect() MockDB {
	return &mockDB{db: make(map[string]string)}
}

/* Get value from DB */
func (m *mockDB) Get(key string) string {
	return m.db[key]
}

/* Create key-value pair in DB */
func (m *mockDB) Create(key, value string) error {
	if _, ok := m.db[key]; ok {
		return errors.New("key already exists")
	}
	m.db[key] = value
	return nil
}

/* Update value in DB */
func (m *mockDB) Update(key, value string) error {
	if _, ok := m.db[key]; !ok {
		return errors.New("key does not exist")
	}
	m.db[key] = value
	return nil
}

/* Delete key-value pair from DB */
func (m *mockDB) Delete(key string) error {
	if _, ok := m.db[key]; !ok {
		return errors.New("key does not exist")
	}
	delete(m.db, key)
	return nil
}
