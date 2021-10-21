<template>
  <v-flex xs12>
    <v-toolbar color="light-blue" dark>
      <v-toolbar-title v-show="!estaBuscando">Clientes</v-toolbar-title>
      <v-toolbar-title v-show="estaBuscando"></v-toolbar-title>
      <v-spacer></v-spacer>
      <v-flex xs12 v-show="estaBuscando">
        <v-text-field
          v-model="busqueda"
          ref="inputBusqueda"
          hide-details
          label="Escriba parte del nombre"
          solo-inverted
          single-line
        ></v-text-field>
      </v-flex>
      <v-btn @click="prepararParaBuscar()" v-show="!estaBuscando" icon>
        <v-icon>search</v-icon>
      </v-btn>
      <v-btn @click="cancelarBusqueda()" v-show="estaBuscando" icon>
        <v-icon>close</v-icon>
      </v-btn>
    </v-toolbar>
    <v-flex xs12>
      <Publicidad></Publicidad>
    </v-flex>
    <v-card>
      <div class="text-xs-center pt-2">
        <v-pagination
          v-model="paginacion.pagina"
          :length="numeroDePaginas"
        ></v-pagination>
      </div>
      <v-data-table
        :loading="cargando"
        :headers="encabezados"
        :items="clientes"
        hide-actions
        class="elevation-1"
      >
        <template slot="items" slot-scope="props">
          <td>{{ props.item.Nombre }}</td>
          <td>{{ props.item.NumeroTelefono }}</td>
          <td>{{ props.item.Cedula }}</td>
          <td>{{ props.item.Fechanac }}</td>
          <td>{{ props.item.Correoelec }}</td>
          <td>{{ props.item.Residencia }}</td>
          <td>{{ props.item.LugarTrab }}</td>
          <td>{{ props.item.Vendedor }}</td>
          <td>{{ props.item.Referencia }}</td>
          <td>{{ props.item.Observaciones }}</td>
          <td class="justify-center layout px-0">
            <v-btn
              target="_blank"
              :href="'tel:' + props.item.NumeroTelefono"
              icon
              ripple
            >
              <v-icon color="blue">phone</v-icon>
            </v-btn>
            <v-btn @click="mostrarHistorial(props.item.Numero)" icon ripple>
              <v-icon color="blue">info</v-icon>
            </v-btn>
            <v-btn @click="editar(props.item)" icon ripple>
              <v-icon color="orange">edit</v-icon>
            </v-btn>
            <v-btn icon ripple @click="eliminar(props.item)">
              <v-icon color="red">delete</v-icon>
            </v-btn>
          </td>
        </template>
        <template slot="no-data">
          <v-alert v-show="!busqueda" :value="true" color="info">
            <div class="text-xs-center">
              <h1>No hay clientes</h1>
              <v-icon class="icono-grande">announcement</v-icon>
              <p>
                No has registrado ningún cliente. Agrega uno con el botón
                <v-icon>add</v-icon>
                de la esquina
              </p>
            </div>
          </v-alert>
          <v-alert v-show="busqueda" :value="true" color="info" dark>
            <div class="text-xs-center">
              <h1>Sin resultados</h1>
              <v-icon class="icono-grande">highlight_off</v-icon>
              <p>
                No hay clientes que coincidan con "
                <strong>{{ busqueda }}</strong
                >"
              </p>
            </div>
          </v-alert>
        </template>
      </v-data-table>
    </v-card>
  </v-flex>
</template>
<script>
import { HTTP_AUTH } from "../../http-common";
import Publicidad from "../Publicidad";

const CLIENTES_POR_PAGINA = 7;
export default {
  components: { Publicidad },
  beforeMount() {
    this.obtener();
  },
  computed: {
    numeroDePaginas() {
      if (this.paginacion.conteoClientes == null) return 0;
      return Math.ceil(this.paginacion.conteoClientes / this.paginacion.limite);
    },
  },
  watch: {
    busqueda() {
      this.obtener();
    },
    paginacion: {
      handler() {
        this.obtener();
      },
      deep: true,
    },
  },
  methods: {
    mostrarHistorial(numeroCliente) {
      this.$emit("mostrar-historial", numeroCliente);
    },
    editar(cliente) {
      this.$emit("editar", cliente);
    },
    eliminar(cliente) {
      this.$emit("eliminar", cliente);
    },
    prepararParaBuscar() {
      this.estaBuscando = true;
      this.$nextTick(this.$refs.inputBusqueda.focus);
    },
    cancelarBusqueda() {
      this.busqueda = "";
      this.estaBuscando = false;
    },
    obtener() {
      let verdaderoOffset =
        this.paginacion.pagina > 1
          ? (this.paginacion.pagina - 1) * this.paginacion.limite
          : 0;
      if (this.busqueda) {
        if (this.deberiaReiniciarPaginacionAlBuscar) {
          this.paginacion.pagina = 1;
          verdaderoOffset = 0;
          this.deberiaReiniciarPaginacionAlBuscar = false;
        }
        HTTP_AUTH.get(
          `buscar/clientes/${encodeURIComponent(
            this.busqueda
          )}/${verdaderoOffset}/${this.paginacion.limite}`
        ).then((resultadosDeBusqueda) => {
          this.clientes = resultadosDeBusqueda.Clientes;
          this.paginacion.conteoClientes = resultadosDeBusqueda.Conteo;
          this.cargando = false;
        });
      } else {
        HTTP_AUTH.get(
          `clientes/${verdaderoOffset}/${this.paginacion.limite}`
        ).then((clientesConConteo) => {
          this.clientes = clientesConConteo.Clientes;
          this.paginacion.conteoClientes = clientesConConteo.Conteo;
          this.cargando = false;
        });
      }
    },
  },
  data: () => ({
    deberiaReiniciarPaginacionAlBuscar: false,
    paginacion: {
      offset: 0,
      limite: CLIENTES_POR_PAGINA,
      conteoClientes: 0,
      pagina: 1,
    },
    busqueda: "",
    estaBuscando: false,
    cargando: false,
    clientes: [],
    encabezados: [
      {
        text: "Nombre",
        align: "left",
        value: "Nombre",
      },
      {
        text: "Teléfono",
        align: "left",
        value: "NumeroTelefono",
      },
      {
        text: "Cédula",
        value: "cedula",
      },
      {
        text: "Fecha nacimiento",
        value: "Fechanac",
      },
      {
        text: "Correo electrónico",
        value: "Correoelec",
      },
      {
        text: "Residencia",
        value: "Residencia",
      },
      {
        text: "Lugar de trabajo",
        value: "LugarTrab",
      },
      {
        text: "Vendedor",
        value: "Vendedor",
      },
      {
        text: "Referencia",
        value: "Referencia",
      },
      {
        text: "Observaciones",
        value: "Observaciones",
      },
      {
        text: "Opciones",
        sortable: false,
      },
    ],
  }),
};
</script>
