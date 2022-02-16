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
	Name   *string
}

type NBTTagByte struct {
	NBTTag
	Value      *byte
	TypedValue *byte
}

func NewByte() NBTTagByte {
	return NBTTagByte{
		NBTTag: NBTTag{
			Format: Nbt_Byte,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
	}
}

type NBTTagShort struct {
	NBTTag
	Value      *[2]byte
	TypedValue *uint8
}

func NewShort() NBTTagShort {
	return NBTTagShort{
		NBTTag: NBTTag{
			Format: Nbt_Short,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
	}
}

type NBTTagInt struct {
	NBTTag
	Value      *[4]byte
	TypedValue *uint32
}

func NewInt() NBTTagInt {
	return NBTTagInt{
		NBTTag: NBTTag{
			Format: Nbt_Int,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
	}
}

type NBTTagLong struct {
	NBTTag
	Value      *[8]byte
	TypedValue *uint64
}

func NewLong() NBTTagLong {
	return NBTTagLong{
		NBTTag: NBTTag{
			Format: Nbt_Long,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
	}
}

type NBTTagFloat struct {
	NBTTag
	Value      *[4]byte
	TypedValue *float32
}

func NewFloat() NBTTagFloat {
	return NBTTagFloat{
		NBTTag: NBTTag{
			Format: Nbt_Float,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
	}
}

type NBTTagDouble struct {
	NBTTag
	Value      *[8]byte
	TypedValue *float64
}

func NewDouble() NBTTagDouble {
	return NBTTagDouble{
		NBTTag: NBTTag{
			Format: Nbt_Double,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
	}
}

type NBTTagByteArray struct {
	NBTTag
	Value  *[]byte
	Length *uint32
}

func NewByteArray() NBTTagByteArray {
	return NBTTagByteArray{
		NBTTag: NBTTag{
			Format: Nbt_ByteArray,
			Name:   nil,
		},
		Value:  nil,
		Length: nil,
	}
}

type NBTTagString struct {
	NBTTag
	Value      *[]byte
	TypedValue *string
	Length     *byte
}

func NewString() NBTTagString {
	return NBTTagString{
		NBTTag: NBTTag{
			Format: Nbt_String,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
		Length:     nil,
	}
}

type NBTTagList struct {
	NBTTag
	Value      *[]byte
	TypedValue *[]ListPayload
	Length     *uint32
}

func NewList() NBTTagList {
	return NBTTagList{
		NBTTag: NBTTag{
			Format: Nbt_List,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
		Length:     nil,
	}
}

type ListPayload struct {
	Id          *byte
	PayloadSize *uint32
	Value       *[]byte
	TypedValue  *interface{}
}

type NBTTagCompound struct {
	NBTTag
	Value      *[]byte
	TypedValue *[]NBTTag
	Length     *uint32
}

func NewCompound() NBTTagCompound {
	return NBTTagCompound{
		NBTTag: NBTTag{
			Format: Nbt_List,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
		Length:     nil,
	}
}

type NBTTagIntArray struct {
	NBTTag
	Length     *uint32
	Value      *[]byte
	TypedValue *[]uint32
}

func NewIntArray() NBTTagIntArray {
	return NBTTagIntArray{
		NBTTag: NBTTag{
			Format: Nbt_List,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
		Length:     nil,
	}
}

type NBTTagLongArray struct {
	NBTTag
	Length     *uint32
	Value      *[]byte
	TypedValue *[]uint64
}

func NewLongArray() NBTTagLongArray {
	return NBTTagLongArray{
		NBTTag: NBTTag{
			Format: Nbt_List,
			Name:   nil,
		},
		Value:      nil,
		TypedValue: nil,
		Length:     nil,
	}
}
