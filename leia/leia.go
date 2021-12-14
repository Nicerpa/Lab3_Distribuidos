package main

import (
	"context"
	"fmt"
	"log"
	"time"

	bp "github.com/Lab3/BrokerProto"
	"google.golang.org/grpc"
)

const (
	brokerAddr = "localhost:60000"
	fulcrum_1  = "localhost:50001"
)

var clockMap = make(map[string][]int32)

var ipMap = make(map[string]string)

func GetNumbersRebelds(planet string, city string) int {
	conn, err := grpc.Dial(brokerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := bp.NewBrokerServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	ip, ok := ipMap[planet]

	if !ok {
		ip = fulcrum_1
	}

	clock, ok := clockMap[planet]

	if !ok {
		clock = []int32{0, 0, 0}
	}

	r, err := c.GetNumberRebelds(ctx, &bp.LeiaReq{Planet: planet, City: city, Ip: ip, Clock: clock})

	if err != nil {
		log.Fatalf("No se pudo recibir informacion de los rebeldes en %s: %v", planet, err)
	}
	clock = r.GetClock()
	rebels := r.GetRebelds()
	ip = r.GetIp()
	fmt.Printf("Se ha recibido el clock desde %s del planeta %s con valor %v", brokerAddr, planet, clock)

	clockMap[planet] = clock
	ipMap[planet] = ip

	return int(rebels)
}

func menuLeia() {
	nombrePlaneta := ""
	nombreCiudad := ""
	flag := true
	numero := 0

	for flag {
		fmt.Println("Ingresa una opción:")
		fmt.Println("0 - GetRebelds")
		fmt.Println("1 - Salir")
		fmt.Scanf("%d\n", &numero)

		if numero == 1 {
			fmt.Println("Saliendo...")
			return
		} else {

			fmt.Println("Ingrese nombrePlaneta")
			fmt.Scanf("%s\n", &nombrePlaneta)
			fmt.Println("Ingrese nombreCiudad")
			fmt.Scanf("%s\n", &nombreCiudad)
			numeroRebeldes := GetNumbersRebelds(nombrePlaneta, nombreCiudad)
			fmt.Println("El numero de rebeldes en la ciudad solicitada es: ", numeroRebeldes)
		}
	}
}

func main() {
	menuLeia()
}
