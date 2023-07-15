package v1

import (
	"atom/internal/database/repository"
	"atom/internal/model"
	"atom/internal/service/fleetManagmentSystem/saeJ1939"
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"strings"
	"time"
)

type Service struct {
	logger   *zap.SugaredLogger
	db       repository.Car
	j1939Srv *saeJ1939.Service
	UnimplementedCarServiceServer
}

func NewService(logger *zap.SugaredLogger, repo repository.Car, j1939Srv *saeJ1939.Service) *Service {
	srv := Service{
		logger:                        logger,
		db:                            repo,
		j1939Srv:                      j1939Srv,
		UnimplementedCarServiceServer: UnimplementedCarServiceServer{},
	}

	return &srv
}

//
///*
//С - cost
//Стоимость эксплуатации =
//С(топлива)+С(налог)+С(ТО)+С(з/ч + шины)+С(страхование) / Пробег или кол-во моточасов.
//*/
//
//// Стоимость эксплуатации 1 км
//func (c *Service) CarUsagePrice(plate string) (int, error) {
//
//	//Стоимость в руб
//	gasPrice := 12.4
//	annualTaxPrice := 12_000
//	carMaintenanceCosts := 20_000
//	carSpareParts := 11_000
//	wheelTireCost := 12_000
//	carInsurance := 12_000
//
//	//TODO: получать пробег автомобиля по его номерному знаку
//
//	return 0, nil
//}

func srvCarTypeToGrpc(t string) CarType {

	switch t {
	case strings.ToUpper(CarType_Carsharing.String()):
		return CarType_Carsharing
	case strings.ToUpper(CarType_Taxi.String()):
		return CarType_Taxi
	case strings.ToUpper(CarType_Delivery.String()):
		return CarType_Delivery
	default:
		return CarType_Undefined
	}
}

func (s *Service) RegisterCar(ctx context.Context, ci *CarInfo) (*CarRegistered, error) {

	car := model.Car{
		Plate:             strings.ToUpper(ci.GetPlate()),
		Model:             strings.ToLower(ci.GetModel()),
		YearOfManufacture: ci.GetYearOfManufacture(),
	}

	car.SetCarType(ci.GetType().String())

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	id, err := s.db.RegisterCar(ctxTimeout, &car)

	if err != nil {
		s.logger.Errorf("RegisterCar failed: %w", err)

		return nil, err
	}

	resp := CarRegistered{}
	resp.Id = int32(id)
	return &resp, nil
}

func (s *Service) AddCarInfo(ctx context.Context, ci *UploadCarInfo) (*IsOK, error) {
	car := model.Car{
		Plate:             strings.ToUpper(ci.GetPlate()),
		Model:             strings.ToLower(ci.GetModel()),
		YearOfManufacture: ci.GetYearOfManufacture(),
	}
	car.SetCarType(ci.GetType().String())

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := s.db.UpdateCarInfo(ctxTimeout, &car, int(ci.GetId()))
	resp := IsOK{}
	if err != nil {
		s.logger.Errorf("AddCarInfo failed: %w", err)
		resp.Ok = false

		return &resp, err
	}

	resp.Ok = true

	return &resp, nil
}

func (s *Service) UpdateCarInfo(ctx context.Context, ci *UploadCarInfo) (*IsOK, error) {
	car := model.Car{
		Plate:             strings.ToUpper(ci.GetPlate()),
		Model:             strings.ToLower(ci.GetModel()),
		YearOfManufacture: ci.GetYearOfManufacture(),
	}
	car.SetCarType(ci.GetType().String())

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := s.db.UpdateCarInfo(ctxTimeout, &car, int(ci.GetId()))
	resp := IsOK{}
	if err != nil {
		s.logger.Errorf("UpdateCarInfo failed: %w", err)
		resp.Ok = false

		return &resp, err
	}

	resp.Ok = true

	return &resp, nil
}

func (s *Service) DeleteCarInfo(ctx context.Context, id *CarID) (*IsOK, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := s.db.DeleteCarInfo(ctxTimeout, int(id.GetId()))
	resp := IsOK{}
	if err != nil {
		s.logger.Errorf("DeleteCarInfo failed: %w", err)
		resp.Ok = false
		return &resp, err
	}

	resp.Ok = true
	return &resp, nil
}

