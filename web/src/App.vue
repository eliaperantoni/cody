<template>
  <div id="app">
    <div class="left"></div>
    <div class="center">
      <header>
        <input type="text" placeholder="Type here to search" v-model="query" @input="search"/>
      </header>
      <Card v-if="showCard" :card="card" />
    </div>
    <div class="right"></div>
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
      card: {
        title: "Delete Slice Element by Index",
        breads: ["Programming Languages", "Go", "Slices"]
      }
    };
  },
  methods: {
    search: debounce(function() {
      if (this.query === "") {
        this.card = {};
        return;
      }
      this.card = {
        title: "Delete Slice Element by Index",
        breads: ["Programming Languages", "Go", "Slices"]
      };
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
@font-face {
  font-family: Roboto;
  src: url("../assets/fonts/roboto.ttf");
}

@font-face {
  font-family: "Fira Code";
  src: url("../assets/fonts/firaCode.woff2");
}

html, body, #app {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
  background: #F8F6F6;
  display: flex;
  flex-direction: row;
  align-items: stretch;

  .left, .right {
    flex: 4;
  }

  Card {
    flex: 1;
  }
}

.center {
  flex: 9;
  display: flex;
  flex-direction: column;
  align-items: stretch;

  > header {
    display: flex;
    flex-direction: row;
    background: white;
    border-radius: 0 0 13px 13px;
    box-shadow: 0 1px 37px rgba(0, 177, 255, 0.16);
    margin-bottom: 18px;

    input {
      flex: 1;
      font-size: 1.1em;
      font-family: Roboto;
      font-weight: 200;
      background: transparent;
      border: none;
      outline-width: 0;
      padding: 0.8em 1.9em;
      color: #979797;

      &::placeholder {
        opacity: 0.5;
      }
    }
  }
}
</style>
