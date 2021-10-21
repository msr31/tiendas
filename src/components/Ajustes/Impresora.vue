<template>
  <v-layout row wrap>
    <v-flex xs12>
      <v-alert type="info" :value="true">
        Recuerde seleccionar únicamente impresoras térmicas. Si no cuenta con una, no toque nada
      </v-alert>
      <v-select
        :loading="cargandoImpresoras"
        :items="impresoras"
        v-model="impresoraSelecionada"
        label="Seleccione su impresora"
        single-line
      ></v-select>
      <v-btn :loading="probandoImpresora" @click="probarCon(impresoraSelecionada)" v-show="impresoraSelecionada"
             color="info">Imprimir ticket de
        prueba
      </v-btn>
    </v-flex>
    <v-snackbar :timeout="5000" :right="true" :bottom="true" v-model="snackbars.impresoraCorrecta">
      Correcto. La impresora funciona perfectamente
      <v-btn flat color="pink" @click.native="snackbars.impresoraCorrecta = false">OK</v-btn>
    </v-snackbar>
    <v-snackbar :timeout="0" :right="true" :bottom="true" v-model="snackbars.impresoraIncorrecta">
      Hubo un error al probar la impresora. Asegúrese de que está conectada
      <v-btn flat color="pink" @click.native="snackbars.impresoraIncorrecta = false">OK</v-btn>
    </v-snackbar>
    <v-snackbar :timeout="2000" :right="true" :bottom="true" v-model="snackbars.impresoraGuardada">
      Impresora guardada
      <v-btn flat color="pink" @click.native="snackbars.impresoraGuardada = false">OK</v-btn>
    </v-snackbar>
    <v-dialog v-model="mostrarDialogo" persistent max-width="500">
      <v-card>
        <v-card-title class="headline">¿Guardar impresora?</v-card-title>
        <v-card-text>
          <p>Se ha determinado que la impresora funciona correctamente. <br>¿Desea guardar esta configuración de
            impresora?</p>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn :loading="guardandoImpresora" color="green darken-1" flat="flat" @click.native="guardarImpresora()">
            Sí
          </v-btn>
          <v-btn color="gray" flat="flat" @click.native="mostrarDialogo = false">No</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-layout>
</template>

<script>
  import {HTTP_AUTH} from "../../http-common";

  export default {
    name: "Impresora",
    beforeMount() {
      this.obtener();
    },
    methods: {
      guardarImpresora() {
        this.guardandoImpresora = true;
        HTTP_AUTH.put("nombre/impresora", this.impresoraSelecionada).then(resultados => {
          this.guardandoImpresora = false;
          this.mostrarDialogo = false;
          if (resultados) {
            this.snackbars.impresoraGuardada = true;
          }
        });
      },
      probarCon(nombreImpresora) {
        this.probandoImpresora = true;
        HTTP_AUTH.get(`probar/impresora/${encodeURIComponent(nombreImpresora)}`)
          .then(resultados => {
            this.probandoImpresora = false;
            if (resultados) {
              this.snackbars.impresoraCorrecta = true;
              this.mostrarDialogo = true;
            } else {
              this.snackbars.impresoraIncorrecta = true;
            }
          })
      },
      obtener() {
        this.cargandoImpresoras = true;
        HTTP_AUTH
          .get("impresoras").then(impresoras => {
          this.impresoras = impresoras;
          this.cargandoImpresoras = false;
        })
          .then(() => {
            HTTP_AUTH.get("nombre/impresora").then(nombreImpresora => {
              if (nombreImpresora) this.impresoraSelecionada = nombreImpresora;
            })
          })
      }
    },
    data: () => ({
      guardandoImpresora: false,
      probandoImpresora: false,
      mostrarDialogo: false,
      snackbars: {
        impresoraCorrecta: false,
        impresoraIncorrecta: false,
        impresoraGuardada: false,
      },
      cargandoImpresoras: false,
      impresoras: [],
      impresoraSelecionada: "",
    })
  }
</script>
