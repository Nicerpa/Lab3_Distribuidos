package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	bp "github.com/Lab3/BrokerProto"
	fp "github.com/Lab3/FulcrumProto"
	"google.golang.org/grpc"
)

const (
	ip_f1 = "localhost:50001"
	ip_f2 = "localhost:50002"
	ip_f3 = "localhost:50003"
	port  = ":60000" // Poner IP completa para el deploy
)

type server struct {
	bp.UnimplementedBrokerServerServer
}

// Consulta al Fulcrum seleccinado por la cantidad de rebeldes en city, planet y su clock
func GetPlanetNumbers(address string, planet string, city string) (int32, []int32) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fp.NewFulcrumServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	r, err := c.GetNumber(ctx, &fp.CityData{Planet: planet, City: city})
	if err != nil {
		log.Fatalf("No se pudo obtener la cantidad de rebeldes en el planeta %s, ciudad %s %v", planet, city, err)
	}

	return r.GetRebelds(), r.GetClock()
}

// Obtiene el valor del clock del planeta planet en el servidor fulcrum en la direccion address
func GetFulcrumClock(address string, planet string) []int32 {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fp.NewFulcrumServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	r, err := c.GetClock(ctx, &fp.PlanetData{Planet: planet})

	if err != nil {
		log.Fatalf("No se pudo recibir clock del planeta %s: %v\n", planet, err)
	}
	clock := r.GetClock()
	fmt.Printf("Se ha recibido el clock desde %s del planeta %s con valor %v\n", address, planet, clock)

	return clock
}

// Retorna el indice del reloj al cual corresponde la direccion ip
func GetFulcrumNum(address string) int {
	if address == ip_f1 {
		return 0
	} else if address == ip_f2 {
		return 1
	} else {
		return 2
	}
}

// Retorna la ip del Fulcrum con la version mas nueva
func LatestVersion(clock_f1 int32, clock_f2 int32, clock_f3 int32, clock int32, ip string) string {
	if clock_f1 > clock && clock_f1 >= clock_f2 && clock_f1 >= clock_f3 {
		return ip_f1
	} else if clock_f2 > clock && clock_f2 >= clock_f1 && clock_f2 >= clock_f3 {
		return ip_f2
	} else if clock_f3 > clock && clock_f3 >= clock_f2 && clock_f1 <= clock_f3 {
		return ip_f3
	} else {
		return ip
	}
}

// Retorna la ip donde se encuentra la ultima version de un planeta
func getSelectedIP(planet string, ip string, clock []int32) string {
	if ip == "" {
		ran := rand.Intn(3)
		if ran == 0 {
			return ip_f1
		} else if ran == 1 {
			return ip_f2
		} else {
			return ip_f3
		}
	} else {
		clock_f1 := GetFulcrumClock(ip_f1, planet)
		clock_f2 := GetFulcrumClock(ip_f2, planet)
		clock_f3 := GetFulcrumClock(ip_f3, planet)

		var index int

		if ip != "" {
			index = GetFulcrumNum(ip)
		} else {
			index = 0
		}

		return LatestVersion(clock_f1[index], clock_f2[index], clock_f3[index], clock[index], ip)
	}
}

// Atiende consulta de Leia sobre el numero de rebeldes en una ciudad
func (s *server) GetNumberRebelds(ctx context.Context, in *bp.LeiaReq) (*bp.Rebelds, error) {
	planet := in.GetPlanet()
	city := in.GetCity()
	ip := in.GetIp()
	clock := in.GetClock()

	selected_ip := getSelectedIP(planet, ip, clock)
	rebelds, new_clock := GetPlanetNumbers(selected_ip, planet, city)

	return &bp.Rebelds{Rebelds: rebelds, Ip: selected_ip, Clock: new_clock}, nil
}

// Atiende peticion de Informantes
func (s *server) GetIPCity(ctx context.Context, in *bp.CityData) (*bp.CityRes, error) {
	planet := in.GetPlanet()
	ip := in.GetIp()
	clock := in.GetClock()

	selected_ip := getSelectedIP(planet, ip, clock)

	fmt.Printf("Se ha enviado la ip %s a un informante\n", selected_ip)
	return &bp.CityRes{IpDes: selected_ip}, nil
}

// Pone el servidor a escuchar por solicitudes
func ListenBrokerServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bp.RegisterBrokerServerServer(s, &server{})
	log.Printf("[*] Mos Eisley gRPC server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

func main() {
	rand.Seed(1)
	ListenBrokerServer()
}
