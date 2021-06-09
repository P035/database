package database

import (

	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Data struct {

	id int
	Nombre string
	Contrasena string
}

func Start() *sql.Tx{

	db, err := sql.Open("mysql", "root:Naranjo7854@/relacional")
	fmt.Println("Conectando con la base de datos.")
	if err != nil {

		fmt.Println("Error conectando a la base de datos", err)
		os.Exit(1)
	}
	fmt.Println("Iniciando transaxión.")
	tx, err := db.Begin()
	if err != nil {

		fmt.Println("Error iniciando transaxión", err)
	}
	return tx
}

func Insertar(tx *sql.Tx) {

	stmt, err := tx.Prepare("INSERT INTO usuarios(ID, Nombre, Contraseña) VALUES(?, ?, ?)")
	fmt.Println(err)
	result, err := stmt.Exec("90", "peo", "Poto")
	fmt.Println(err)
	fmt.Println(result.LastInsertId())
}

func Tabla(tx *sql.Tx, query string) Data{

	data, err := tx.Query(query)
	if err != nil{

		fmt.Println("Error ejecutando Query", err)
	}
	var Lectura Data
	data.Next()
	data.Scan(&Lectura.id, &Lectura.Nombre, &Lectura.Contrasena)
	return Lectura
}
