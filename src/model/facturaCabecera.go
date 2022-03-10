package model

type facturaCabecera struct {
	numero string
	fecha  string
	tempo  int
}

func GetFacturas() []facturaCabecera {
	var facturas []facturaCabecera
	f1 := facturaCabecera{numero: "001", fecha: "01/01/2020", tempo: 1}
	facturas = append(facturas, f1)
	facturas = append(facturas, f1)
	return facturas
}
