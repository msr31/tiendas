<template>
  <v-dialog v-model="mostrar" persistent max-width="500">
    <v-card>
      <v-card-title class="headline">Correo electrónico existente</v-card-title>
      <v-card-text>
        <p>Parece que ya hay un negocio registrado con ese correo electrónico. Intenta de nuevo con otra dirección.
          Si continúas teniendo problemas, contáctame en <a href="https://parzibyte.me/blog">https://parzibyte.me/blog</a>
        </p>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="success" @click.native="cerrarDialogo()">Cerrar</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
  import {CorreoSoporteYContacto} from "../../constantes";

  export default {
    name: "DialogoNegocioExistente",
    props: ["mostrar"],
    methods: {
      cerrarDialogo() {
        this.$emit("cerrar-dialogo");
      },
      data: () => ({
        correo: CorreoSoporteYContacto,
      })
    }
  }
</script>

<style scoped>

</style>
