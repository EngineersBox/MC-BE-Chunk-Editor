package db

import (
	"errors"
	"os"

	leveldb "github.com/syndtr/goleveldb/leveldb"
)

type World struct {
	path string
	Transactor
}

func LoadWorld(path string) (World, error) {
	world := World{
		path: path,
		Transactor: Transactor{
			conn: nil,
		},
	}
	worldLevelDB := path + "/db"

	info, err := os.Stat(worldLevelDB)
	if os.IsNotExist(err) {
		return world, errors.New("Could not open world at: " + worldLevelDB)
	}
	if err != nil {
		return world, err
	}
	if !info.IsDir() {
		return world, errors.New("Path does not point to a directory")
	}
	world.Transactor.conn, err = leveldb.OpenFile(worldLevelDB, nil)
	if err != nil {
		if world.Transactor.conn != nil {
			_ = world.Transactor.conn.Close()
		}
		return world, err
	}
	return world, nil
}
