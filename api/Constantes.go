package main

const (
	/*Algunas constantes usadas para respuesta en JSON. Verificar constantes.js*/
	RespuestaErrorNegocioExistente          = 0
	RespuestaErrorRegistrandoNegocio        = 1
	RespuestaNegocioRegistradoCorrectamente = 2
	RespuestaLoginNegocioNoVerificado       = 3
	RespuestaLoginCorrecto                  = 4
	RespuestaLoginError                     = 5
	RespuestaLoginIncorrecto                = 6
	// Constantes del sistema
	ImpresionNativa                      = false
	TamanioMaximoArchivoImportacionExcel = 5 << 20 // 5MB
	NombreArchivoExportadoCSV            = "ProductosExportados_SPOS3.csv"
	NombreArchivoExportadoExcel          = "ProductosExportados_SPOS3.xlsx"
	LimiteProductosMasVendidos           = 10
	DiasParaApartadosProximosAVencer     = 7
	LimiteProductosMenosVendidos         = 10
	LimiteProductosNuncaVendidos         = 10
	LimiteAutoCompletadoClientes         = 10
	MensajePieDeTicket                   = ""
	MIME                                 = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	NumeroAvisosAntesDeEliminarCuenta    = 3
	DiasParaMarcarCuentaComoInactiva     = 7
	AsuntoReactivarCuenta                = ""
	MensajeCompartir                     = ""
	UrlAcortadaCompartirSistema          = ""
	RutaDirectorioContenidoEstatico      = "/static/"
	NombreSesion                         = "SublimePosSesion"
	NombreBaseDeDatosSesiones            = "sesiones.sqlite"
	ClaveCookie                          = "905ff3f39809b9f9e204204596e234589a99a8fab5e691c08474584a5f3b7523"
	EdadDeSesionEnSegundos               = 60 * 60 * 24 // Un día en segundos
	DominioPermitidoCORS                 = "http://localhost:8080"
	PuertoServidor                       = ":2106" // Debe coincidir con constantes.js
	RutaGeneralNoNecesitaComprobacion    = "RUTA_NO_AUTH"
	ControladorBDSQLite3                 = "sqlite3"
	ControladorBDMySQL                   = "mysql"
	// Cambia aquí tu motor: (por el momento solo puede ser ControladorBDSQLite3 o ControladorBDMySQL)
	ControladorBD = ControladorBDMySQL
	/*
		Si usas MySQL recuerda darle al usuario que coloques aquí todos los permisos, y debes crear
		una base de datos como lo indique la constante "NombreBaseDeDatosNegociosMySQL", misma que debe
		ser accedida desde el usuario que coloques
	*/
	HostMySQL                            = "tcp(192.168.0.110:3307)" // Debe ir entre tcp()
	UsuarioMySQL                         = "maximo"
	PasswordMySQL                        = "manzana15"
	NombreBaseDeDatosNegociosMySQL       = "negocios_spos"
	NombreDirectorioParaSubidas          = "subidas"
	CadenaConexionBDUsuarios             = "negocio_parzibyte.me_"
	CadenaConexionBDNegocios             = "negocios_spos3_by_parzibyte.me.db"
	PrefijoBDNegocios                    = "spos"
	NombreArchivoEsquemaSQLInit          = "init.sql"
	NombreArchivoEsquemaSQLNegocios      = "esquema_negocios_sqlite.sql"
	NombreArchivoEsquemaSQLSistemas      = "esquema_spos_sqlite.sql"
	NombreArchivoEsquemaSQLInitMySQL     = "init.sql"
	NombreArchivoEsquemaSQLNegociosMySQL = "esquema_negocios.sql"
	NombreArchivoEsquemaSQLSistemasMySQL = "esquema_spos.sql"
	NombrePrimerUsuarioAdmin             = "admin"
	ClaveAPICorreo                       = ""
	PrefijoRutasAdmin                    = "/auth"
	GmailServidor                        = ""
	GmailPuerto                          = ""
	GmailCorreo                          = ""
	GmailPass                            = ""
	UrlBaseApp                           = ""
	UrlBasePaginaWeb                     = ""
	CorreoSoporteYContacto               = ""
)
