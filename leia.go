package main

import (
	"context"
	"fmt"
	"log"
	"time"

	bp "github.com/Lab3/BrokerProto"
	"google.golang.org/grpc"
)

func GetNumbersRebelds(planet string, city string) (int, string, []int32) {
	// definir globalmente una constante address = "localhost:60000"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := bp.NewBrokerServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	r, err := c.GetNumberRebelds(ctx, &bp.LeiaReq{Planet: planet, City: city, Ip: ipMap[planet], Clock: clockMap[planet]})

	if err != nil {
		log.Fatalf("No se pudo recibir informacion de los rebeldes en %s: %v", planet, err)
	}
	clock := r.GetClock()
	rebels := r.GetRebelds()
	ip := r.GetIp()
	fmt.Printf("Se ha recibido el clock desde %s del planeta %s con valor %v", address, planet, clock)

	return int(rebels), ip, clock
}

func main() {

}
