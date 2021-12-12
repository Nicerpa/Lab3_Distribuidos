/*
  TO DO
  - Programar distincion entre slave y master (2 minutos)
  - Programar timer de 2 minutos
*/

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	fp "github.com/Lab3/FulcrumProto"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	fp.UnimplementedFulcrumServerServer
}

const (
	index = 0
	port  = ":50001" // Poner IP completa para el deploy
)

var fulcrum_addresses = []string{"localhost:50002", "localhost:50003"}

var clockMap = make(map[string][]int32)

// -----------------------------Funciones Utiles --------------------------------------------------

// Busca el numero almacenado de rebeldes para un planeta
func getNumbers(planet string, city string) int32 {
	input, err := ioutil.ReadFile(planet)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	var num int

	for _, line := range lines {
		if strings.Contains(line, city) {
			// planet city [num]
			num, _ = strconv.Atoi(strings.Split(line, " ")[2])

		}
	}

	return int32(num)
}

// Cambia las lineas de un archivo por los strings almacenados en line []string
func updateFileContent(lines []string, filename string) {
	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

// Lee las lineas de un archivo y las almacena en un arreglo de strings
func readFileContent(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")

	return lines
}

// Agrega el log de una accion, si es que no existe el archivo log entonces lo crea
func logAccion(command string, planet string, city string, newValue string) {
	f, err := os.OpenFile(strings.Join([]string{planet, ".log"}, ""), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s %s %s %s\n", command, planet, city, newValue)); err != nil {
		log.Println(err)
	}
}

// Añade una ciudad al archivo de un planeta
func addCity(planetName string, cityName string, rebels int32) {
	f, err := os.OpenFile(planetName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s %s %d\n", planetName, cityName, rebels)); err != nil {
		log.Println(err)
	}

	logAccion("AddCity", planetName, cityName, strconv.Itoa(int(rebels)))
}

// Borra una ciudad del archivo de un planeta
func deleteCity(planetName string, cityName string) {
	input, err := ioutil.ReadFile(planetName)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	var lineToDelete int

	for i, line := range lines {
		if strings.Contains(line, cityName) {
			lineToDelete = i
		}
	}
	copy(lines[lineToDelete:], lines[lineToDelete+1:]) // Shift a[i+1:] left one index.
	lines[len(lines)-1] = ""                           // Erase last element (write zero value).
	lines = lines[:len(lines)-1]                       // Truncate slice.

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(planetName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	logAccion("DeleteCity", planetName, cityName, "")
}

// Modifica el nombre de una ciudad en el archivo de un planeta
func updateCityName(planet string, city string, newName string) {
	input, err := ioutil.ReadFile(planet)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, city) {
			l := strings.Split(line, " ")
			l[1] = newName
			lines[i] = strings.Join(l, " ")
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(planet, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	logAccion("UpdateName", planet, city, newName)
}

// Modifica la cantidad de rebeldes de una ciudad del archivo de un planeta
func updateCityRebelds(planet string, city string, newValue int32) {
	input, err := ioutil.ReadFile(planet)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, city) {
			l := strings.Split(line, " ")
			l[2] = strconv.Itoa(int(newValue))
			lines[i] = strings.Join(l, " ")
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(planet, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	logAccion("UpdateNumber", planet, city, strconv.Itoa(int(newValue)))
}

//-------------------------------------------------------------------------------------------------
//							Funcionalidades de Servidor

// Retorna el clock correspondiente a planet
func (s *server) GetClock(ctx context.Context, in *fp.PlanetData) (*fp.PlanetClock, error) {
	planet := in.GetPlanet()

	return &fp.PlanetClock{Clock: clockMap[planet]}, nil
}

// Crea una ciudad en un planeta
func (s *server) AddCity(ctx context.Context, in *fp.NewCity) (*fp.CityStatus, error) {
	planet := in.GetPlanet()
	city := in.GetCity()
	rebelds := in.GetRebelds()

	addCity(planet, city, rebelds)

	if _, ok := clockMap[planet]; !ok {
		clockMap[planet] = []int32{0, 0, 0}
	}
	clockMap[planet][index] = clockMap[planet][index] + 1

	return &fp.CityStatus{Clock: clockMap[planet]}, nil
}

// Destruye una ciudad en el planeta planet
func (s *server) DeleteCity(ctx context.Context, in *fp.DelCity) (*fp.CityStatus, error) {
	planet := in.GetPlanet()
	city := in.GetCity()

	deleteCity(planet, city)

	clockMap[planet][index] = clockMap[planet][index] + 1

	return &fp.CityStatus{Clock: clockMap[planet]}, nil
}

// Hace update a una ciudad
func (s *server) UpdateCity(ctx context.Context, in *fp.UpCity) (*fp.CityStatus, error) {
	planet := in.GetPlanet()
	city := in.GetCity()
	flag := in.GetFlag()

	if flag == 0 {
		newValue := in.GetNum()
		updateCityRebelds(planet, city, newValue)
	} else {
		newName := in.GetNewname()
		updateCityName(planet, city, newName)
	}

	clockMap[planet][index] = clockMap[planet][index] + 1
	return &fp.CityStatus{Clock: clockMap[planet]}, nil
}

// Retorna el numero de rebeldes en una ciudad, en un planeta
func (s *server) GetNumber(ctx context.Context, in *fp.CityData) (*fp.CityRes, error) {
	planet := in.GetPlanet()
	city := in.GetCity()

	rebelds := getNumbers(planet, city)

	return &fp.CityRes{Clock: clockMap[planet], Rebelds: rebelds}, nil
}

// Retorna la lista de planetas que el servidor maneja
func (s *server) RequestPlanetList(ctx context.Context, in *emptypb.Empty) (*fp.PlanetList, error) {
	planets := make([]string, len(clockMap))

	i := 0
	for planet := range clockMap {
		planets[i] = planet
		i++
	}

	return &fp.PlanetList{List: planets}, nil
}

// Retorna el clock y el contenido del archivo log respectivo a un planeta especifico
func (s *server) RequestLog(ctx context.Context, in *fp.LogReq) (*fp.Log, error) {
	planet := in.GetName()

	clock := clockMap[planet]

	input, err := ioutil.ReadFile(strings.Join([]string{planet, ".log"}, ""))
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	return &fp.Log{LogFile: lines, Clock: clock}, nil
}

func (s *server) UpdateFile(ctx context.Context, in *fp.NewData) (*fp.Status, error) {
	file := in.GetFile()
	clock := in.GetClock()
	planet := in.GetPlanet()

	clockMap[planet] = clock

	updateFileContent(file, planet)
	// TODO: AGREGAR UNA FUNCION QUE HAGA RESET A LOS LOG (TODOS!!!!!!!)
	logName := strings.Join([]string{planet, "log"}, ".")
	updateFileContent([]string{}, logName)
	return &fp.Status{StatusFlag: 0}, nil
}

// Pone el servidor fulcrum a escuchar
func ListenFulcrumServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fp.RegisterFulcrumServerServer(s, &server{})
	log.Printf("[*] Fulcrum %d gRPC server listening at %v", index+1, lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//-----------------------------------------------------------------------------------------------

//-----------------------------------------------------------------------------------------------
//								Nodo Coordinador

// Obtiene la lista de los planetas registrados en el servidor en address
func GetPlanetList(address string) []string {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fp.NewFulcrumServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	r, err := c.RequestPlanetList(ctx, &emptypb.Empty{})

	if err != nil {
		log.Fatalf("No se pudo recibir lista de planetas desde %s: %v", address, err)
	}

	planetList := r.GetList()

	fmt.Printf("Se ha recibido la lista de planetas desde la direccion %s", address)
	return planetList
}

// Obtiene el clock y el archivo log de un planeta en el servidor en address
func GetLogByPlanet(planet string, address string) ([]int32, []string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fp.NewFulcrumServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	r, err := c.RequestLog(ctx, &fp.LogReq{Name: planet})

	if err != nil {
		log.Fatalf("No se pudo recibir el registro del planeta %s desde %s: %v", planet, address, err)
	}

	clock := r.GetClock()
	file := r.GetLogFile()

	return clock, file
}

// Ejecuta los comandos en un archivo log sobre el planeta local
func consistenciaINADOR(file []string, planet string) {
	for _, line := range file {
		lineArray := strings.Split(line, " ")
		command := lineArray[0]
		city := lineArray[2]

		if command == "AddCity" {
			value, _ := strconv.Atoi(lineArray[3])
			addCity(planet, city, int32(value))
		} else if command == "UpdateName" {
			newName := lineArray[3]
			updateCityName(planet, city, newName)
		} else if command == "UpdateNumber" {
			newValue, _ := strconv.Atoi(lineArray[3])
			updateCityRebelds(planet, city, int32(newValue))
		} else {
			deleteCity(planet, city)
		}
	}
}

// Envia la informacion nueva para que los otros servers la almacenen
func sendUpdatedPlanetData(planet string, clock []int32, address string) {
	lines := readFileContent(planet)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fp.NewFulcrumServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	r, err := c.UpdateFile(ctx, &fp.NewData{Clock: clock, File: lines, Planet: planet})

	if err != nil {
		log.Fatalf("No se pudo enviar informacion actualizada del planeta %s a %s: %v", planet, address, err)
	}
	r.GetStatusFlag()
	fmt.Printf("Se ha enviado correctamente informacion para actualizar el planeta %s a %s", planet, address)
}

// Actualiza archivos locales en base a la data almacenada en los otros servidores fulcrum
func CheckConsistency() {
	for _, fulcrumAddress := range fulcrum_addresses {
		planetList := GetPlanetList(fulcrumAddress)

		for _, planet := range planetList {
			clock, fileArray := GetLogByPlanet(planet, fulcrumAddress)
			consistenciaINADOR(fileArray, planet)

			v1 := clockMap[planet][index]
			v2 := clock[index]

			if v1 > v2 {
				clockMap[planet] = []int32{v1, v1, v1}
			} else {
				clockMap[planet] = []int32{v2, v2, v2}
			}
		}
	}

	for _, fulcumAddress := range fulcrum_addresses {
		for planet, clockPlanet := range clockMap {
			sendUpdatedPlanetData(planet, clockPlanet, fulcumAddress)
		}
	}
}

//-----------------------------------------------------------------------------------------------

func main() {
	ListenFulcrumServer()
}
