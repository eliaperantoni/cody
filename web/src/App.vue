<template>
  <div id="app">
    <Sidebar />

    <div class="header-container">
      <Header v-model="query" @input="search" />
    </div>

    <div class="content">
      <transition appear name="content" mode="out-in">
        <Card v-if="showCard" :card="card" :key="card.title" />
        <Hint v-else />
      </transition>
    </div>

    <transition name="footer">
      <Footer v-if="query != ''" />
    </transition>
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Card from "./components/Card.vue";
import Hint from "./components/Hint.vue";
import Sidebar from "./components/Sidebar.vue";
import Footer from "./components/Footer.vue";
import axios from "axios";
import { debounce, isEmpty } from "lodash";

export default {
  name: "app",
  components: { Card, Hint, Sidebar, Footer, Header },
  data() {
    return {
      query: "",
      card: {
        title: `Delete Slice Element by Index ${Math.random()}`,
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
        title: `Delete Slice Element by Index ${Math.random()}`,
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
%center {
  min-width: 540px;
  width: 46%;
  @media screen and (orientation: portrait) {
    width: 100%;
    min-width: auto;
  }
}

@font-face {
  font-family: Roboto;
  src: url("../assets/fonts/roboto.ttf");
}

@font-face {
  font-family: "Fira Code";
  src: url("../assets/fonts/firaCode.woff2");
}

body {
  background: #f8f6f6;
  margin: 0;
  position: relative;
}

html,
body {
  min-height: 100vh;
}

#app {
  padding: 4em 0 6em 0;
  overflow-x: hidden;
}

.header-container {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  display: flex;
  flex-direction: row;
  justify-content: center;
  z-index: 10;
  .header {
    @extend %center;
  }
}

.sidebar {
  position: fixed;
  top: 60px;
  left: 80px;

  @media screen and (orientation: portrait) {
    display: none;
  }
}

.content {
  @extend %center;
  margin: 0 auto;
}

.card,
.hint {
  width: 100%;
  box-sizing: border-box;
}

.footer {
  position: absolute;
  bottom: 24px;
  left: 0;
  right: 0;
  text-align: center;
}

.content .hint {
  transform: rotate(-12deg);
  margin-top: 10vh;
}

.content {
  &-enter {
    opacity: 0.7;
    &-to {
      opacity: 1;
    }
    &-active {
      transition: all 0.1s ease-in-out;
    }
  }
}

.footer {
  &-enter {
    opacity: 0;
    transform: translateY(10px);
    &-to {
      transform: translateY(0);
      opacity: 1;
    }
    &-active {
      transition: all 0.1s ease-in-out;
    }
  }
  &-leave {
    transform: translateY(0);
    opacity: 1;
    &-to {
      transform: translateY(10px);
      opacity: 0;
    }
    &-active {
      transition: all 0.1s ease-in-out;
    }
  }
}
</style>
