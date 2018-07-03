<template lang="pug">
#app
  tm-cookie-consent
  .top-bar
    .left
      .key testnet:
      .value {{ config.chain }}
    .right
      a(href="https://explorecosmos.network" target="_blank") explorer #[i.material-icons assessment]
  img(src="~assets/brandmark.png" width="512" height="512")
  router-view
  notifications(:notifications='notifications' theme='cosmos')
  #bottom &copy; 2018 Interchain Foundation
</template>


<script>
import { mapGetters } from "vuex";
import { TmCookieConsent } from "@tendermint/ui";
import Notifications from "@nylira/vue-notifications";

export default {
  name: "app",
  components: {
    Notifications,
    TmCookieConsent
  },
  computed: {
    ...mapGetters(["notifications", "config"])
  },
  mounted() {}
};
</script>

<style lang="stylus" src="./styles/app.styl"></style>

<style lang="stylus">
@import '~variables'

img
  position absolute
  top -7rem
  left -7rem
  z-index 1

.top-bar
  position fixed
  top 0
  right 0
  background var(--accent)
  text-align center
  height 3rem
  width 100%
  z-index 100

  display flex
  justify-content space-between
  line-height 3rem

  .left
    display flex
    padding 0 1rem
  .key
    text-transform: lowercase
    margin-right 0.5rem
  .value
    font-weight bold
    color var(--bright)
  .right
    a
      display block
      background var(--mc)
      padding 0 1rem
      color var(--bright)
      display flex
      align-items center
      i
        margin-left 0.5rem
#bottom
  color var(--bc)
  font-size 0.75rem
  padding 1rem 0
  text-align: center;
</style>
