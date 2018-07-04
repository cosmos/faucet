<template lang="pug">
.section
  .section-header
    h1 Don't have a testnet address?
    p There's two ways to get one. The first is by using Voyager, the crypto wallet for Cosmos. If you know command-line-fu, you can also generate an address with the Cosmos SDK.
  .section-main
    btn(type="anchor" href="https://cosmos.network/testnet" value="Join the latest testnet" size="lg" target="_blank" icon="wifi_tethering")
</template>

<script>
import Btn from "@nylira/vue-button";
export default {
  name: "section-join",
  components: {
    Btn
  }
};
</script>
