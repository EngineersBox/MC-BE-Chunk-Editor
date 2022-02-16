package chunk

const (
	ChunkVersion        ChunkDataType = 0x2c
	Data2D              ChunkDataType = 0x2d
	Data2DLegacy        ChunkDataType = 0x2e
	SubChunkPrefix      ChunkDataType = 0x2f
	LegacyTerrain       ChunkDataType = 0x30
	BlockEntity         ChunkDataType = 0x31
	Entity              ChunkDataType = 0x32
	PendingTicks        ChunkDataType = 0x33
	BlockExtraData      ChunkDataType = 0x35
	BiomeState          ChunkDataType = 0x35
	UNUSED              ChunkDataType = 0x37
	BorderBlocks        ChunkDataType = 0x38
	HardCodedSpawnAreas ChunkDataType = 0x39
	RandomTicks         ChunkDataType = 0x3a
	Checksums           ChunkDataType = 0x3b
	ChunkVersionV116100 ChunkDataType = 0x76
)

type ChunkDataType byte

type ChunkData struct {
	Sort         ChunkDataType
	Header       *[]byte
	NbtBody      *[]byte
	TypedNbtBody *[]interface{}
}

func NewChunkData() ChunkData {
	return ChunkData{
		Sort:         0x00,
		Header:       nil,
		NbtBody:      nil,
		TypedNbtBody: nil,
	}
}
