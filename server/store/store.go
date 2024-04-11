package store

type Storer interface {
	Get(k string) (string, error)
	Set(k, v string) error
}

type MemStore struct {
	keyToValues map[string]string
}

func NewMemStore() Storer {
	return &MemStore{
		keyToValues: make(map[string]string),
	}
}

func (s *MemStore) Get(k string) (string, error) {
    v := s.keyToValues[k]
	return v, nil
}

func (s *MemStore) Set(k, v string) error {
    s.keyToValues[k] = v
	return nil
}
