<template>
  <v-dialog v-model="mostrar" persistent max-width="500">
    <v-card>
      <v-card-title class="headline">Registrar cliente</v-card-title>
      <v-card-text>
        <v-form ref="formulario">
          <v-container fluid grid-list-md>
            <v-layout row wrap>
              <v-flex xs12>
                <v-text-field ref="nombreCliente" prepend-icon="face" label="Nombre completo" type="text"
                              v-model="cliente.nombre" :rules="reglas.nombre" hint="Nombre y apellidos"
                              required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field :mask="mascara" prepend-icon="phone" label="Número de teléfono" type="text"
                              v-model="cliente.numeroTelefono" :rules="reglas.numeroTelefono"
                              hint="Número telefónico de 8 dígitos"
                              required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field :mask="mascara2" prepend-icon="face" label="Cédula" type="text"
                              v-model="cliente.cedula" :rules="reglas.cedula"
                              hint="Cédula del cliente"
                              required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field :mask="mascara3" prepend-icon="face" label="Fecha de Nacimiento" type="text"
                              v-model="cliente.fechanac" :rules="reglas.fechanac"
                              hint="Fecia de Nacimiento del cliente"
                              required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field ref="correoelec" prepend-icon="face" label="Correo Electronico" type="text"
                              v-model="cliente.correoelec" :rules="reglas.correoelec"
                              hint="Correo Electronico del cliente"
                              required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field ref="residencia" prepend-icon="face" label="Residencia" type="text"
                              v-model="cliente.residencia" :rules="reglas.residencia"
                              hint="Residencia del cliente"
                              required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field ref="lugartrab" prepend-icon="face" label="Lugar de Trabajo" type="text"
                              v-model="cliente.lugartrab" :rules="reglas.lugartrab"
                              hint="Lugar de trabajo del cliente"
                              required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field ref="vendedor" prepend-icon="face" label="Vendedor" type="text"
                              v-model="cliente.vendedor" :rules="reglas.vendedor"
                              hint="Nombre del vendedor"
                              required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field ref="referencias" prepend-icon="face" label="Referencia" type="text"
                              v-model="cliente.referencia" :rules="reglas.referencia"
                              hint="Referencias del cliente"
                              required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field ref="observaciones" prepend-icon="face" label="Observaciones" type="text"
                              v-model="cliente.observaciones" :rules="reglas.observaciones"
                              hint="Observaciones del cliente"
                              required></v-text-field>
              </v-flex>
            </v-layout>
          </v-container>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn :loading="cargando" color="green darken-1" flat="flat" @click.native="guardar()">Guardar</v-btn>
        <v-btn color="gray" flat="flat" @click.native="cerrarDialogo()">Cerrar</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
<script>
  import {HTTP_AUTH} from "../../http-common";

  export default {
    watch: {
      mostrar(estaMostrado) {
        if (estaMostrado) this.enfocarInputPrincipal();
      }
    },
    methods: {
      enfocarInputPrincipal() {
        this.$nextTick(this.$refs.nombreCliente.focus);
      },
      resetearFormulario() {
        this.$refs.formulario.reset();
      },
      cerrarDialogo() {
        this.resetearFormulario();
        this.$emit("cerrar-dialogo");
      },
      guardar() {
        if (this.$refs.formulario.validate()) {
          let cliente = Object.assign({}, this.cliente);
          this.cargando = true;
          HTTP_AUTH.post("cliente", cliente).then(clienteGuardado => {
            this.cargando = false;
            if (clienteGuardado.Numero) {
              this.resetearFormulario();
              this.$emit("cliente-guardado", clienteGuardado);
              this.enfocarInputPrincipal();
            } else {
              /**Handle error here */
            }
          });
        }
      }
    },
    props: ["mostrar"],
    data: () => ({
      mascara: "####-####",
      mascara2: "#-####-####",
      mascara3: "##-##-####",
      cargando: false,
      cliente: {},
      reglas: {
        nombre: [
          nombre => {
            if (!nombre) return "Introduzca el nombre completo del cliente";
            if (!/^[a-zA-Z\sáéíóúÁÉÍÓÚ]*$/.test(nombre))
              return "El nombre sólo puede llevar letras";
            return true;
          }
        ],
        numeroTelefono: [
          numeroTelefono => {
            if (!numeroTelefono) return "Introduzca el número de teléfono";
            if (numeroTelefono.length !== 8)
              return "El número debe tener 8 dígitos";
            return true;
          }
        ],
        cedula: [
          cedula => {
            if (!cedula) return "Introduzca el número de cédula";
            if (cedula.length !== 9)
              return "El número debe tener 9 dígitos";
            return true;
          }
        ],
        fechanac: [
          fechanac => {
            if (!fechanac) return "Introduzca la fecha de nacimiento del cliente";
            if (fechanac.length !== 8)
              return "El número debe tener 8 dígitos";
            return true;
          }
        ],
        correoelec: [
          correoelec => {
            if (!correoelec) return "Introduzca correo electronico del cliente";
            if (!/^[a-zA-Z\sáéíóúÁÉÍÓÚ/@/.]*$/.test(correoelec))
              return "El Correo Electronico debe estar correcto";
            return true;
          }
        ],
        residencia: [
          residencia => {
            if (!residencia) return "Introduzca la fecha de nacimiento del cliente";
            if (!/^[a-zA-Z\sáéíóúÁÉÍÓÚ/0-9]*$/.test(residencia))
              return "La residencia debe estar correcta";
            return true;
          }
        ],
        lugartrab: [
          lugartrab => {
            if (!lugartrab) return "Introduzca la fecha de nacimiento del cliente";
            if (!/^[a-zA-Z\sáéíóúÁÉÍÓÚ/0-9]*$/.test(lugartrab))
              return "Escriba bien el lugar de trabajo";
            return true;
          }
        ],
        vendedor: [
          vendedor => {
            if (!vendedor) return "Introduzca la fecha de nacimiento del cliente";
            if (!/^[a-zA-Z\sáéíóúÁÉÍÓÚ]*$/.test(vendedor))
              return "Escriba bien el nombre del vendedor";
            return true;
          }
        ],
        referencia: [
          referencia => {
            if (!referencia) return "Introduzca la referencia del cliente";
            if (!/^[a-zA-Z\sáéíóúÁÉÍÓÚ/0-9/-]*$/.test(referencia))
              return "La fecha debe estar correcta";
            return true;
          }
        ],
        observacion: [
          observacion => {
            if (!observacion) return "Introduzca alguna observación del cliente";
            if (!/^[a-zA-Z\sáéíóúÁÉÍÓÚ/0-9/.]*$/.test(observacion))
              return "Hay un caracter no aceptado en la descripción";
            return true;
          }
        ]
      }
    })
  };
</script>
