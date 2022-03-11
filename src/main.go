package main

import (
	"database/sql"
	"facturacion/controller"
	"facturacion/model"
	"fmt"
)

// This method delete a row
func testDelete(dbConnection *sql.DB) {
	rows, err := controller.Delete(dbConnection, "DELETE from album where id = ?", []interface{}{7})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("INFO: Rows affected -> %d\n", rows)
}

// This method update a dummy value
func testUpdate(dbConnection *sql.DB) {
	rows, err := controller.Update(dbConnection, "UPDATE album set title = ? where artist = ?", []interface{}{"UPDATE_TITULO", "John Coltrane"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Info: Rows affected -> %d\n", rows)
}

// This method inserts a dummy value
func testInsert(dbConnection *sql.DB) {
	var price float64 = 1.54556
	rows, err := controller.Add(dbConnection, "INSERT INTO album(title, artist, price) VALUES(?, ?, ?)", []interface{}{"Titulo", "ARTISTA", price})
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
	}
	fmt.Printf("Info: Index of the new row -> %d\n", rows)
}

// This method queries all albums
func testSelectAll(dbConnection *sql.DB) {
	rows := controller.Get(dbConnection, "SELECT * FROM FACTURACABECERA", []interface{}{})
	facturas := model.MapRowsToFacturaCabecera(rows)
	fmt.Println(facturas)

}

func main() {
	dbConnection := controller.Connect()
	testSelectAll(dbConnection)
}
