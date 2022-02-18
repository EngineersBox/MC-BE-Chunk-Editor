package nbt

const (
	Nbt_End Format = iota
	Nbt_Byte
	Nbt_Short
	Nbt_Int
	Nbt_Long
	Nbt_Float
	Nbt_Double
	Nbt_ByteArray
	Nbt_String
	Nbt_List
	Nbt_Compound
	Nbt_IntArray
	Nbt_longArray
)

type Format byte

type NBTTag struct {
	Format Format
	Value  interface{}
}

type NBTTagByte struct {
	Name       *string
	Value      *byte
	TypedValue *byte
}

func NewByte() NBTTag {
	return NBTTag{
		Format: Nbt_Byte,
		Value: NBTTagByte{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
		},
	}
}

type NBTTagShort struct {
	Name       *string
	Value      *[2]byte
	TypedValue *uint16
}

func NewShort() NBTTag {
	return NBTTag{
		Format: Nbt_Short,
		Value: NBTTagShort{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
		},
	}
}

type NBTTagInt struct {
	Name       *string
	Value      *[4]byte
	TypedValue *uint32
}

func NewInt() NBTTag {
	return NBTTag{
		Format: Nbt_Int,
		Value: NBTTagInt{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
		},
	}
}

type NBTTagLong struct {
	Name       *string
	Value      *[8]byte
	TypedValue *uint64
}

func NewLong() NBTTag {
	return NBTTag{
		Format: Nbt_Long,
		Value: NBTTagLong{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
		},
	}
}

type NBTTagFloat struct {
	Name       *string
	Value      *[4]byte
	TypedValue *float32
}

func NewFloat() NBTTag {
	return NBTTag{
		Format: Nbt_Float,
		Value: NBTTagFloat{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
		},
	}
}

type NBTTagDouble struct {
	Name       *string
	Value      *[8]byte
	TypedValue *float64
}

func NewDouble() NBTTag {
	return NBTTag{
		Format: Nbt_Double,
		Value: NBTTagDouble{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
		},
	}
}

type NBTTagByteArray struct {
	Name   *string
	Value  []byte
	Length *uint32
}

func NewByteArray() NBTTag {
	return NBTTag{
		Format: Nbt_ByteArray,
		Value: NBTTagByteArray{
			Name:   nil,
			Value:  nil,
			Length: nil,
		},
	}
}

type NBTTagString struct {
	Name       *string
	Value      []byte
	TypedValue *string
	Length     *byte
}

func NewString() NBTTag {
	return NBTTag{
		Format: Nbt_String,
		Value: NBTTagString{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
			Length:     nil,
		},
	}
}

type NBTTagList struct {
	Name       *string
	Value      []byte
	TypedValue []ListPayload
	Length     *uint32
}

func NewList() NBTTag {
	return NBTTag{
		Format: Nbt_List,
		Value: NBTTagList{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
			Length:     nil,
		},
	}
}

type ListPayload struct {
	Id          *byte
	PayloadSize *uint32
	Value       []byte
	TypedValue  *interface{}
}

type NBTTagCompound struct {
	Value      []byte
	TypedValue []NBTTag
	Length     *uint32
}

func NewCompound() NBTTag {
	return NBTTag{
		Format: Nbt_Compound,
		Value: NBTTagCompound{
			Value:      nil,
			TypedValue: nil,
			Length:     nil,
		},
	}
}

type NBTTagIntArray struct {
	Name       *string
	Length     *uint32
	Value      []byte
	TypedValue []uint32
}

func NewIntArray() NBTTag {
	return NBTTag{
		Format: Nbt_IntArray,
		Value: NBTTagIntArray{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
			Length:     nil,
		},
	}
}

type NBTTagLongArray struct {
	Name       *string
	Length     *uint32
	Value      []byte
	TypedValue []uint64
}

func NewLongArray() NBTTag {
	return NBTTag{
		Format: Nbt_longArray,
		Value: NBTTagLongArray{
			Name:       nil,
			Value:      nil,
			TypedValue: nil,
			Length:     nil,
		},
	}
}
