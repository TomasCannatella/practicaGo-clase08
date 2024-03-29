/*
Ejercicio 1 - Datos de clientes
Un estudio contable necesita acceder a los datos de sus empleados para poder realizar
distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo .txt.

Tendrás que desarrollar la funcionalidad para poder leer el archivo .txt
que nos indica el cliente, sin embargo, no han pasado el archivo a leer por nuestro programa.

Desarrollá el código necesario para leer los datos del archivo llamado “customers.txt”
(recordá lo visto sobre el pkg “os”).
Dado que no contamos con el archivo necesario,
se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al intentar
leer un archivo que no existe, mostrando el mensaje “The indicated file was not found or is damaged”.

Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
*/

package main

import (
	"fmt"
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
	defer fmt.Println("Ejecución Finalizada")

	fmt.Println("Reading file customer.txt")
	file := readFile("ejericio1/customer.txt")
	fmt.Println(file)

}
