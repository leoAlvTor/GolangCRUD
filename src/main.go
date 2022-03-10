package main

import (
	"database/sql"
	"facturacion/controller"
	"facturacion/model"
	"fmt"
)

func main() {
	var dbConnection *sql.DB
	dbConnection = controller.Connect()

	var albums []model.Album

	albums, _ = controller.GetAll(dbConnection, "SELECT * FROM album WHERE artist = ? and price >= ?", []interface{}{"John Coltrane", "56"})
	fmt.Println(albums)
	controller.Disconnect()
}
