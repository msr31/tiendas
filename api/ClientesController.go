package main

import (
	"log"
)

type ClientesController struct {
	AjustesUsuario AjustesDeUsuarioLogueado
}

func (clientesController *ClientesController) historial(idCliente int) *HistorialCliente {
	/*
	   Simplemente regresamos las ventas y apartados que ha hecho :)
	*/
	ac := ApartadosController{
		AjustesUsuario: AjustesDeUsuarioLogueado{
			httpRequest:        clientesController.AjustesUsuario.httpRequest,
			httpResponseWriter: clientesController.AjustesUsuario.httpResponseWriter,
		},
	}
	vc := VentasContadoController{
		AjustesUsuario: AjustesDeUsuarioLogueado{
			httpRequest:        clientesController.AjustesUsuario.httpRequest,
			httpResponseWriter: clientesController.AjustesUsuario.httpResponseWriter,
		},
	}
	return &HistorialCliente{
		Apartados:       ac.deUnCliente(idCliente),
		VentasAlContado: vc.deUnCliente(idCliente),
	}
}

func (clientesController *ClientesController) ultimoNumero() (ultimoNumero int, err error) {
	db, err := clientesController.AjustesUsuario.obtenerBaseDeDatos()

	if err != nil {
		log.Fatalf("Error abriendo base de datos: %v", err)
		panic(err)
	}

	defer db.Close()
	filas, err := db.Query("SELECT idCliente FROM clientes ORDER BY idCliente DESC LIMIT 1;")
	if err != nil {
		log.Printf("Error al consultar el último número de cliente:\n%q", err)
		return -1, err
	}

	defer filas.Close()

	if !filas.Next() {
		return 1, nil
	}
	var ultimoRowid int
	err = filas.Scan(&ultimoRowid)
	if err != nil {
		log.Printf("Error al escanear idCliente de productos:\n%q", err)
		return -1, err
	}
	return ultimoRowid, nil
}

