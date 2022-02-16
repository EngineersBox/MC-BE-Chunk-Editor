package main

import (
	"encoding/binary"
	"fmt"

	db "github.com/EngineersBox/MC-BE-Chunk-Editor/src/db"
)

const testPath string = `C:\Users\DEG3NERAT3_\AppData\Local\Packages\Microsoft.MinecraftUWP_8wekyb3d8bbwe\LocalState\games\com.mojang\minecraftWorlds\j4ELYtZvAAA=`

func main() {
	world, err := db.LoadWorld(testPath)
	if err != nil {
		panic(err)
	}
	key1 := make([]byte, 4)
	binary.BigEndian.PutUint32(key1, 0x6e000000)
	key2 := make([]byte, 4)
	binary.BigEndian.PutUint32(key2, 0xeaffffff)
	key3 := make([]byte, 2)
	binary.BigEndian.PutUint16(key3, 0x2f02)

	key := append(key1, key2...)
	key = append(key, key3...)
	println(fmt.Sprintf("%x", key))

	data, err := world.Transactor.Get(key)
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("%d", len(data)))
}
