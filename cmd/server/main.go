package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	contraseña := "micontraseña"

	contraseñaHasheada, err := bcrypt.GenerateFromPassword([]byte(contraseña), bcrypt.DefaultCost)

	if err != nil {

		log.Fatal(err)
		return
	}

	fmt.Println("contraseña hasheada: ", string(contraseñaHasheada))
	contraseñaVerificacion := "micontraseña"

	err = bcrypt.CompareHashAndPassword(contraseñaHasheada, []byte(contraseñaVerificacion))

	if err != nil {
		fmt.Println("Contraseña incorrecta")
	} else {
		fmt.Println("Contraseña correcta")
	}

}
