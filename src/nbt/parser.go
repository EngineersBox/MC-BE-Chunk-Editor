package nbt

import (
	"bytes"
	"errors"
)

const (
	State_Finished ParserState = iota
	State_Running
	State_Exception
	State_Idle
)

type ParserState uint8

type NBTParser struct {
	State        ParserState
	StateMessage *string
	buf          bytes.Buffer
	tags         []NBTTag
}

func NewParser() NBTParser {
	return NBTParser{
		State:        State_Idle,
		StateMessage: nil,
		tags:         nil,
	}
}

func (p *NBTParser) FromByteBuffer(buf bytes.Buffer) error {
	if buf.Len() <= 0 || buf.Cap() <= 0 {
		return errors.New("Cannot parse empty buffer")
	}
	p.buf = buf
	return nil
}

func (p *NBTParser) Parse() error {
	return nil
}

func (p *NBTParser) GetParsedTags() []NBTTag {
	return p.tags
}
