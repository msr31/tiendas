import {DIA_EN_MILISEGUNDOS} from "./constantes";

export const FUNCIONES = {
  /**
   * Devuelve la fecha que se le pasa, más la hora a las 12 de la noche
   *
   * Por ejemplo, si se le pasa algo como 2018-06-20 devuelve 2018-06-20T00:00:00
   *
   * */
  agregarHoraCeroAFecha(fecha) {
    return `${fecha}T00:00:00`
  },
  /**
   * Devuelve la fecha de hoy en el formato:
   * YYYY-MM-DD
   *
   * Por ejemplo:
   *
   * 2018-06-20
   * */
  hoyComoCadena() {
    let fecha = new Date(Date.now() + 864e5),
      anio = fecha.getFullYear(),
      mes = this.agregarCerosALaIzquierdaSiEsNecesario(fecha.getMonth() + 1),
      dia = this.agregarCerosALaIzquierdaSiEsNecesario(fecha.getDate());
    return `${anio}-${mes}-${dia}`;
  },
  /**
   * Lo único que hace es convertir la fecha a algo como
   * 2018-05-24
   *
   * Por cierto, getMonth() devuelve el mes pero tomando en cuenta que enero es 0 y
   * diciembre es 11
   */
  formatearFecha(fecha) {
    if (!fecha instanceof Date)
      throw new TypeError("La fecha debe ser un objeto de tipo Date");

    let mes = fecha.getMonth() + 1, diaDelMes = fecha.getDate();
    return `${fecha.getFullYear()}-${this.agregarCerosALaIzquierdaSiEsNecesario(mes)}-${this.agregarCerosALaIzquierdaSiEsNecesario(diaDelMes)}`;
  },
  componerFechaParaFin(fecha) {
    fecha = new Date(fecha);
    fecha.setTime(fecha.getTime() + fecha.getTimezoneOffset() * 60 * 1000); //Necesario si viene de cadena. Ver https://parzibyte.me/blog/2018/03/05/ajustando-fechas-javascript/
    fecha.setTime(fecha.getTime() + DIA_EN_MILISEGUNDOS);
    return this.formatearFecha(fecha);
  },
  componerFechaParaInicio(fecha) {
    return fecha;
  },
  esteAnioComoCadena() {
    return new Date().getFullYear().toString();
  },
  esteMesComoCadena() {
    let mes = new Date().getMonth() + 1;

    return this.agregarCerosALaIzquierdaSiEsNecesario(mes);
  },
  agregarCerosALaIzquierdaSiEsNecesario(valor) {
    valor = valor.toString();
    return valor.length < 2 ? `0${valor}` : valor.toString();
  },
  colorHexadecimalAleatorio() {
    let colores = ["#f2476a", "#fb654e", "#eb2d3a", "#add8e6", "#90ee90", "#ffcb7e", "#ff9464", "#d5e389", "#d0b191", "#c18390", "#F44336", "#9C27B0", "#673AB7", "#3F51B5", "##2196F3", "#4CAF50", "#8BC34A", "#CDDC39", "##FFEB3B", "#FF9800", "#FF5722"];
    return colores[Math.floor(Math.random() * colores.length)];
  },
  /*
  * @deprecated
  * */
  _colorHexadecimalAleatorio() {
    return "#000000".replace(/0/g, () => (~~(Math.random() * 16)).toString(16));
  }
};
