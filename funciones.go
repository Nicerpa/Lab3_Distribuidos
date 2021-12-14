package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	myfile *os.FileInfo
	e      error
)

func main() {

	// GetNumberRebelds("test.txt", "xupalo", "jorgePetardo")
	// menuInformantes()
	// updateCityName("test", "ciudad_9", "xuxetumare")
	// updateCityRebelds("test", "ciudad_9", 100)
	// truncate("test.txt")

}
func updateCityName(nombrePlaneta string, nombreCiudad string, newName string) {
	nombreArchivo := nombrePlaneta + ".txt"
	file, err := os.Open(nombreArchivo)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	sentencia := ""
	sentencia += nombrePlaneta + " " + nombreCiudad
	var text []string
	pertenencia := 0

	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), sentencia) {
			splitear := strings.Split(fileScanner.Text(), " ")
			sentencia = nombrePlaneta + " " + newName + " " + splitear[2]
			text = append(text, sentencia)
			pertenencia += 1
		} else {

			text = append(text, fileScanner.Text())
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()

	if pertenencia != 0 {
		file2, err := os.Create(nombreArchivo)

		if err != nil {
			log.Fatalf("Error when opening file: %s", err)
		}

		for i := 0; i < len(text); i++ {
			if i == len(text)-1 {

				file2.WriteString(text[i])
			} else {

				file2.WriteString(text[i] + "\n")
			}
		}

		file2.Close()
	} else {
		fmt.Println("No existe la ciudad o el planeta indicados")
	}

}

func updateCityRebelds(nombrePlaneta string, nombreCiudad string, newValue int) {
	nombreArchivo := nombrePlaneta + ".txt"
	file, err := os.Open(nombreArchivo)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	sentencia := ""
	sentencia += nombrePlaneta + " " + nombreCiudad
	var text []string
	pertenencia := 0

	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), sentencia) {

			sentencia += " " + strconv.Itoa(newValue)
			text = append(text, sentencia)
			pertenencia += 1
		} else {

			text = append(text, fileScanner.Text())
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()

	if pertenencia != 0 {
		file2, err := os.Create(nombreArchivo)

		if err != nil {
			log.Fatalf("Error when opening file: %s", err)
		}

		for i := 0; i < len(text); i++ {
			if i == len(text)-1 {

				file2.WriteString(text[i])
			} else {

				file2.WriteString(text[i] + "\n")
			}
		}

		file2.Close()
	} else {
		fmt.Println("No existe la ciudad o el planeta indicados")
	}

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
			//getRebelds()
		}
	}
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
				// agregarCiudad()
			}
			fmt.Println("Ciudad ", nombreCiudad, " agregada con ", numeroRebeldes, " rebeldes")

		} else if numero == 2 {
			fmt.Println("Ingrese nuevo nombreCiudad")
			fmt.Scanf("%s\n", &nombreCiudadNueva)
			updateCityName(nombrePlaneta, nombreCiudad, nombreCiudadNueva)
			fmt.Println("Ciudad actualizada ", nombreCiudad, " -> ", nombreCiudadNueva)

		} else if numero == 3 {
			fmt.Println("Ingrese numeroRebeldes")
			fmt.Scanf("%d\n", &numeroRebeldes)
			updateCityRebelds(nombrePlaneta, nombreCiudad, numeroRebeldes)
			fmt.Println("Rebeldes actualizados a ", numeroRebeldes)
		} else if numero == 4 {
			//eliminarCiudad()
			fmt.Println("Ciudad ", nombreCiudad, " eliminada")
		}
	}

}
