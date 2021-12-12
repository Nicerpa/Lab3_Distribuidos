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

	fp "github.com/Lab3/FulcrumProto"
	"google.golang.org/grpc"
)

type server struct {
	fp.UnimplementedFulcrumServerServer
}

const (
	index = 0
	port  = ":50001" // Poner IP completa para el deploy
)

var clockMap = make(map[string][]int32)

// -----------------------------Funciones Utiles --------------------------------------------------
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

func addCity(planetName string, cityName string, rebels int32) {
	f, err := os.OpenFile(planetName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s %s %d\n", planetName, cityName, rebels)); err != nil {
		log.Println(err)
	}
}

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
}

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
}

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
}

//-------------------------------------------------------------------------------------------------

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

func main() {
	ListenFulcrumServer()
}
