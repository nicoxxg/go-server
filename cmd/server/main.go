package main

import (
	"database/sql"
	"fmt"
	"log"

	handlerClient "github.com/nicoxxg/go-server/cmd/server/handler/cliente"

	handlerTurno "github.com/nicoxxg/go-server/cmd/server/handler/turno"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/nicoxxg/go-server/internal/domain/cliente"
	"github.com/nicoxxg/go-server/internal/domain/turno"
)

const (
	puerto = ":8080"
)

func main() {
	db := connectDB()
	router := gin.New()
	router.Use(gin.Recovery())
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	turnoRepository := turno.NewTurnoRepository(db)

	turnoService := turno.NewTurnoService(turnoRepository)

	turnoController := handlerTurno.NewTurnoController(turnoService)

	clienteRepository := cliente.NewClienteRepository(db)

	clientService := cliente.NewClientService(clienteRepository)

	clientController := handlerClient.NewClientController(clientService)

	router.GET("/clientes", clientController.FindAll())

	router.POST("/turno", turnoController.Save())

	router.GET("/turno", turnoController.FindAllTurnos())

	router.GET("/cliente/:id", clientController.FindClientById())

	router.GET("/cliente/email", clientController.FindClienteByEmail())

	router.POST("/cliente", clientController.SaveClient())

	router.PATCH("/cliente/update", clientController.UpdateClient())

	if err := router.Run(puerto); err != nil {
		panic(err)
	}

}

// contraseña := "micontraseña"

// 	contraseñaHasheada, err := bcrypt.GenerateFromPassword([]byte(contraseña), bcrypt.DefaultCost)

// 	if err != nil {

// 		log.Fatal(err)
// 		return
// 	}

// 	fmt.Println("contraseña hasheada: ", string(contraseñaHasheada))
// 	contraseñaVerificacion := "micontraseña"

// 	err = bcrypt.CompareHashAndPassword(contraseñaHasheada, []byte(contraseñaVerificacion))

// 	if err != nil {
// 		fmt.Println("Contraseña incorrecta")
// 	} else {
// 		fmt.Println("Contraseña correcta")
// 	}

func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = "root"
	dbPassword = "root"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "go_server"
	cadenaConexion := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=UTF8", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", cadenaConexion)

	if err != nil {
		panic(err)
	}

	// checkea la coneccion
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}
