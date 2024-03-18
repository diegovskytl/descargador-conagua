package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// agregar diccionario

func main() {

	estadosMx := map[int]string{
		0:  "ags",  // Aguascalientes
		1:  "bc",   // Baja California - 2
		2:  "bcs",  // Baja California Sur - 3
		3:  "camp", // Campeche - 4
		4:  "coah",
		5:  "col",
		6:  "chis", // Chiapas - 7
		7:  "chih", // Chihuahua - 8
		8:  "cdmx", // Ciudad de México
		9:  "dgo",  // Durango
		10: "gto",  // Guanajuato
		11: "gro",  // Guerrero - 12
		12: "hgo",  // Hidalgo
		13: "jal",  // Jalisco
		14: "edomex",
		15: "mich", // Michoacán - 16
		16: "mor",  // Morelos - 17
		17: "nay",  // Nayarit
		18: "nl",   // Nuevo León
		19: "oax",  // Oaxaca - 20
		20: "pue",  // Puebla - 21
		21: "qro",  // Querétaro
		22: "qr",   // Quintana Roo
		23: "slp",  // San Luis Potosí
		24: "sin",  // Sinaloa
		25: "son",  // Sonora - 26
		26: "tab",  // Tabasco
		27: "tams", // Tamaulipas
		28: "tlax", // Tlaxcala - 29
		29: "ver",  // Veracruz
		30: "yuc",  // Yucatán
		31: "zac",  // Zacatecas
	}

	const baseURL string = "https://smn.conagua.gob.mx/tools/RESOURCES/Mensuales/"

	fmt.Println("El URL base es ", baseURL)
	for i := range 32 {

		fmt.Println("Descargando archivos para", estadosMx[i])
		var edoURL string = baseURL + estadosMx[i] + "/"

		for j := range 5 {
			prefijoTXT := fmt.Sprintf("%05d", i+1)
			sufijoTXT := fmt.Sprintf("%03d", j)
			nombreArch := prefijoTXT + sufijoTXT
			txtFileUrl := edoURL + nombreArch + ".TXT"
			fmt.Println("Descargando desde", txtFileUrl)

			response, err := http.Get(txtFileUrl)
			if err != nil {
				fmt.Println("Error al obtener el archivo", estadosMx[i]+"/"+nombreArch+".TXT")
				continue
			}
			defer response.Body.Close()

			//procesamiento: -> revisar cómo empieza, quitar espacios del principio, colocar comas antes de cada espacio
			reader := bufio.NewReader(response.Body)

			firstLine, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Error al leer la primera línea del http response")
				continue
			}

			if !strings.HasPrefix(firstLine, "<") {

				var locationSuffix string
				var trimedLines string

				//limpieza
				limpiador := bufio.NewScanner(reader)

				numLinea := 0

				for limpiador.Scan() {

					var CSVline string

					linea := limpiador.Text()
					linea = strings.TrimSpace(linea)

					palabras := strings.Fields(linea)

					if strings.HasPrefix(linea, "DESV.") {
						CSVline = linea
					} else {

						for _, palabra := range palabras {
							CSVline += palabra + ", "
						}

					}

					if numLinea == 5 {
						locationSuffix = linea //agregar localidad y municipio al nombre del archivo
					}
					trimedLines += CSVline + "\n"
					numLinea++
				}

				file, err := os.Create(estadosMx[i] + "_" + locationSuffix + "_" + nombreArch + ".csv")
				if err != nil {
					fmt.Println("Error al crear el archivo", nombreArch)
					continue
				}
				defer file.Close()

				_, err = file.WriteString(trimedLines)
				if err != nil {
					fmt.Println("hostia, falló la copia")
					continue
				}

				fmt.Println("Archivo descargado correctamente:", nombreArch)
			} else {
				fmt.Println("el archivo no eran datos climáticos")
			}
		}
	}
}
