<template lang="pug">
.section-header
  h1 Cosmos Testnet Faucet
  p Hello intrepid spaceperson! Use this faucet to get tokens for the #[span.chain {{ config.chain }}] testnet. Please don't abuse this service&mdash;the number of available tokens is limited.
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "faucet-header",
  computed: {
    ...mapGetters(["config"])
  }
};
</script>

<style lang="stylus">
@import '~variables'

.section-header
  padding 0 1rem 1.5rem
  span.chain
    white-space nowrap

  strong
    font-weight 500

  h1
    font-size 1.5rem
    font-weight 600
    color var(--bright)
    line-height 1.25

  h1 + p
    margin-top 0.75rem
  p
    color var(--dim)
</style>
