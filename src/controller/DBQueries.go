package controller

import (
	"database/sql"
	"facturacion/model"
	"fmt"
)

func Get(db *sql.DB, query string, params []interface{}) ([]interface{}, error) {
	var objects []interface{}
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
		objects = append(objects, album)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}
	return objects, nil
}

func Add(db *sql.DB, query string, params []interface{}) (int64, error) {
	result, err := db.Exec(query, params...)
	if err != nil {
		return 0, fmt.Errorf("error: error while adding a new record: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error: error while getting lastInsertId: %v", err)
	}
	return id, nil
}

func Delete(db *sql.DB, query string, params []interface{}) (int64, error) {
	result, err := db.Exec(query, params)
	if err != nil {
		return 0, fmt.Errorf("error: error while deleting a record: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error: error while getting rowsAffected: %v", err)
	}
	return rowsAffected, nil
}

func Update(db *sql.DB, query string, params []interface{}) (int64, error) {
	result, err := db.Exec(query, params)
	if err != nil {
		return 0, fmt.Errorf("error: error while updating a record: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error: error while getting rowsAffected: %v", err)
	}
	return rowsAffected, nil
}
