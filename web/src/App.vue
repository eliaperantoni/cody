<template>
  <div id="app">
    <title v-text="title" />

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

    <transition appear name="next">
      <div v-if="showCard" class="next">►</div>
    </transition>

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
      window.scrollTo(0, 0);
    }, 150)
  },
  computed: {
    showCard() {
      return !isEmpty(this.card);
    },
    title() {
      return this.showCard ? this.card.title : "Cody";
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
$center-min-width: 540px;
$center-width: 46%;

%center {
  min-width: $center-min-width;
  width: $center-width;
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
body,
#app {
  min-height: 100vh;
  overflow-x: hidden;
  scroll-behavior: smooth;
}

.header-container {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  display: flex;
  flex-direction: row;
  justify-content: center;
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
  position: relative;
  margin: 0 auto;
  padding: 4em 0 6em 0;
}

.card,
.hint {
  width: 100%;
  box-sizing: border-box;
}

.content .hint {
  transform: rotate(-12deg);
  margin-top: 10vh;
  p {
    padding: 0;
  }
}

.next {
  $height: 100px;

  border-radius: 0 14px 14px 0;
  height: $height;
  width: 50px;
  position: fixed;
  top: 100px;
  color: white;
  font-size: 1.3em;
  line-height: $height; // Makes arrow vertically centered
  text-align: center;
  user-select: none;
  transform-origin: top left;

  cursor: pointer;

  background: linear-gradient(
    to top left,
    rgb(0, 162, 255) 0%,
    deepskyblue 100%
  );
  box-shadow: 1px 2px 12px rgba(0, 191, 255, 0.6);

  left: $center-width + (100% - $center-width) * 1/ 2;
  @media screen and (max-width: $center-min-width * (1 / 0.46)) {
    left: calc($center-min-width + (100% - $center-min-width) / 2);
  }
  @media screen and (orientation: portrait) {
    display: none;
  }
}

.header-container {
  z-index: 2;
}

.content {
  z-index: 1;
}

.next {
  z-index: 0;
}

.footer {
  position: absolute;
  bottom: 24px;
  left: 0;
  right: 0;
  text-align: center;
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

.next {
  &-enter {
    opacity: 0.4;
    transform: scaleX(0.5);
    &-to {
      transform: none;
      opacity: 1;
    }
    &-active {
      transition: all .2s ease-out;
    }
  }
}

.footer {
  &-enter {
    opacity: 0;
    transform: translateY(10px);
    &-to {
      transform: none;
      opacity: 1;
    }
    &-active {
      transition: all 0.1s ease-in-out;
    }
  }
  &-leave {
    transform: none;
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
