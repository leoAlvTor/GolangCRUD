package main

import (
	"database/sql"
	"facturacion/controller"
	"facturacion/model"
	"fmt"
)

//This method inserts a dummy value
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
	var albums []interface{}
	albums, _ = controller.Get(dbConnection, "SELECT * FROM album", []interface{}{})
	for _, tmp := range albums {
		album, _ := tmp.(model.Album)
		fmt.Printf("Album ID: %d,\nalbum title: %s,\nalbum artist: %s,\nalbum price: %f\n", album.ID, album.Title, album.Artist, album.Price)
	}
}

// This method queries all albums from where artist is John Coltrane and price is greater or equal than $56
func testSelectWithParams(dbConnection *sql.DB) {
	var albums []interface{}
	albums, _ = controller.Get(dbConnection, "SELECT * FROM album WHERE artist = ? and price >= ?", []interface{}{"John Coltrane", "56"})
	for _, tempAlbum := range albums {
		album, _ := tempAlbum.(model.Album)
		fmt.Printf("Album ID: %d,\nalbum title: %s,\nalbum artist: %s,\nalbum price: %f\n", album.ID, album.Title, album.Artist, album.Price)
	}
}

func main() {
	dbConnection := controller.Connect()
	//testSelectWithParams(dbConnection)
	testSelectAll(dbConnection)
	//testInsert(dbConnection)
	controller.Disconnect()
}
