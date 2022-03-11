package controller

import (
	"database/sql"
	"fmt"
)

func Get(db *sql.DB, query string, params []interface{}) (rows *sql.Rows) {
	rows, err := db.Query(query, params...)
	if err != nil {
		return nil
	}
	return rows
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
	result, err := db.Exec(query, params...)
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
	result, err := db.Exec(query, params...)
	if err != nil {
		return 0, fmt.Errorf("error: error while updating a record: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error: error while getting rowsAffected: %v", err)
	}
	return rowsAffected, nil
}
