package utils

import (
	"bufio"
	"bytes"
	"client/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func IniciarConfiguracion(filePath string) *globals.Config {
	var config *globals.Config
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func LeerConsola() {
	// Leer de la consola
	reader := bufio.NewReader(os.Stdin) 
	log.Println("Ingrese los mensajes") 
	text, _ := reader.ReadString('\n')  
	log.Print(text)  
}

func GenerarYEnviarPaquete() {
	paquete := Paquete{}
	// Leemos y cargamos el paquete

	log.Printf("paqute a enviar: %+v", paquete)
	// Enviamos el paqute
}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("error codificando mensaje: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		log.Printf("error codificando mensajes: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensajes a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func ConfigurarLogger() {
	logFile, err := os.OpenFile("tp0.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil { // si err no es nil, significa que ocurrió un error al intentar abrir/crear el archivo
		panic(err) //detiene la ejecución del programa y muestra el error. El programa no continua si no se puede configuarar el logger
	}
	mw := io.MultiWriter(os.Stdout, logFile) // se crea un escritor Writer que envia los datos a multiples destinos. 
	log.SetOutput(mw) // se configura para que le logger envie mensajes a mw, que a suvez los redirige tanto a consola como a tp0.log
}

// os.Stdout representa la salida estandar (consola)
// mv es un escritor combinado que enviará cualquier mensaje de log tanto a consola como al archivo
// Ejemplo de uso: log.Println("Este es un mensaje de prueba")