func (clientesController *ClientesController) nuevo(cliente *Cliente) {
	ayudante := AyudanteBaseDeDatos{
		nombreTabla:    "clientes",
		AjustesUsuario: clientesController.AjustesUsuario,
	}
	ayudante.insertar(map[string]interface{}{
		"nombreCompleto": cliente.Nombre,
		"numeroTelefono": cliente.NumeroTelefono,
		"cedula":         cliente.Cedula,
		"fechanac":       cliente.Fechanac,
		"correoelec":     cliente.Correoelec,
		"residencia":     cliente.Residencia,
		"lugartrab":      cliente.LugarTrab,
		"vendedor":       cliente.Vendedor,
		"referencia":     cliente.Referencia,
		"observaciones":  cliente.Observaciones,
	})
	ultimoNumero, _ := clientesController.ultimoNumero()
	cliente.Numero = ultimoNumero
}
func (clientesController *ClientesController) actualizar(cliente *Cliente) {
	ayudante := AyudanteBaseDeDatos{
		nombreTabla:    "clientes",
		AjustesUsuario: clientesController.AjustesUsuario,
	}
	ayudante.actualizarDonde("idCliente", cliente.Numero, map[string]interface{}{
		"nombreCompleto": cliente.Nombre,
		"numeroTelefono": cliente.NumeroTelefono,
		"cedula":         cliente.Cedula,
		"fechanac":       cliente.Fechanac,
		"correoelec":     cliente.Correoelec,
		"residencia":     cliente.Residencia,
		"lugartrab":      cliente.LugarTrab,
		"vendedor":       cliente.Vendedor,
		"referencia":     cliente.Referencia,
		"observaciones":  cliente.Observaciones,
	})
}
func (clientesController *ClientesController) eliminar(cliente *Cliente) {
	if cliente.Numero == 1 {
		//Previene la eliminación del cliente mostrador
		return
	}
	ayudante := AyudanteBaseDeDatos{
		nombreTabla:    "clientes",
		AjustesUsuario: clientesController.AjustesUsuario,
	}
	ayudante.eliminarDonde("idCliente", cliente.Numero)
}
func (clientesController *ClientesController) todos(offset, limite int) ClientesConConteo {
	db, err := clientesController.AjustesUsuario.obtenerBaseDeDatos()
	clientesConteo := ClientesConConteo{}

	if err != nil {
		panic(err)
	}

	defer db.Close()
	filas, err := db.Query(`SELECT idCliente, nombreCompleto, numeroTelefono, cedula, fechanac, correoelec, residencia, lugartrab, vendedor, referencia, observaciones
	FROM clientes ORDER BY
	idCliente DESC LIMIT ? OFFSET ?`, limite, offset)
	if err != nil {
		log.Printf("Error al realizar la consulta para obtener todos los clientes:\n%q", err)
		return clientesConteo
	}

	defer filas.Close()

	clientes := []Cliente{}
	for filas.Next() {
		var cliente Cliente
		err := filas.Scan(&cliente.Numero, &cliente.Nombre, &cliente.NumeroTelefono, &cliente.Cedula, &cliente.Fechanac, &cliente.Correoelec, &cliente.Residencia, &cliente.LugarTrab, &cliente.Vendedor, &cliente.Referencia, &cliente.Observaciones)
		if err != nil {
			log.Printf("Error al escanear todos los clientes:\n%q", err)
		}
		clientes = append(clientes, cliente)
	}

	if err = filas.Err(); err != nil {
		log.Printf("Error en las filas al escanear todos los clientes:\n%q", err)
	}
	clientesConteo.Clientes = clientes
	conteo, _ := clientesController.conteo()
	clientesConteo.Conteo = conteo
	return clientesConteo
}
func (c *ClientesController) conteo() (conteo int, err error) {
	db, err := c.AjustesUsuario.obtenerBaseDeDatos()

	if err != nil {
		log.Fatalf("Error abriendo base de datos: %v", err)
		panic(err)
	}

	defer db.Close()
	filas, err := db.Query("SELECT COUNT(idCliente) FROM clientes;")
	if err != nil {
		log.Printf("Error al conteo productos:\n%q", err)
		return 0, err
	}

	defer filas.Close()

	if !filas.Next() {
		return 0, nil
	}
	var total int
	err = filas.Scan(&total)
	if err != nil {
		log.Printf("Error al conteo productos:\n%q", err)
		return 0, err
	}
	return total, nil
}
func (c *ClientesController) conteoBusqueda(nombre string) (conteo int, err error) {
	db, err := c.AjustesUsuario.obtenerBaseDeDatos()

	if err != nil {
		log.Fatalf("Error abriendo base de datos: %v", err)
		panic(err)
	}

	defer db.Close()
	filas, err := db.Query("SELECT COUNT(idCliente) FROM clientes WHERE nombreCompleto LIKE ?;",
		"%"+nombre+"%")
	if err != nil {
		log.Printf("Error al conteo clientes de búsqueda:\n%q", err)
		return 0, err
	}

	defer filas.Close()

	if !filas.Next() {
		return 0, nil
	}
	var total int
	err = filas.Scan(&total)
	if err != nil {
		log.Printf("Error al conteo clientes de búsqueda:\n%q", err)
		return 0, err
	}
	return total, nil
}
func (clientesController *ClientesController) buscar(limite, offset int, nombre string) ClientesConConteo {
	db, err := clientesController.AjustesUsuario.obtenerBaseDeDatos()
	clientesConConteo := ClientesConConteo{}

	if err != nil {
		panic(err)
	}

	defer db.Close()
	filas, err := db.Query(`SELECT idCliente, nombreCompleto, numeroTelefono, cedula, fechanac, correoelec, residencia, lugartrab, vendedor, referencia, observaciones
	FROM clientes WHERE nombreCompleto LIKE ? ORDER BY idCliente DESC LIMIT ? OFFSET ?`, "%"+nombre+"%", limite, offset)
	if err != nil {
		log.Printf("Error en consulta al realizar una búsqueda en clientes:\n%q", err)
		return clientesConConteo
	}

	defer filas.Close()

	clientes := []Cliente{}
	for filas.Next() {
		var cliente Cliente
		err := filas.Scan(&cliente.Numero, &cliente.Nombre, &cliente.NumeroTelefono, &cliente.Cedula, &cliente.Fechanac, &cliente.Correoelec, &cliente.Residencia, &cliente.LugarTrab, &cliente.Vendedor, &cliente.Referencia, &cliente.Observaciones)
		if err != nil {
			log.Printf("Error al escanear resultados de búsqueda en clientes:\n%q", err)
		}
		clientes = append(clientes, cliente)
	}

	if err = filas.Err(); err != nil {
		log.Printf("Error:\n%q", err)
	}
	conteo, _ := clientesController.conteoBusqueda(nombre)
	clientesConConteo.Conteo = conteo
	clientesConConteo.Clientes = clientes
	return clientesConConteo
}
func (clientesController *ClientesController) buscarParaAutocompletado(nombre string) []Cliente {
	db, err := clientesController.AjustesUsuario.obtenerBaseDeDatos()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	filas, err := db.Query("SELECT idCliente, nombreCompleto, numeroTelefono, cedula, fechanac, correoelec, residencia, lugartrab, vendedor, referencia, observaciones FROM clientes WHERE nombreCompleto LIKE ? ORDER BY idCliente DESC LIMIT ?;", "%"+nombre+"%", LimiteAutoCompletadoClientes)
	if err != nil {
		log.Printf("Error en consulta al realizar una búsqueda en clientes:\n%q", err)
		return nil
	}

	defer filas.Close()

	clientes := []Cliente{}
	for filas.Next() {
		var cliente Cliente
		err := filas.Scan(&cliente.Numero, &cliente.Nombre, &cliente.NumeroTelefono, &cliente.Cedula, &cliente.Fechanac, &cliente.Correoelec, &cliente.Residencia, &cliente.LugarTrab, &cliente.Vendedor, &cliente.Referencia, &cliente.Observaciones)
		if err != nil {
			log.Printf("Error al escanear resultados de búsqueda en clientes:\n%q", err)
		}
		clientes = append(clientes, cliente)
	}

	if err = filas.Err(); err != nil {
		log.Printf("Error:\n%q", err)
	}
	return clientes
}

func (clientesController *ClientesController) porRowid(idCliente int) *Cliente {
	db, err := clientesController.AjustesUsuario.obtenerBaseDeDatos()

	if err != nil {
		panic(err)
	}
	defer db.Close()
	filas, err := db.Query("SELECT idCliente, nombreCompleto, numeroTelefono, cedula, fechanac, correoelec, residencia, lugartrab, vendedor, referencia, observaciones FROM clientes WHERE idCliente = ? LIMIT 1;", idCliente)
	if err != nil {
		log.Printf("Error al consultar cliente por código:\n%q", err)
		return nil
	}

	defer filas.Close()

	if !filas.Next() {
		return nil
	}
	var cliente Cliente
	err = filas.Scan(&cliente.Numero, &cliente.Nombre, &cliente.NumeroTelefono, &cliente.Cedula, &cliente.Fechanac, &cliente.Correoelec, &cliente.Residencia, &cliente.LugarTrab, &cliente.Vendedor, &cliente.Referencia, &cliente.Observaciones)
	if err != nil {
		log.Printf("Error al escanear cliente por idCliente:\n%q", err)
		return nil
	}
	return &cliente
}
