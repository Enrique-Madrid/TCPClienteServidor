package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

var (
	i            bool
	canalUltimo  string
	nombreUltimo string
)

//Función encargada de recibir mensajes de otro cliente
func cargarMensaje(conn net.Conn, canal string) {
	for {
		b := make([]byte, 2e+8)
		clienteMsg, err := conn.Read(b)
		if err != nil {
			fmt.Println("Conexión con el servidor perdida")
			os.Exit(0)
			break
		}
		if i == false {
			arrayMeta := strings.Split(string(b[:clienteMsg]), ":")
			//Determina el canal del mensaje
			canalUltimo = arrayMeta[1]
			//Determina el nombre del cliente
			nombreUltimo = arrayMeta[0]
			i = true
		} else {
			if canal == canalUltimo {
				fmt.Println("Se recibió el archivo", nombreUltimo, "del canal #", canalUltimo)
				//Crea un archivo con el nombre del canal y el nombre del cliente
				destino, err := os.OpenFile("./"+canalUltimo+"_"+nombreUltimo, os.O_WRONLY|os.O_CREATE, 0666)
				_, err = destino.Write(b[:clienteMsg])
				if err != nil {
					log.Fatal(err)
				}
				destino.Close()
				i = false
			}
		}

	}
}

func main() {
	serConn, err := net.Dial("tcp", "localhost:5555")
	if err != nil {
		log.Fatal(err)
	}

	//Un ciclo para escribir mensajes
	for {

		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		comando := strings.Split(reader.Text(), " ")

		switch comando[0] {
		case "exit":
			//En dado caso el usuario ingrese exit, se cierra el programa
			serConn.Close()
			os.Exit(0)
		case "./client":
			//En dado caso el usuario ingrese ./client, se crea un nuevo canal
			if comando[1] == "log" && comando[2] != "" {
				canal := comando[2]
				logCliente(canal, serConn)

			} else {
				fmt.Println("Ingrese ./client log <canal>")
			}

		default:
			fmt.Println("Comando no reconocido")
		}

	}
}

//Se encarga de ejecutar los demas comandos luego de logearse como <cliente>
func logCliente(canal string, conn net.Conn) {
	//Se hace un hilo para recibir mensajes de otro cliente
	go cargarMensaje(conn, canal)
	fmt.Println("Logeado en el canal #", canal)
	for {

		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		comando := strings.Split(reader.Text(), " ")

		switch comando[0] {
		case "exit":
			conn.Close()
			os.Exit(0)
		case "./client":
			if comando[1] == "channel" && comando[2] != "" {
				canalAEnviar := comando[2]
				fmt.Println("El canal #", canalAEnviar, "se ha seleccionado para recibir paquetes")
				enviarMensaje(canalAEnviar, conn)
			} else {
				fmt.Println("./client channel <canal> \n Ingrese el canal al quiere enviar")
			}
		default:
			fmt.Println("Comando no reconocido")
		}

	}
}

func enviarMensaje(canal string, conn net.Conn) {

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	comando := strings.Split(reader.Text(), " ")

	switch comando[0] {
	case "exit":
		conn.Close()
		os.Exit(0)
	case "./client":
		if comando[1] == "send" && comando[2] != "" {
			enviarArchivo(conn, canal, comando[2])

		} else {
			fmt.Println("./client send <archivo>")
		}
	default:
		fmt.Println("Comando no reconocido")
	}
}

func enviarArchivo(conn net.Conn, canal, dir string) {
	arrayNom := strings.Split(dir, "/")
	nombreArchivo := arrayNom[len(arrayNom)-1]
	fmt.Println("Enviando archivo:", nombreArchivo, "al canal:", canal)
	archivo, err := os.Open(dir)
	datosArchivo, err := ioutil.ReadAll(archivo)
	if err != nil {
		log.Fatal(err)
	}

	metaDatos := nombreArchivo + ":" + canal
	conn.Write([]byte(metaDatos))
	conn.Write([]byte(datosArchivo))
	archivo.Close()
}
