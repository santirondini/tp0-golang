package main

import (
	"client/globals"
	"client/utils"
	// "fmt"
	"log"
	"os"
	"bufio"
)

func textoLeido () string {
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')
	return texto
}

func main() {
	
	
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	// log.Println("Hola soy un log")
	
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// validar que la config este cargada correctamente

	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	// loggeamos el valor de la config
	
	log.Println(globals.ClientConfig)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él	

	// enviar un mensaje al servidor con el valor de la config

	// utils.EnviarMensaje(globals.ClientConfig.Ip,globals.ClientConfig.Puerto,globals.ClientConfig.Mensaje)
	// me tir aun error de enviando mensaje a la ip y al puerto



	// leer de la consola el mensaje
		for {
			palabra := textoLeido()
			if palabra == " " || palabra == "\n" {
				break
			}
			log.Println(palabra)
		}

	// generamos un paquete y lo enviamos al servidor
	
	utils.GenerarYEnviarPaquete()

}