<template lang="pug">
#faucet
  #form
    header
      h1 Cosmos Testnet Faucet
      h2 Hello intrepid explorer! You can use this form to get tokens for the #[strong `gaia-6000`] testnet. Don't have a testnet address yet? #[a(href="https://cosmos.network/testnet") Join the testnet!]
    form(v-on:submit.prevent='onSubmit', method='post')
      .li-form
        label(for='faucet-captcha') Captcha
        vue-recaptcha#faucet-captcha(:sitekey='recaptcha')
      .li-form
        label(for='faucet-address') Testnet Address
        field#faucet-address(type='text', v-model='address', required='', size='lg' placeholder='Your testnet address')
      .li-form
        btn(type='submit', value='Send me tokens', size='lg', color='primary')
  links
  #bottom
    | &copy; 2018 Interchain Foundation
</template>

<script>
import axios from "axios";
import VueRecaptcha from "vue-recaptcha";
import Btn from "@nylira/vue-button";
import Field from "@nylira/vue-field";
import Links from "../components/Links";
export default {
  name: "faucet",
  components: {
    Btn,
    Field,
    Links,
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

#faucet
  max-width 32rem
  width 100%
  margin 0 auto

#form
  padding 0 1rem
  background var(--app-fg)
  margin 0 0 3rem

  header
    padding 1rem 0

    strong
      font-weight 500

  h1
    font-size 1.5rem
    font-weight 500
    color bright

  h2
    color var(--dim)
    font-size sm

  label
    display none

  .li-form
    padding 1rem 0
    border-top 1px solid var(--bc)

  .ni-btn
    width 100%

#bottom
  color var(--bc)
  font-size 0.75rem
  padding 1rem 0
</style>
