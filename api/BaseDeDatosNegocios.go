package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func obtenerBDNegocios() (*sql.DB, error) {
	var cadenaConexion string
	if ControladorBD == ControladorBDSQLite3 {
		cadenaConexion = CadenaConexionBDNegocios
	} else {
		cadenaConexion = fmt.Sprintf("%s:%s@%s/%s", UsuarioMySQL, PasswordMySQL, HostMySQL, NombreBaseDeDatosNegociosMySQL)
	}
	db, err := sql.Open(ControladorBD, cadenaConexion)
	if err != nil {
		log.Printf("Error abriendo base de datos maestra: %v", err)
		return nil, err
	}
	return db, nil
}
