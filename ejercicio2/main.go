/*
Ejercicio 2 - Imprimir datos

A continuación, vamos a crear un archivo “customers.txt”
con información de los clientes del estudio.

Ahora que el archivo sí existe, el panic no debe ser lanzado.

Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir
los datos que contenga. En el caso de no poder leerlo, se debe lanzar un “panic”.

Recordemos que siempre que termina la ejecución, independientemente del resultado,
debemos tener un “defer” que nos indique que la ejecución finalizó.

También recordemos cerrar los archivos al finalizar su uso.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(path string) (file *os.File) {

	file, err := os.Open(path)

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	return
}

func main() {

	fmt.Println("Reading file customer.txt")

	file := readFile("./customer.txt")

	defer func() {
		fmt.Println("Ejecución Finalizada")
		file.Close()
	}()

	bytes, _ := io.ReadAll(file)

	fmt.Println(string(bytes))

}
