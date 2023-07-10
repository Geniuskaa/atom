package main

import (
	carSrvGrpcV1 "atom/internal/service/generated/car/v1"
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile("certs/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

var (
	ADDR string
	PORT string
)

func main() {
	flag.StringVar(&ADDR, "a", "", "")
	flag.StringVar(&PORT, "p", "", "")
	flag.Parse()

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGHUP, syscall.SIGINT,
		syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	tlsCreds, err := loadTLSCredentials()
	if err != nil {
		panic(err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(tlsCreds),
	}

	conn, err := grpc.DialContext(mainCtx, net.JoinHostPort(ADDR, PORT), opts...)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	client := carSrvGrpcV1.NewCarServiceClient(conn)

	sc := bufio.NewScanner(os.Stdin)

	stdInCh := make(chan string, 1)

	go func(sc *bufio.Scanner, ch chan string) {
		for sc.Scan() {
			ch <- sc.Text()
		}
	}(sc, stdInCh)

	var userMsg string

	for {
		fmt.Println("Выберите запрос (введите номер), который хотите послать:")
		fmt.Println("1. Зарегистрировать автомобиль")
		fmt.Println("2. Посмотреть информацию о машине")
		fmt.Println("3. Добавить информацию о машине")
		fmt.Println("4. Обновить информацию о машине")
		fmt.Println("5. Удалить информацию о машине")
		fmt.Println("6. Получить информацию о пробеге машины")
		fmt.Println("7. Посмотреть информацию по всем машинам")
		fmt.Println("8. Добавить пробег автомобиля")
		fmt.Println("9. Получить информацию о стоимости использования авто за 1 км")
		fmt.Println("10. Запустить симулятор передачи данных датчиков автомобиля на сервер")
		fmt.Println("99. Выйти")

		select {
		case <-mainCtx.Done():
			return
		case userMsg = <-stdInCh:
		}

		num, err := strconv.Atoi(userMsg)
		if err != nil {
			fmt.Println("Пожалуйста, введите положительное число!")
			time.Sleep(time.Second * 2)
			clearConsole()
			continue
		}

		ctxTimeOut, cancel := context.WithTimeout(mainCtx, time.Minute*10)

		switch num {
		case 1:
			clearConsole()
			regCar(ctxTimeOut, stdInCh, client)
			time.Sleep(time.Second * 5)
		case 2:
			clearConsole()
			getCarInfo(ctxTimeOut, stdInCh, client)
			time.Sleep(time.Second * 5)
		case 3:
			clearConsole()
			addCarInfo(ctxTimeOut, stdInCh, client)
			time.Sleep(time.Second * 5)
		case 4:
			clearConsole()
			updateCarInfo(ctxTimeOut, stdInCh, client)
			time.Sleep(time.Second * 5)
		case 5:
			clearConsole()
			deleteCarInfo(ctxTimeOut, stdInCh, client)
			time.Sleep(time.Second * 5)
		case 6:
			clearConsole()
			getCarMileage(ctxTimeOut, stdInCh, client)
			time.Sleep(time.Second * 5)
		case 7:
			clearConsole()
			getAllCarsInfo(ctxTimeOut, client)
			time.Sleep(time.Second * 5)
		case 8:
			clearConsole()
			addCarMileage(ctxTimeOut, stdInCh, client)
			time.Sleep(time.Second * 5)
		case 9:
			clearConsole()
			getCarUsageCost(ctxTimeOut, stdInCh, client)
			time.Sleep(time.Second * 5)
		case 10:
			clearConsole()
			simulateSensors(ctxTimeOut, client)
			time.Sleep(time.Second * 3)
		default:
			clearConsole()
			fmt.Println("Досвидания!")
			return
		}

		cancel()
	}

}

func clearConsole() {
	cmdName := "clear"
	cmd := exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// TODO: finish it
func simulateSensors(ctxTimeOut context.Context, client carSrvGrpcV1.CarServiceClient) {
	stream, err := client.SendCarMetrics(ctxTimeOut)
	if err != nil {
		fmt.Println("Произошла ошибка симуляции данных с датчиков")
		return
	}

	for {
		select {
		case <-ctxTimeOut.Done():
			stream.CloseSend()
			fmt.Println("Симуляция завершена")
			return
		default:
			time.Sleep(time.Millisecond * 1000)
		}

		sensor1 := rand.Int63n(250) - 40
		sensor2 := rand.Int63n(250) - 40
		sensor3 := rand.Int63n(671) - 273

		part1 := strconv.FormatInt(sensor1, 16)
		part2 := strconv.FormatInt(sensor2, 16)
		part3 := strconv.FormatInt(sensor3, 16)

		str := fmt.Sprintf("%s%s%s", part1, part2, part3)

		if len(str) < 16 {
			i := 16 - len(str)
			for ; i > 0; i-- {
				str += "0"
			}
		}

		fmt.Println(len(str))
		per := int32(1000)
		req := carSrvGrpcV1.CarMetrics{
			ID:     []byte("0CFEEE00"),
			Data:   []byte(str),
			Period: &per,
		}
		err := stream.Send(&req)

		if err != nil {
			stream.CloseSend()
			fmt.Println("Симуляция завершена")
			return
		}
	}

}

func regCar(ctxTimeOut context.Context, stdInCh <-chan string, client carSrvGrpcV1.CarServiceClient) {
	fmt.Println("Пожалуйста заполните требуемые поля:")
	fmt.Println("Пожалуйста вводите номер английскими буквами, иначе будет ошибка.")
	fmt.Print("Номер автомобиля с кодом региона (Например: K191AK116): ")
	plate := <-stdInCh

	fmt.Print("Модель автомобиля (Например: BMW x5 30d): ")
	carModel := <-stdInCh

	fmt.Print("Год выпуска автомобиля (Например: 2016): ")
	year := <-stdInCh

	fmt.Print("Тип автомобиля. Выбрать из диапазона (carsharing, taxi, delivery): ")
	carType := <-stdInCh

	var carTypeProto carSrvGrpcV1.CarType

	switch carType {
	case "carsharing":
		carTypeProto = carSrvGrpcV1.CarType_Carsharing
	case "taxi":
		carTypeProto = carSrvGrpcV1.CarType_Taxi
	case "delivery":
		carTypeProto = carSrvGrpcV1.CarType_Delivery
	default:
		carTypeProto = carSrvGrpcV1.CarType_Undefined
	}

	req := carSrvGrpcV1.CarInfo{
		Plate:             plate,
		Model:             carModel,
		YearOfManufacture: year,
		Type:              carTypeProto,
	}

	resp, err := client.RegisterCar(ctxTimeOut, &req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}
	fmt.Println("Вы успешно зарегистрировали автомобиль! Вашему авто присвоен ID = ", resp.Id)
	fmt.Println("Сохраните его, чтобы в дальнейшем вы смогли обновлять информацию об автомобиле.")
}

func getCarInfo(ctxTimeOut context.Context, stdInCh <-chan string, client carSrvGrpcV1.CarServiceClient) {
	fmt.Println("Пожалуйста заполните требуемые поля:")
	fmt.Println("Пожалуйста вводите номер английскими буквами, иначе будет ошибка.")
	fmt.Print("Номер автомобиля с кодом региона (Например: K191AK116): ")
	plate := <-stdInCh
	req := carSrvGrpcV1.CarPlate{Plate: plate}

	resp, err := client.GetCarInfo(ctxTimeOut, &req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}
	fmt.Println("Номер автомобиля: ", resp.GetPlate())
	fmt.Println("Модель автомобиля: ", resp.GetModel())
	fmt.Println("Год выпуска автомобиля: ", resp.GetYearOfManufacture())
	fmt.Println("Тип автомобиля: ", resp.GetType())
}

func addCarInfo(ctxTimeOut context.Context, stdInCh <-chan string, client carSrvGrpcV1.CarServiceClient) {
	fmt.Println("Пожалуйста заполните требуемые поля:")
	fmt.Print("Введите ID автомобиля, у которого вы хотите добавить данные: ")
	id := <-stdInCh

	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	fmt.Println("Пожалуйста вводите номер английскими буквами, иначе будет ошибка.")
	fmt.Print("Номер автомобиля с кодом региона (Например: K191AK116): ")
	plate := <-stdInCh

	fmt.Print("Модель автомобиля (Например: BMW x5 30d): ")
	carModel := <-stdInCh

	fmt.Print("Год выпуска автомобиля (Например: 2016): ")
	year := <-stdInCh

	fmt.Print("Тип автомобиля. Выбрать из диапазона (carsharing, taxi, delivery): ")
	carType := <-stdInCh

	var carTypeProto carSrvGrpcV1.CarType

	switch carType {
	case "carsharing":
		carTypeProto = carSrvGrpcV1.CarType_Carsharing
	case "taxi":
		carTypeProto = carSrvGrpcV1.CarType_Taxi
	case "delivery":
		carTypeProto = carSrvGrpcV1.CarType_Delivery
	default:
		carTypeProto = carSrvGrpcV1.CarType_Undefined
	}

	req := carSrvGrpcV1.UploadCarInfo{
		Id:                int32(intID),
		Plate:             plate,
		Model:             carModel,
		YearOfManufacture: year,
		Type:              carTypeProto,
	}

	resp, err := client.AddCarInfo(ctxTimeOut, &req)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	if resp.Ok {
		fmt.Println("Вы успешно добавили информацию об автомобиле!")
		return
	}
	fmt.Println("Что-то пошло не так, нам не удалось добавить информацию об автомобиле!")
}

func updateCarInfo(ctxTimeOut context.Context, stdInCh <-chan string, client carSrvGrpcV1.CarServiceClient) {
	fmt.Println("Пожалуйста заполните требуемые поля:")
	fmt.Print("Введите ID автомобиля, у которого вы хотите обновить данные:")

	var id, plate, carModel, year, carType string

	select {
	case id = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	fmt.Println("Пожалуйста вводите номер английскими буквами, иначе будет ошибка.")
	fmt.Print("Номер автомобиля с кодом региона (Например: K191AK116): ")

	select {
	case plate = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	fmt.Print("Модель автомобиля (Например: BMW x5 30d): ")

	select {
	case carModel = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	fmt.Print("Год выпуска автомобиля (Например: 2016): ")

	select {
	case year = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	fmt.Print("Тип автомобиля. Выбрать из диапазона (carsharing, taxi, delivery): ")

	select {
	case carType = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	var carTypeProto carSrvGrpcV1.CarType

	switch carType {
	case "carsharing":
		carTypeProto = carSrvGrpcV1.CarType_Carsharing
	case "taxi":
		carTypeProto = carSrvGrpcV1.CarType_Taxi
	case "delivery":
		carTypeProto = carSrvGrpcV1.CarType_Delivery
	default:
		carTypeProto = carSrvGrpcV1.CarType_Undefined
	}

	req := carSrvGrpcV1.UploadCarInfo{
		Id:                int32(intID),
		Plate:             plate,
		Model:             carModel,
		YearOfManufacture: year,
		Type:              carTypeProto,
	}

	resp, err := client.UpdateCarInfo(ctxTimeOut, &req)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	if resp.Ok {
		fmt.Println("Вы успешно обновили информацию об автомобиле!")
		return
	}
	fmt.Println("Что-то пошло не так, нам не удалось обновить информацию об автомобиле!")
}

func deleteCarInfo(ctxTimeOut context.Context, stdInCh <-chan string, client carSrvGrpcV1.CarServiceClient) {
	fmt.Println("Пожалуйста заполните требуемые поля:")
	fmt.Print("Введите ID автомобиля, который вы хотите удалить:")
	id := <-stdInCh

	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	req := carSrvGrpcV1.CarID{Id: int32(intID)}

	ok, err := client.DeleteCarInfo(ctxTimeOut, &req)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	if ok.Ok {
		fmt.Println("Вы успешно удалили информацию об автомобиле!")
		return
	}
	fmt.Println("Что-то пошло не так, нам не удалось удалить информацию об автомобиле!")
}

func getAllCarsInfo(ctxTimeOut context.Context, client carSrvGrpcV1.CarServiceClient) {
	req := empty.Empty{}

	resp, err := client.GetAllCarsInfo(ctxTimeOut, &req)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	for i, val := range resp.Cars {
		fmt.Printf("%d. Car plate: %s, model: %s, year of manudacture: %s, type: %s\n", i+1, val.Plate, val.Model,
			val.YearOfManufacture, val.Type.String())
	}

}

func getCarMileage(ctxTimeOut context.Context, stdInCh <-chan string, client carSrvGrpcV1.CarServiceClient) {
	fmt.Println("Пожалуйста заполните требуемые поля:")
	fmt.Println("Пожалуйста вводите номер английскими буквами, иначе будет ошибка.")
	fmt.Print("Номер автомобиля с кодом региона (Например: K191AK116): ")
	var plate string

	select {
	case plate = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	req := carSrvGrpcV1.CarPlate{Plate: plate}

	mileage, err := client.GetCarMileage(ctxTimeOut, &req)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	fmt.Printf("Пробег автомобиля с номером %s составляет %d км\n", strings.ToUpper(plate), mileage.Mileage)
}

func addCarMileage(ctxTimeOut context.Context, stdInCh <-chan string, client carSrvGrpcV1.CarServiceClient) {
	fmt.Println("Пожалуйста заполните требуемые поля:")
	fmt.Print("Введите ID автомобиля, пробег которого вы хотите увеличить: ")
	var id, amount string

	select {
	case id = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	fmt.Print("Введите пробег в км которое проехало авто (Указывать только целочисленное значение): ")

	select {
	case amount = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	if intAmount < 0 {
		fmt.Println("Пробег должен быть положительным!")
		return
	}

	req := carSrvGrpcV1.CarMileageUpload{
		Id:      int32(intID),
		Mileage: int32(intAmount),
	}

	resp, err := client.AddCarMileage(ctxTimeOut, &req)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	if resp.Ok {
		fmt.Println("Вы успешно изменили информацию о пробеге автомобиля!")
		return
	}
	fmt.Println("Что-то пошло не так, нам не удалось увеличить пробег автомобиле!")
}

func getCarUsageCost(ctxTimeOut context.Context, stdInCh <-chan string, client carSrvGrpcV1.CarServiceClient) {
	fmt.Println("Пожалуйста заполните требуемые поля:")
	fmt.Println("Пожалуйста вводите номер английскими буквами, иначе будет ошибка.")
	fmt.Print("Номер автомобиля с кодом региона (Например: K191AK116): ")
	var plate, fuelCosts, carTax, autoMCosts, autoCCosts, insuranceCost string

	select {
	case plate = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	fmt.Print("Расходы на топливо: ")

	select {
	case fuelCosts = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	intFuel, err := strconv.Atoi(fuelCosts)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	fmt.Print("Расходы на налог: ")

	select {
	case carTax = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	intTax, err := strconv.Atoi(carTax)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	fmt.Print("Расходы на ТО: ")

	select {
	case autoMCosts = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	intM, err := strconv.Atoi(autoMCosts)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	fmt.Print("Расходы на запчасти: ")

	select {
	case autoCCosts = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	intC, err := strconv.Atoi(autoCCosts)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	fmt.Print("Расходы на страховку: ")

	select {
	case insuranceCost = <-stdInCh:
	case <-ctxTimeOut.Done():
		return
	}

	intInsurance, err := strconv.Atoi(insuranceCost)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	req := carSrvGrpcV1.CarPlateAndCosts{
		Plate:                plate,
		FuelCosts:            int32(intFuel),
		CarTax:               int32(intTax),
		AutoMaintenanceCosts: int32(intM),
		AutoConsumablesCost:  int32(intC),
		InsuranceCost:        int32(intInsurance),
	}

	resp, err := client.GetCarUsageCost(ctxTimeOut, &req)
	if err != nil {
		fmt.Println("Что-то пошло не так. Попробуйте снова!")
		return
	}

	fmt.Printf("Каждый км пути на авто с номером %s обходится в %d руб\n", strings.ToUpper(plate), resp.CostInRUB)
}
