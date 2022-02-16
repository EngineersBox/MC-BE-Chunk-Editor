package db

type LevelDB struct {
	Path       string
	Connection *interface{}
}

func Connect(worldPath string) LevelDB {
	return LevelDB{}
}

func (ldb *LevelDB) Get(key string) ([]byte, error) {
	// TODO: Implement this
	return []byte{}, nil
}

func (ldb *LevelDB) Store(key string, data []byte) error {
	// TODO: Implement this
	return nil
}
