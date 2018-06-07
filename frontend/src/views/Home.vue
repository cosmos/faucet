<template lang="pug">
#page
  #form
    header
      h1 Cosmos Testnet Faucet
      h2 testnet version: #[a(href="https://explorecosmos.network" target="_blank") gaia-6000]
    form(v-on:submit.prevent='onSubmit', method='post')
      .li-form
        label(for='faucet-captcha') Captcha
        vue-recaptcha#faucet-captcha(:sitekey='recaptcha')
      .li-form
        label(for='faucet-address') Testnet Address
        field#faucet-address(type='text', v-model='address', required='', size='lg' placeholder='Your testnet address')
      .li-form
        btn(type='submit', value='Claim tokens', size='lg', color='primary')
  #links
    a(href="https://cosmos.network" target="_blank")
      .key Cosmos Website
      .value cosmos.network
    a(href="https://explorecosmos.network" target="_blank")
      .key Network Explorer
      .value explorecosmos.network
    a(href="https://riot.im/app/#/room/#cosmos:matrix.org" target="_blank")
      .key Community Chat
      .value #cosmos
    a(href="https://riot.im/app/#/room/#cosmos_validators:matrix.org" target="_blank")
      .key Validator Chat
      .value #cosmos_validators
  #bottom
    | &copy; 2018 Interchain Foundation
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

#page
  max-width 32rem
  width 100%
  margin 0 auto

#form
  padding 0 1rem
  background var(--app-fg)
  margin 0 0 3rem

  header
    padding 1rem 0

  h1
    font-size 1.5rem
    font-weight 500
    color bright

  h2
    color var(--dim)

  label
    display none

  .li-form
    padding 1rem 0
    border-top 1px solid var(--bc)

  .ni-btn
    width 100%

#links
  a
    display block
    line-height 2rem
    border-top 1px solid var(--bc-dim)
    display flex
    justify-content space-between
    padding 0 1rem
    margin-bottom 0.25rem
    .key
      color var(--dim)
    .value
      color var(--link)
    &:hover
      background var(--hover-bg)
      .key
        color var(--bright)
      .value
        color var(--hover)
#bottom
  color var(--bc)
  font-size 0.75rem
  padding 2rem 1rem 1rem
</style>
