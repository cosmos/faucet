<template lang="pug">
header#form-header
  h1 Cosmos Testnet Faucet
  p Hello intrepid explorer! You can use this form to get tokens for the #[strong `gaia-6000`] testnet. Don't have a testnet address yet? #[a(href="https://cosmos.network/testnet") Join the testnet!]
</template>

<style lang="stylus">
@import '~variables'

#form-header
  padding 0 1rem 1.5rem

  strong
    font-weight 500

  h1
    font-size 1.5rem
    font-weight 600
    color bright
    margin-bottom 0.5rem

  p
    color var(--dim)
    font-size 0.875rem
</style>
