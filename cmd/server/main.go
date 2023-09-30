package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	handlerClient "github.com/nicoxxg/go-server/cmd/server/handler/cliente"
	"github.com/nicoxxg/go-server/pkg/middleware"

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

	security := middleware.NewSecurity(clienteRepository)

	router.GET("/api/clientes", middleware.Verification(), middleware.RolVerification("admin"), clientController.FindAll())

	router.GET("/api/current", middleware.Verification(), clientController.GetClientCurrent())

	router.POST("/api/turno", turnoController.Save())

	router.GET("/api/turno", turnoController.FindAllTurnos())

	router.GET("/api/cliente/:id", clientController.FindClientById())

	router.GET("/api/cliente/email", clientController.FindClienteByEmail())

	router.POST("/api/cliente", clientController.SaveClient())

	router.PATCH("/api/cliente/update/:id", clientController.UpdateClient())

	router.POST("/api/login", security.Logger())

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
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = string(os.Getenv("DATABASE_USERNAME"))
	dbPassword = string(os.Getenv("DATABASE_PASSWORD"))
	dbHost = string(os.Getenv("DATABASE_HOST"))
	dbPort = string(os.Getenv("DATABASE_PORT"))
	dbName = string(os.Getenv("DATABASE_NAME"))
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
