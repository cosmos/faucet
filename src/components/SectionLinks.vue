<template lang="pug">
.section
  .section-header
    h1 Community Links
  .section-main
    list-item(
      href="https://riot.im/app/#/room/#cosmos:matrix.org"
      target="_blank"
      title="Community Chat"
      subtitle="#cosmos")
    list-item(
      href="https://riot.im/app/#/room/#cosmos_validators:matrix.org"
      target="_blank"
      title="Validator Chat"
      subtitle="#cosmos_validators")
    list-item(
      href="https://explorecosmos.network"
      target="_blank"
      title="Testnet Explorer"
      subtitle="explorecosmos.network")
    list-item(
      href="https://cosmos.network"
      target="_blank"
      title="Cosmos Website"
      subtitle="cosmos.network")
</template>

<script>
import ListItem from "../components/NiListItem.vue";
export default {
  name: "faucet",
  components: {
    ListItem
  },
  data: () => ({}),
  methods: {}
};
</script>
