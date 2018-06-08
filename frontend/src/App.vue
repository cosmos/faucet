<template lang="pug">
#app
  img(src="~assets/brandmark.png")
  router-view
  notifications(:notifications='notifications' theme='cosmos')
  #bottom &copy; 2018 Interchain Foundation
</template>


<script>
import { mapGetters } from "vuex";
import Notifications from "@nylira/vue-notifications";
export default {
  name: "app",
  components: {
    Notifications
  },
  computed: {
    ...mapGetters(["notifications"])
  },
  mounted() {}
};
</script>

<style lang="stylus" src="./styles/app.styl"></style>

<style lang="stylus">
@import '~variables'

img
  position absolute
  top -20rem
  left -20rem
  z-index 1

#bottom
  color var(--bc)
  font-size 0.75rem
  padding 1rem 0
  text-align: center;
</style>
