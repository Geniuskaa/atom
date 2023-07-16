package saeJ1939

import (
	"fmt"
	"strconv"
)

// Сообщение таково плана будем получать от источника данных, например от автобус(грузовика).
// Формат и тип данных не финален, но очень близок к документациям найденным в интернете.
// Возможно поля структуры будут завернуты в []byte
type J1939_MSG_Packed struct {
	// In ms
	Period uint16
	// In hexadecimal format. Example: 0CF00400
	ID string
	// Always consists from 8 bytes in hexadecimal format. Example: E1 DF 2F 64 00 64 FF FA
	Data string
}

// Сообщение, которое мы "распаковали" на сервере, узнали группу к которой принадлежат данные.
// Далее на основе документации группы мы соотносим нужное кол-во байт поля Data к определенной характеристике,
// переводим их в 10-ую систему и получаем человекопонятные значения с датчиков, которыми мы можем оперировать
// как нам вздумается.
type J1939_MSG_Unpacked struct {
	Priority uint16
	// It (PGN) refers to few SPNs, which you can find at PGNtoSPNs.Storage
	PGNreference uint16
	// In ms
	Period uint16
	// Always consists from 8 bytes in hexadecimal format. Example: E1 DF 2F 64 00 64 FF FA
	Data string
}

// Each key(PGN) can store few SPN and if it so - you should remember, that first value in slice has zero offset
// and next values`es offsets are the sum of SPN DataLen previous ones.
// So you have to begin to call GetOffset() func at SPN from first value in slice till the last one.
type PGNtoSPNs struct {
	Storage map[uint16][]SPN
}

type SPN interface {
	GetOffset(prevOffset uint8) uint8
	GetDataLen() uint8
	GetDataOffset() any
	GetResolution() string
}

type SPN_int struct {
	ID uint16
	// In bytes. Example, data may take from 1 byte to 4 byte.
	DataLen uint8
	// Offset in its units from value that we get from J1939_MSG_Unpacked.Data after decimal conversion.
	Offset int16
	Range  struct {
		From int32
		To   int32
	}
	Resolution string
	Type       string
}

func (s SPN_int) GetOffset(prevOffset uint8) uint8 {
	return prevOffset + s.DataLen
}

func (s SPN_int) GetDataLen() uint8 {
	return s.DataLen
}

func (s SPN_int) GetDataOffset() any {
	return s.Offset
}

func (s SPN_int) GetResolution() string {
	return s.Resolution
}

type SPN_float struct {
	ID uint16
	// In bytes
	DataLen uint8
	Offset  float32
	Range   struct {
		From float32
		To   float32
	}
	Resolution string
	Type       string
}

func (s SPN_float) GetOffset(prevOffset uint8) uint8 {
	return prevOffset + s.DataLen
}

func (s SPN_float) GetDataLen() uint8 {
	return s.DataLen
}

func (s SPN_float) GetDataOffset() any {
	return s.Offset
}

func (s SPN_float) GetResolution() string {
	return s.Resolution
}

func (m *J1939_MSG_Packed) Unpack() (J1939_MSG_Unpacked, error) {
	unpackedMsg := J1939_MSG_Unpacked{
		Data:   m.Data,
		Period: m.Period,
	}

	// base 16 for hexadecimal
	intFromHex, err := strconv.ParseUint(m.ID, 16, 64)
	if err != nil {
		return unpackedMsg, fmt.Errorf("Unpack failed: %w", err)
	}

	binFromInt := strconv.FormatUint(intFromHex, 2)
	offsetL := len(binFromInt) - 8
	offsetR := len(binFromInt) - 8

	//sourceAddr := binFromInt[offsetL:]
	offsetL -= 8
	pduSpecific := binFromInt[offsetL:offsetR]
	offsetL -= 8
	offsetR -= 8
	pduFormat := binFromInt[offsetL:offsetR]
	rawPGNref := fmt.Sprintf("%s%s", pduFormat, pduSpecific)
	cleanPGN, err := strconv.ParseUint(rawPGNref, 2, 16)
	if err != nil {
		return unpackedMsg, fmt.Errorf("Unpack failed: %w", err)
	}
	unpackedMsg.PGNreference = uint16(cleanPGN)

	offsetL -= 1
	offsetR -= 8
	//dataPage := binFromInt[offsetL:offsetR]

	offsetL--
	offsetR--
	//extendedDataPage := binFromInt[offsetL:offsetR]

	offsetR -= 1
	rawPriority := binFromInt[:offsetR]
	cleanPriority, err := strconv.ParseUint(rawPriority, 10, 16)
	if err != nil {
		return unpackedMsg, fmt.Errorf("Unpack failed: %w", err)
	}
	unpackedMsg.Priority = uint16(cleanPriority)

	//Первые 3 бита

	return unpackedMsg, nil
}
