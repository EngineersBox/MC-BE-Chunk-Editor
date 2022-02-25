package db

import leveldb "github.com/engineersbox/goleveldb/leveldb"

type Transactor struct {
	conn *leveldb.DB
}

func (transactor *Transactor) Get(key []byte) ([]byte, error) {
	raw, err := transactor.conn.Get(key, nil)
	if err != nil {
		return nil, err
	}
	value := make([]byte, len(raw))
	copy(value, raw)
	return value, nil
}

func (transactor *Transactor) Store(key []byte, data []byte) error {
	return transactor.conn.Put(key, data, nil)
}

func (transactor *Transactor) Delete(key []byte) error {
	return transactor.conn.Delete(key, nil)
}
