<template lang="pug">
#faucet
  .section
    faucet-header
    form(v-on:submit.prevent='onSubmit', method='post')
      form-group(:error='$v.fields.captcha.$error'
        field-id='faucet-captcha' field-label='Captcha')
        vue-recaptcha#faucet-captcha(
          ref="recaptcha"
          @verify="onVerify"
          @expired="onExpired"
          :sitekey="sitekey")
        form-msg(name='Captcha' type='required' v-if='!$v.fields.captcha.required')
      form-group(:error='$v.fields.address.$error'
        field-id='faucet-address' field-label='Send To')
        field#faucet-address(
          type='text'
          v-model='fields.address'
          placeholder='Testnet address'
          size="lg")
        form-msg(name='Address' type='required' v-if='!$v.fields.address.required')
        form-msg(name='Address' type='bech32' :body="bech32error" v-else-if='!$v.fields.address.bech32Validate')
      form-group
        btn(v-if='sending' value='Sending...' disabled color="primary" size="lg")
        btn(v-else @click='onSubmit' value="Send me tokens" color="primary" size="lg" icon="send")
  section-join
  section-links
</template>

<script>
import axios from "axios";
import VueRecaptcha from "vue-recaptcha";
import { required } from "vuelidate/lib/validators";
import b32 from "../scripts/b32";
import Btn from "@nylira/vue-button";
import Field from "@nylira/vue-field";
import FormGroup from "../components/NiFormGroup";
import FormMsg from "../components/NiFormMsg";
import FaucetHeader from "../components/FaucetHeader";
import SectionJoin from "../components/SectionJoin.vue";
import SectionLinks from "../components/SectionLinks.vue";
export default {
  name: "faucet",
  components: {
    Btn,
    Field,
    FormGroup,
    FaucetHeader,
    FormMsg,
    SectionJoin,
    SectionLinks,
    VueRecaptcha
  },
  data: () => ({
    fields: {
      captcha: "",
      address: ""
    },
    sending: false,
    sitekey: "6LdqyV0UAAAAAEqgBxvSsDpL2aeTEgkz_VTz1Vi1"
  }),
  methods: {
    resetForm() {
      this.fields.address = "";
      this.fields.captcha = "";
      this.$refs.recaptcha.reset(); // Direct call reset method
      this.$v.$reset();
    },
    onVerify() {
      this.fields.captcha = "verified";
    },
    onExpired() {
      this.$refs.recaptcha.reset(); // Direct call reset method
    },
    async onSubmit() {
      this.$v.$touch();
      if (this.$v.$error) return;

      this.sending = true;
      let address = this.fields.address;
      axios
        .post("/claim", {
          address: address
        })
        .then(() => {
          this.sending = false;
          this.$store.commit("notify", {
            title: "Successfully Sent",
            body: `Sent tokens to ${address}`
          });
          this.resetForm();
        })
        .catch(err => {
          this.sending = false;
          this.$store.commit("notifyError", {
            title: "Error Sending",
            body: `An error occurred while trying to send: "${err.message}"`
          });
        });
    },
    bech32Validate(param) {
      try {
        b32.decode(param);
        this.bech32error = null;
        return true;
      } catch (error) {
        this.bech32error = error.message;
        return false;
      }
    }
  },
  validations() {
    return {
      fields: {
        address: {
          required,
          bech32Validate: this.bech32Validate
        },
        captcha: { required }
      }
    };
  }
};
</script>

<style lang="stylus">
@import '~variables'

#faucet
  max-width 40rem
  width 100%
  margin 0 auto

.section
  margin 0.5rem
  padding 1rem
  background var(--app-bg)
  position relative
  z-index 10
  label
    display none

  input:-webkit-autofill
    -webkit-text-fill-color var(--txt) !important
    -webkit-box-shadow 0 0 0px 3rem var(--app-fg) inset

  .section-main
    padding 0 1rem

@media screen and (min-width: 375px)
  .section
    padding 2rem 1rem

@media screen and (min-width: 768px)
  .section
    padding 3rem 2rem
    margin 1rem
</style>
