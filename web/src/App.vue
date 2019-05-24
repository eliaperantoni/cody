<template>
  <div id="app">
    <title>{{card.id || "Cody"}}</title>
    <input type="text" placeholder="Type here to search" v-model="query" @input="search">
    <Card :card="card" v-if="showCard"></Card>
  </div>
</template>

<script>
import Card from "./components/Card.vue";
import axios from "axios";
import { debounce, isEmpty } from "lodash";

export default {
  name: "app",
  components: { Card },
  data() {
    return {
      query: "",
      card: {}
    };
  },
  methods: {
    search: debounce(function() {
      if (this.query == "") {
        this.card = {};
        return
      }
      axios.get("https://randomuser.me/api/").then(response => {
        let seed = response.data.info.seed
        this.card = {
          id: seed
        };
        window.history.replaceState(
          null,
          seed,
          seed
        )
      });
    }, 150)
  },
  computed: {
    showCard() {
      return !isEmpty(this.card);
    }
  },
  mounted() {
    if (window.injected != undefined) {
      this.card = window.injected;
    }
  }
};
</script>

<style lang="scss">
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
