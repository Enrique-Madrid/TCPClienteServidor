package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//Aquí se guardan todas las conexiones, mensajes en arrays
var (
	e            bool
	canalUltimo  string
	nombreUltimo string
	conns        []net.Conn
	connCh       = make(chan net.Conn)
	closeCh      = make(chan net.Conn)
	msgCh        = make(chan string)
)

func baseDeDatos() (conexionDB *sql.DB) {
	driver := "mysql"
	user := "root"
	pass := ""
	dbName := "servidordatos"
	conexionDB, err := sql.Open(driver, user+":"+pass+"@tcp(127.0.0.1)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return conexionDB
}

//Este método se encarga de cargar los mensajes que llegan del cliente
func cargarMensaje(conn net.Conn) {
	for {
		b := make([]byte, 2e+8)
		mensaje, err := conn.Read(b)
		if err != nil {
			break
		}
		if !e {
			arrayMSG := strings.Split(string(b[:mensaje]), ":")
			//Determina el canal del mensaje
			canalUltimo = arrayMSG[1]
			//Determina el nombre del cliente
			nombreUltimo = arrayMSG[0]
			e = true
		} else {
			var tamaño float32 = float32(len(b[:mensaje])) / 1e+6
			fmt.Println(tamaño)
			insertarDatos(canalUltimo, nombreUltimo, tamaño)
			e = false
		}
		msg := string(b[:mensaje])
		msgCh <- msg
		//Luego de cargar el mensaje en el servidor, lo devuelve al cliente
		pubMsg(conn, msg)

	}
	closeCh <- conn
}

//Inicia la interface web
func iniciarWebPage() {
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	//Abre la página web php llamada index.php
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.php")
	})
	http.ListenAndServe(":5556", nil)
}

func insertarDatos(canal, nombre string, peso float32) {
	//Aquí se crea la base de datos
	db := baseDeDatos()
	//Aquí se crea la tabla de mensajes
	insert, err := db.Prepare("INSERT INTO datosenviados2(id, canal, nombre, peso) VALUES('','" + canal + "','" + nombre + "','" + fmt.Sprintf("%f", peso) + "')")
	if err != nil {
		log.Fatal(err)
	}
	insert.Exec()
	log.Println("Base de datos creada")
}

func main() {
	//Aquí se crea el servidor
	server, err := net.Listen("tcp", ":5555")
	go iniciarWebPage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Servidor Iniciado...")
	//Aquí se crea un hilo que se encarga de aceptar conexiones
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatal(err)
			}

			conns = append(conns, conn)
			connCh <- conn
		}
	}()

	//Aquí se encarga de cargar los mensajes que llegan del cliente
	for {
		select {
		case conn := <-connCh:
			go cargarMensaje(conn)
			fmt.Println("Cliente conectado")
		case msg := <-msgCh:
			_ = msg
			println("Archivo enviado con éxito")
		case conn := <-closeCh:
			fmt.Println("Se desconectó un cliente")
			removerConn(conn)
		}
	}

}

//Esta función se encarga de enviar los mensajes a los clientes
func pubMsg(conn net.Conn, msg string) {
	for i := range conns {
		if conns[i] != conn {
			conns[i].Write([]byte(msg))
		}
	}

}

//En dado caso un usuario se desconecte, esta función borra su conexión del array
func removerConn(conn net.Conn) {
	var i int
	for i = range conns {
		if conns[i] == conn {
			break
		}

	}
	conns = append(conns[i:], conns[:i+1]...)
}
