package model

import (
	"database/sql"
	"fmt"
)

type FacturaCabecera struct {
	Numero string
	Fecha  string
	Tempo  int
}

type FacturaDetalle struct {
	Codigo         int
	ProductoCodigo string
}

func MapRowsToFacturaCabecera(rows *sql.Rows) []FacturaCabecera {
	var facturas []FacturaCabecera
	for rows.Next() {
		var factura FacturaCabecera
		if err := rows.Scan(&factura.Numero, &factura.Fecha, &factura.Tempo); err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		facturas = append(facturas, factura)
	}

	return facturas
}
