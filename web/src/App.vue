<template>
  <div id="app">
    <div class="left"></div>
    <div class="center">
      <header>
        <input type="text" placeholder="Type here to search" v-model="query">
      </header>
      <Card :card="card"/>
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
        breads: [
          "Programming Languages",
          "Go",
          "Slices"
        ]
      }
    }
  },
  methods: {
    search: debounce(function() {
      return
      if (this.query == "") {
        this.card = {};
        return;
      }
      // TODO Do search
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

<style lang="sass">
html, body, #app
  height: 100%
  width: 100%
  margin: 0
  padding: 0
  background: #F8F6F6
  display: flex
  flex-direction: row
  align-items: stretch
  .left, .right
    flex: 4
  .center
    flex: 9
    display: flex
    flex-direction: column
    align-items: stretch

.center > header
  display: flex
  flex-direction: row
  background: white
  border-radius: 0 0 13px 13px
  box-shadow: 0 1px 37px rgba(0, 177, 255, 0.16)
  margin-bottom: 18px
  input
    flex: 1
    font-size: 1.1em
    font-family: Roboto
    font-weight: 200
    background: transparent
    border: none
    outline-width: 0
    padding: .8em 1.9em
    color: #979797
    &::placeholder
      opacity: .5

Card
  flex: 1
</style>