func (s *Service) GetCarInfo(ctx context.Context, p *CarPlate) (*CarInfo, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	car, err := s.db.GetCarInfo(ctxTimeout, p.GetPlate())
	if err != nil {
		s.logger.Errorf("GetCarInfo failed: %w", err)

		return nil, err
	}

	resp := CarInfo{}
	resp.Plate = car.Plate
	resp.Model = car.Model
	resp.Type = srvCarTypeToGrpc(car.Type.String())
	resp.YearOfManufacture = car.YearOfManufacture

	return &resp, nil
}

func (s *Service) GetAllCarsInfo(ctx context.Context, e *empty.Empty) (*AllCarsInfo, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	cars, err := s.db.GetAllCarsInfo(ctxTimeout)
	if err != nil {
		s.logger.Errorf("GetAllCarsInfo failed: %w", err)
		fmt.Println("Db error!")
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	carDTOs := make([]*CarInfo, 0, len(cars))
	for _, val := range cars {
		car := CarInfo{
			Plate:             val.Plate,
			Model:             val.Model,
			YearOfManufacture: val.YearOfManufacture,
			Type:              srvCarTypeToGrpc(val.Type.String()),
		}

		carDTOs = append(carDTOs, &car)
	}

	resp := AllCarsInfo{Cars: carDTOs}

	fmt.Println("Returned cars!")
	return &resp, status.Error(codes.Unavailable, "")
}

func (s *Service) GetCarMileage(ctx context.Context, p *CarPlate) (*CarMileage, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	mileage, err := s.db.GetCarMileage(ctxTimeout, strings.ToUpper(p.GetPlate()))
	if err != nil {
		s.logger.Errorf("GetCarMileage failed: %w", err)

		return nil, err
	}

	resp := CarMileage{
		Mileage: int32(mileage),
	}

	return &resp, nil
}

func (s *Service) AddCarMileage(ctx context.Context, cm *CarMileageUpload) (*IsOK, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := s.db.AddCarMileage(ctxTimeout, int(cm.GetId()), int(cm.Mileage))
	resp := IsOK{}
	if err != nil {
		s.logger.Errorf("AddCarMileage failed: %w", err)
		resp.Ok = false

		return nil, err
	}

	resp.Ok = true

	return &resp, nil

}

// С - cost
// Стоимость эксплуатации =
// С(топлива)+С(налог)+С(ТО)+С(з/ч + шины)+С(страхование) / Пробег или кол-во моточасов.
func (s *Service) GetCarUsageCost(ctx context.Context, cpc *CarPlateAndCosts) (*CarUsageCost, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	mileage, err := s.db.GetCarMileage(ctxTimeout, strings.ToUpper(cpc.GetPlate()))
	if err != nil {
		s.logger.Errorf("GetCarUsageCost failed: %w", err)

		return nil, status.Errorf(codes.NotFound, "Info for car with such plate not found!")
	}

	sum := (cpc.FuelCosts + cpc.CarTax + cpc.AutoMaintenanceCosts + cpc.AutoConsumablesCost + cpc.InsuranceCost) / int32(mileage)

	resp := CarUsageCost{CostInRUB: sum}

	return &resp, nil
}

func (s *Service) SendCarMetrics(stream CarService_SendCarMetricsServer) error {
	for {
		m, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err := stream.SendAndClose(nil)
				if err != nil {
					s.logger.Errorf("SendCarMetrics failed during SendAndClose: %w", err)
				}
				return status.Errorf(codes.OK, "")
			}

			err := stream.SendAndClose(nil)
			if err != nil {
				s.logger.Errorf("SendCarMetrics failed during SendAndClose: %w", err)
			}
			return status.Errorf(codes.Unavailable, "some problem with sent data")
		}

		msg := saeJ1939.J1939_MSG_Packed{}

		if *m.Period < 0 || *m.Period > 32767 {
			continue
		}

		msg.Period = uint16(m.GetPeriod())

		if len(m.ID) != 8 {
			continue
		}

		msg.ID = string(m.ID)

		if len(m.Data) != 16 {
			continue
		}

		msg.Data = string(m.Data)

		unpackedMsg, err := msg.Unpack()
		if err != nil {
			continue
		}

		//Now it just println data from 'sensors'
		s.j1939Srv.PullData(unpackedMsg)

	}
}
