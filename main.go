package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strconv"

	db "github.com/EngineersBox/MC-BE-Chunk-Editor/src/db"
	keys "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/keys"
	hexdump "github.com/EngineersBox/hexdump-format"
)

const (
	// testPath = `C:\Users\DEG3NERAT3_\AppData\Local\Packages\Microsoft.MinecraftUWP_8wekyb3d8bbwe\LocalState\games\com.mojang\minecraftWorlds\j4ELYtZvAAA=`
	testPath = "/Users/jackkilrain/Desktop/pob"

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

func writeToFile(file string, data []byte) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	bytesWritten, err := f.Write(data)
	if bytesWritten != len(data) {
		return fmt.Errorf(
			"did not write all data to file, only wrote %d/%d bytes",
			bytesWritten,
			len(data),
		)
	}
	return err
}

func createKey(x int, y int, z int, worldType string, tagType string, chunkCoords bool) keys.HexKey {
	var keyParameters keys.LDBKeyParameters = keys.LDBKeyParameters{
		Coords: keys.Coordinates{
			X: &x,
			Y: &y,
			Z: &z,
		},
		Attrs: keys.Attributes{
			WorldType:   &worldType,
			TagType:     &tagType,
			ChunkCoords: &chunkCoords,
		},
	}
	return keyParameters.CalculateHexKey()
}

func main() {
	world, err := db.LoadWorld(testPath)
	if err != nil {
		panic(err)
	}

	hexKey := createKey(110, 2, -22, "Overworld", "SubChunkPrefix", true)
	key, err := stringKeyToBytes(hexKey.ToString())
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("LevelDB Key: %x", key))

	data, err := world.Transactor.Get(key)
	if err != nil {
		panic(err)
	}
	writeToFile("chunkdata.bin", data)
	writeToFile("chunkdata.hex", []byte(hexdump.CreateHexdumpText(data)))
}
