package saeJ1939

import (
	"fmt"
	"strconv"
)

type Service struct {
	cache PGNtoSPNs
}

func NewService() *Service {

	srv := Service{}
	srv.cache = warmUpCache()

	return &srv
}

func warmUpCache() PGNtoSPNs {
	//TODO: сейчас функция захордкожена, но впоследствии предполагается, что данные PGN и SPN будут догружаться в БД
	// и сервер при старте будет подтягивать все значения себе в кэш, чтобы быстро распаковывать сообщения и записывать их.

	pgn := PGNtoSPNs{map[uint16][]SPN{
		65262: {
			SPN_int{
				ID:      110,
				DataLen: 1,
				Offset:  -40,
				Range: struct {
					From int32
					To   int32
				}{
					From: -40,
					To:   210,
				},
				Resolution: "deg C",
				Type:       "Measured",
			}, SPN_int{
				ID:      174,
				DataLen: 1,
				Offset:  -40,
				Range: struct {
					From int32
					To   int32
				}{
					From: -40,
					To:   210,
				},
				Resolution: "deg C",
				Type:       "Measured",
			}, SPN_int{
				ID:      175,
				DataLen: 2,
				Offset:  -273,
				Range: struct {
					From int32
					To   int32
				}{
					From: -273,
					To:   398,
				},
				Resolution: "deg C",
				Type:       "Measured",
			}},
	}}

	return pgn
}

func (s *Service) PullData(msg J1939_MSG_Unpacked) {
	if val, exst := s.cache.Storage[msg.PGNreference]; exst {
		offset := uint8(0)
		for _, spn := range val {
			num := msg.Data[offset : offset+spn.GetDataLen()]
			rawDataOffset := spn.GetDataOffset()

			switch rawDataOffset.(type) {
			case int16:
				measuredData, err := strconv.ParseInt(num, 16, 64)
				if err != nil {
					continue
				}
				dataOffset := rawDataOffset.(int16)
				fmt.Println(measuredData-int64(dataOffset), " ", spn.GetResolution())
			default:
				continue
			}

			offset = spn.GetOffset(offset)
		}
	}
}
