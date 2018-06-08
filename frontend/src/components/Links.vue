<template lang="pug">
#links
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
    title="Network Explorer"
    subtitle="explorecosmos.network")
  list-item(
    href="https://cosmos.network"
    target="_blank"
    title="Cosmos Website"
    subtitle="cosmos.network")
  #bottom
    | &copy; 2018 Interchain Foundation
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

<style lang="stylus">
@import '~variables'

#links
  display block
  padding 0 1rem

#bottom
  color var(--bc)
  font-size 0.75rem
  padding 1rem 0

@media screen and (min-width: 375px)
  #links
    padding 0 2rem
</style>
