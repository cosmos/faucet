<template lang="pug">
#home
  h1 Cosmos Testnet Faucet
  form(v-on:submit.prevent='onSubmit', method='post')
    .li-form
      label(for='faucet-captcha') Captcha
      vue-recaptcha#faucet-captcha(:sitekey='recaptcha')
    .li-form
      label(for='faucet-address') Testnet Address
      field#faucet-address(type='text', v-model='address', required='', size='lg' placeholder='Your testnet address')
    .li-form
      btn(type='submit', value='Claim tokens', size='lg', color='primary')
</template>

<script>
import axios from "axios";
import VueRecaptcha from "vue-recaptcha";
import Btn from "@nylira/vue-button";
import Field from "@nylira/vue-field";
export default {
  name: "home",
  components: {
    Btn,
    Field,
    VueRecaptcha
  },
  data: () => ({
    address: "",
    recaptcha: "6LdqyV0UAAAAAEqgBxvSsDpL2aeTEgkz_VTz1Vi1"
  }),
  methods: {
    onSubmit() {
      axios
        .post("/claim", {
          address: this.address
        })
        .then(() => {
          this.$store.commit("notify", {
            title: "Claim Succcessful",
            body: "Refresh your wallet to get your tokens"
          });
        })
        .catch(() => {
          this.$store.commit("notifyError", {
            title: "Claim Error",
            body: "There was an error claiming tokens"
          });
        });
    }
  }
};
</script>

<style lang="stylus">
@import '~variables'

#home
  max-width 32rem
  width 100%
  padding 0 1rem
  margin 0 auto
  background var(--app-fg)

  h1
    font-size 1.25rem
    font-weight 500
    padding 1rem 0

  label
    display none

  .li-form
    padding 1rem 0
    border-top 1px solid var(--bc)

  .ni-btn
    width 100%
</style>
