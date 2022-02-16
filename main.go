package main

import (
	"fmt"

	chunk "github.com/EngineersBox/MC-BE-Chunk-Editor/src/chunk"
	nbt "github.com/EngineersBox/MC-BE-Chunk-Editor/src/nbt"
)

func main() {
	var test = chunk.NewChunkData()
	test.TypedNbtBody = &[]interface{}{
		nbt.NewByte(),
	}
	fmt.Println(*test.TypedNbtBody)
}
