<template>
  <v-layout row wrap>
    <v-flex xs12>

      <h1 class="display-1">{{
        procesoTerminado ?
        "Cuenta verificada correctamente":
        "Estamos verificando tu cuenta"}}</h1>
      <h2 class="subheading">{{procesoTerminado ? "" : paso}}</h2>
      <v-progress-linear :color="procesoTerminado ? 'success' : 'info'" height="20"
                         :value="procesoTerminado ? 100 : valor"></v-progress-linear>
      <p v-show="!procesoTerminado">Por favor, espera un momento.</p>

    </v-flex>
    <v-flex v-show="procesoTerminado" xs12 class="text-xs-center">
      <p class="headline">Todo en orden, ahora puedes iniciar sesión. Tus credenciales <strong>fueron
        enviadas a tu correo
        electrónico</strong> cuando te registraste</p>
      <v-btn @click="irAlLogin()" large color="success">Iniciar sesión</v-btn>
      <h1 class="display-1">Corre la voz</h1>
      <p class="subheading">Hemos creado una aplicaci&oacute;n de punto de venta, lo que nos motiva.
        <br>
        <strong>Por favor, comparte</strong></p>
      
    
    </v-flex>
  </v-layout>
</template>

<script>
  import {EventBus} from "../main";
  import {HTTP} from "../http-common";
  import {EnlaceCompartir, MensajeCompartir} from "../constantes";

  export default {
    name: "VerificarNegocio",
    mounted() {
      EventBus.$emit("ocultarMenu");
      EventBus.$emit("ponerTitulo", "Iniciar sesión");
      this.idIntervalo = setInterval(() => this.siguientePaso(), 1000);
      this.verificarToken(this.$route.params.token);
    },
    data: () => ({
      enlaces: {
       
      },
      idIntervalo: null,
      valor: 0,
      procesoTerminado: false,
      numeroPaso: -1,
      pasos: ["Creando base de datos", "Preparando primer uso", "Registrando primer usuario", "Dando permisos al usuario", "A punto de terminar..."],
    }),
    computed: {
      paso() {
        if (this.numeroPaso < 0) return "";
        if (this.numeroPaso < this.pasos.length) {
          return this.pasos[this.numeroPaso];
        }
        return "";
      }
    },
    methods: {
      irAlLogin() {
        this.$router.push({name: "Login"});
      },
      verificarToken(token) {
        HTTP.get(`negocio/verificar/${token}`)
          .then(respuesta => {
            console.log("Al verificar con el servidor:", respuesta);
          }).finally(() => {
          this.procesoTerminado = true;
        });
      },
      siguientePaso() {
        if (this.numeroPaso < this.pasos.length - 1) {
          this.numeroPaso++;
          this.valor += (100 / (this.pasos.length + 1));
        } else {
          clearInterval(this.idIntervalo);
        }
      }
    },
  }
</script>

<style scoped>

</style>
