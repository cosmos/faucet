<template lang="pug">
#app
  router-view
  notifications(:notifications='notifications' theme='cosmos')
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
