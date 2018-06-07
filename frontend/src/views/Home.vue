<template>
  <div class="home">
    <h1>Cosmos Testnet Faucet</h1>
      <form v-on:submit.prevent="onSubmit" method="post">
        <label>Address</label>
        <input type="text" v-model="address" required>
        <br>
        <input type="submit" value="Claim">
      </form>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "home",
  components: {},
  data: () => ({
    address: ""
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
