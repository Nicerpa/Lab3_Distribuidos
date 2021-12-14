package main

import (
	"context"
	"fmt"
	"log"
	"time"

	bp "github.com/Lab3/BrokerProto"
	fp "github.com/Lab3/FulcrumProto"
	"google.golang.org/grpc"
)

const (
	brokerAddr = "localhost:60000"
)

var ipMap = make(map[string]string)

var clockMap = make(map[string][]int32)

func getIpCity(planeta string, clock []int32) string {
	conn, err := grpc.Dial(brokerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := bp.NewBrokerServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	ip, ok := ipMap[planeta]

	if !ok {
		ip = ""
	}

	r, err := c.GetIPCity(ctx, &bp.CityData{Planet: planeta, Clock: clock, Ip: ip})

	if err != nil {
		log.Fatalf("No se pudo obtener ip de servidor fulcrum desde el broker: %v", err)
	}

	return r.GetIpDes()
}

func agregarCiudad(planeta string, ciudad string, rebeldes int32) {
	clock, ok := clockMap[planeta]

	if !ok {
		clock = []int32{0, 0, 0}
	}

	ip := getIpCity(planeta, clock)

	fmt.Printf("Se ha recibido la ip %s desde el broker\n", ip)

	conn, err := grpc.Dial(ip, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fp.NewFulcrumServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	r, err := c.AddCity(ctx, &fp.NewCity{Planet: planeta, City: ciudad, Rebelds: rebeldes})

	if err != nil {
		log.Fatalf("No se pudo crear nueva ciudad %s en el planeta %s: %v\n", ciudad, planeta, err)
	}

	ipMap[planeta] = ip
	clockMap[planeta] = r.GetClock()

	fmt.Printf("Ciudad %s creada correctamente en el planeta %s con %d rebeldes\n", ciudad, planeta, rebeldes)
}

func updateCityName(nombrePlaneta string, nombreCiudad string, nombreCiudadNueva string) {
	clock, ok := clockMap[nombrePlaneta]

	if !ok {
		clock = []int32{0, 0, 0}
	}

	ip := getIpCity(nombrePlaneta, clock)

	conn, err := grpc.Dial(ip, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fp.NewFulcrumServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	r, err := c.UpdateCity(ctx, &fp.UpCity{Planet: nombrePlaneta, City: nombreCiudad, Newname: nombreCiudadNueva, Num: -1, Flag: 1})

	ipMap[nombrePlaneta] = ip
	clockMap[nombrePlaneta] = r.GetClock()

	fmt.Printf("Ciudad %s actualizada correctamenta a %s en el planeta %s\n", nombreCiudad, nombreCiudadNueva, nombrePlaneta)
}

func updateCityRebelds(nombrePlaneta string, nombreCiudad string, numeroRebeldes int) {
	clock, ok := clockMap[nombrePlaneta]

	if !ok {
		clock = []int32{0, 0, 0}
	}

	ip := getIpCity(nombrePlaneta, clock)

	conn, err := grpc.Dial(ip, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fp.NewFulcrumServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	r, err := c.UpdateCity(ctx, &fp.UpCity{Flag: 0, Planet: nombrePlaneta, City: nombreCiudad, Num: int32(numeroRebeldes), Newname: ""})

	ipMap[nombrePlaneta] = ip
	clockMap[nombrePlaneta] = r.GetClock()

	fmt.Printf("Numero de rebeldes actualizado a %d en la ciudad %s, planeta %s\n", numeroRebeldes, nombreCiudad, nombrePlaneta)
}

func eliminarCiudad(planeta string, ciudad string) {
	clock, ok := clockMap[planeta]

	if !ok {
		clock = []int32{0, 0, 0}
	}

	ip := getIpCity(planeta, clock)

	conn, err := grpc.Dial(ip, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fp.NewFulcrumServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	r, err := c.DeleteCity(ctx, &fp.DelCity{Planet: planeta, City: ciudad})

	ipMap[planeta] = ip
	clockMap[planeta] = r.GetClock()

	fmt.Printf("La ciudad %s en el planeta %s fue eliminada correctamente\n", ciudad, planeta)
}

func menuInformantes() {
	numero := 0
	nombrePlaneta := ""
	nombreCiudad := ""
	nombreCiudadNueva := ""
	bul := 0
	numeroRebeldes := 0
	flag := true

	for flag {

		fmt.Println("Ingresa una opción:")
		fmt.Println("1 - AddCity")
		fmt.Println("2 - UpdateName")
		fmt.Println("3 - UpdateNumber")
		fmt.Println("4 - DeleteCity")
		fmt.Println("5 - Salir")
		fmt.Scanf("%d\n", &numero)

		if numero == 5 {
			fmt.Println("Saliendo...")
			return
		}
		fmt.Println("Ingrese nombrePlaneta")
		fmt.Scanf("%s\n", &nombrePlaneta)
		fmt.Println("Ingrese nombreCiudad")
		fmt.Scanf("%s\n", &nombreCiudad)

		if numero == 1 {
			fmt.Println("¿Desea ingresar numeroRebeldes?")
			fmt.Println("Si - 0")
			fmt.Println("No - 1")
			fmt.Scanf("%d\n", &bul)
			if bul == 0 {
				fmt.Println("Ingrese numeroRebeldes")
				fmt.Scanf("%d\n", &numeroRebeldes)
			}
			agregarCiudad(nombrePlaneta, nombreCiudad, int32(numeroRebeldes))

		} else if numero == 2 {
			fmt.Println("Ingrese nuevo nombreCiudad")
			fmt.Scanf("%s\n", &nombreCiudadNueva)
			updateCityName(nombrePlaneta, nombreCiudad, nombreCiudadNueva)
			//fmt.Println("Ciudad actualizada ", nombreCiudad, " -> ", nombreCiudadNueva)

		} else if numero == 3 {
			fmt.Println("Ingrese numeroRebeldes")
			fmt.Scanf("%d\n", &numeroRebeldes)
			updateCityRebelds(nombrePlaneta, nombreCiudad, numeroRebeldes)
			//fmt.Println("Rebeldes actualizados a ", numeroRebeldes)
		} else if numero == 4 {
			eliminarCiudad(nombrePlaneta, nombreCiudad)
			//fmt.Println("Ciudad ", nombreCiudad, " eliminada")
		}
		numero = 0
		nombrePlaneta = ""
		nombreCiudad = ""
		nombreCiudadNueva = ""
		bul = 0
		numeroRebeldes = 0
	}

}

func main() {
	menuInformantes()
}
