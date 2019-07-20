<template>
  <div id="app">
    <title v-text="title" />

    <Sidebar />

    <div class="header-container">
      <Header v-model="query" @input="search" />
    </div>

    <div class="content" id="content">
      <transition
        appear
        name="card"
        @after-leave="updateContentHeight"
        @enter="updateContentHeight"
      >
        <Card v-if="showCard" :card="card" :key="card.title" />
      </transition>
    </div>

    <transition appear name="hint">
      <Hint v-if="!showCard" />
    </transition>

    <transition appear name="next">
      <div v-if="showCard" class="next">►</div>
    </transition>

    <transition name="footer">
      <Footer v-if="showCard" />
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
    }, 150),
    updateContentHeight() {
      let heights = ["12em"];

      const card = document.getElementsByClassName("card")[0];

      if (card !== undefined) {
        heights.push(`${card.clientHeight}px`);
      }

      const heightStr = `calc(${heights.join(" + ")})`;
      console.log(heightStr);

      const content = document.getElementById("content");
      content.style.height = heightStr;
    }
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

  // Necessary for the footer
  // to always be on the bottom
  // of the page, even when the
  // page is shorter than the viewport
  position: relative;
}

html,
body,
#app {
  // Necessary for the footer
  // to always be on the bottom
  // of the page, even when the
  // page is shorter than the viewport
  min-height: 100vh;
  scroll-behavior: smooth;
  // Makes sure that nothing
  // overflows the viewport
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

.content {
  position: relative;
  .card {
    position: absolute;
    top: 4em;
  }
}

.card,
.hint {
  width: 100%;
  box-sizing: border-box;
}

.hint {
  position: fixed;
  top: 24vh;
  transform: rotate(-12deg);
  font-size: 3vw;
  user-select: none;

  @media screen and (min-width: 1000px) {
    font-size: 1em;
  }
  p {
    padding: 0;
  }
}

.next {
  $height: 100px;

  @mixin shadow($opacity, $blur) {
    box-shadow: 1px 2px $blur rgba(0, 191, 255, $opacity);
  }

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
  transition: filter 0.1s ease-in-out, box-shadow 0.1s ease-in-out;

  cursor: pointer;

  background: linear-gradient(
    to top left,
    rgb(0, 162, 255) 0%,
    deepskyblue 100%
  );

  @include shadow(0.6, 12px);
  &:hover {
    filter: saturate(2);
    @include shadow(0.9, 18px);
  }

  left: $center-width + (100% - $center-width) * 1/ 2;
  @media screen and (max-width: $center-min-width * (1 / 0.46)) {
    left: calc(#{$center-min-width} + (100% - #{$center-min-width}) / 2);
  }
  @media screen and (orientation: portrait) {
    display: none;
  }
}

.header-container {
  z-index: 4;
}

.footer {
  z-index: 3;
}

.content {
  z-index: 2;
}

.hint {
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

.card {
  &-enter {
    opacity: 0.2;
    transform: translateY(-20px);
    &-to {
      transform: none;
      opacity: 1;
    }

    &-active {
      transition: all 0.1s ease-out;
    }
  }
  &-leave {
    transform: none;
    opacity: 1;

    &-to {
      transform: translateY(50px);
      opacity: 0.2;
    }

    &-active {
      transition: all 0.1s ease-in;
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
      transition: all 0.2s ease-out;
    }
  }
  &-leave {
    opacity: 1;
    transform: none;

    &-to {
      transform: translateY(50px) scaleX(0.5);
      opacity: 0.2;
    }

    &-active {
      transition: all 0.1s ease-in;
    }
  }
}

.hint {
  &-enter {
    opacity: 0;

    &-to {
      opacity: 1;
    }

    &-active {
      transition: all 0.3s ease-out;
      transition-delay: .1s;
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
      transition-delay: .2s;
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
