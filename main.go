package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	db "github.com/EngineersBox/MC-BE-Chunk-Editor/src/db"
)

const (
	testPath = `C:\Users\DEG3NERAT3_\AppData\Local\Packages\Microsoft.MinecraftUWP_8wekyb3d8bbwe\LocalState\games\com.mojang\minecraftWorlds\j4ELYtZvAAA=`

	overworldKeyLength = 20
	otherDimKeyLength  = 28
)

func stringKeyToBytes(key string) ([]byte, error) {
	keyLen := len(key)
	if keyLen != overworldKeyLength && keyLen != otherDimKeyLength {
		return nil, errors.New("Key must be 10 bytes for overworld or 14 bytes for another dimension")
	}
	worldAttrCount := keyLen / 8 // Each world attribute is 4 bytes, 8 hex chars
	var keyBytes []byte = []byte{}
	for i := 0; i < worldAttrCount; i++ {
		bytes := make([]byte, 4)
		intVal, err := strconv.ParseUint(key[i*8:(i+1)*8], 16, 32)
		if err != nil {
			return nil, err
		}
		binary.BigEndian.PutUint32(bytes, uint32(intVal))
		keyBytes = append(keyBytes, bytes...)
	}
	bytes := make([]byte, 2)
	intVal, err := strconv.ParseUint(key[keyLen-4:], 16, 16)
	if err != nil {
		return nil, err
	}
	binary.BigEndian.PutUint16(bytes, uint16(intVal))
	keyBytes = append(keyBytes, bytes...)
	return keyBytes, nil
}

func main() {
	world, err := db.LoadWorld(testPath)
	if err != nil {
		panic(err)
	}

	key, err := stringKeyToBytes("6e000000eaffffff2f02")
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("%x", key))

	data, err := world.Transactor.Get(key)
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("%x", data[4800:]))
}
