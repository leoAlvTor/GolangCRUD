package controller

import (
	"database/sql"
	"facturacion/model"
	"fmt"
)

func GetAll(db *sql.DB, query string, params []interface{}) ([]model.Album, error) {

	var albums []model.Album
	fmt.Println(params)
	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("error while executing statement: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var album model.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("error while reading rows, error %v", err)
		}
		albums = append(albums, album)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return albums, nil
}
