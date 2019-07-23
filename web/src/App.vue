<template>
  <div id="app">
    <title v-text="title" />

    <Sidebar />

    <div class="header-container">
      <Header v-model="query" @input="search" />
    </div>

    <div class="content" id="content">
      <transition appear @enter="animeCardEnter" @leave="animeCardLeave">
        <Card v-if="showCard" :card="card" :key="card.title" />
      </transition>
    </div>

    <transition appear @enter="animeArrowEnter" @leave="animeArrowLeave">
      <div v-if="showArrowLeft" @click="arrowLeft" class="arrow left">◄</div>
    </transition>

    <transition appear @enter="animeArrowEnter" @leave="animeArrowLeave">
      <div v-if="showArrowRight" @click="arrowRight" class="arrow right">►</div>
    </transition>

    <transition appear name="hint">
      <Hint v-if="!showCard" />
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
import anime from "animejs";
import { debounce } from "lodash";

export default {
  name: "app",
  components: { Card, Hint, Sidebar, Footer, Header },
  data() {
    return {
      query: "",
      cards: [],
      cardIndex: 0,

      anime: "contextSwitch"
    };
  },
  computed: {
    card() {
      return this.cards[this.cardIndex];
    },
    showCard() {
      return this.card != undefined;
    },
    title() {
      return this.showCard ? this.card.title : "Cody";
    },
    showArrowRight() {
      return this.showCard && this.cardIndex != this.cards.length - 1;
    },
    showArrowLeft() {
      return this.showCard && this.cardIndex != 0;
    }
  },
  created() {
    if (window.injected != undefined) {
      this.cards = [window.injected];
    }

    window.onpopstate = event => {
      Object.assign(this, event.state);
    };
    this.updateHistory();
  },
  methods: {
    search: debounce(async function() {
      this.anime = "contextSwitch";

      if (this.query === "") {
        this.cards = [];
      } else {
        let res = await axios.get(
          `http://localhost:5000/search?q=${this.query}`
        );
        this.cards = res.data;
      }

      this.scrollToTop();
      this.cardIndex = 0;

      this.updateHistory();
    }, 100),
    // This function makes it possible to have cards positioned
    // absolutely inside their container so that we can have
    // cool animations. The problem with absolutely positioned
    // elements is that their size is not taken into account
    // when positioning other elements. This function forces
    // the container div to take the height of its content.
    updateContentHeight(card) {
      let heights = ["12em"];

      if (card !== null) {
        heights.push(`${card.clientHeight}px`);
      }

      const heightStr = `calc(${heights.join(" + ")})`;

      const content = document.getElementById("content");
      content.style.height = heightStr;
    },
    updateHistory() {
      const data = this.$data;
      if (this.card == null) {
        window.history.pushState(data, `Index`, `/`);
      } else {
        window.history.pushState(
          data,
          `Card ${this.card.id}`,
          `/card/${this.card.id}`
        );
      }
    },
    arrowRight() {
      this.anime = "right";
      this.cardIndex++;
      this.scrollToTop();
      this.updateHistory();
    },
    arrowLeft() {
      this.anime = "left";
      this.cardIndex--;
      this.scrollToTop();
      this.updateHistory();
    },
    scrollToTop() {
      window.scrollTo(0, 0);
    },
    animeCardEnter(el, done) {
      let anim;
      switch (this.anime) {
        case "contextSwitch":
          anim = {
            translateY: [-20, 0],
            opacity: [0.4, 1]
          };
          break;
        case "right":
          anim = {
            translateX: [20, 0],
            opacity: [0.7, 1]
          };
          break;
        case "left":
          anim = {
            translateX: [-20, 0],
            opacity: [0.7, 1]
          };
          break;
      }
      anime({
        targets: el,
        duration: 100,
        easing: "easeOutQuad",
        complete: () => {
          this.updateContentHeight(el);
          done();
        },
        ...anim
      });
      done();
    },
    animeCardLeave(el, done) {
      if (this.anime != "contextSwitch") {
        done();
        return;
      }
      anime({
        targets: el,
        opacity: [1, 0],
        duration: 100,
        easing: "easeInQuad",
        translateY: [0, 40],
        complete: () => {
          if (!this.showCard) this.updateContentHeight(null);
          done();
        }
      });
    },
    animeArrowEnter(el, done) {
      anime({
        targets: el,
        opacity: [0.4, 1],
        duration: 100,
        easing: "easeOutQuad",
        complete: done,
        scaleX: [0.5, 1]
      });
    },
    animeArrowLeave(el, done) {
      let anim;
      switch (this.anime) {
        case "contextSwitch":
          anim = {
            translateY: [0, 40],
            scaleX: [1, 0.5]
          };
          break;
        case "right":
        case "left":
          anim = {
            scaleX: [1, 0.5]
          };
          break;
      }
      anime({
        targets: el,
        opacity: [1, 0],
        duration: 100,
        easing: "easeInQuad",
        complete: done,
        ...anim
      });
    }
  }
};
</script>

<style lang="scss">
$center-min-width: 540px;
$center-width: 46%;
$side-width: (100% - $center-width) / 2;
$max-width-before-fixed: $center-min-width * 100 / ($center-width / 1%);

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
}

body {
  overflow-y: scroll;
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

  @media screen and (max-width: 1300px) {
    display: none;
  }
}

.content {
  @extend %center;
  margin: 0 auto;
  position: relative;
  .card {
    position: absolute;
    top: 4em;
    width: 100%;
    box-sizing: border-box;
  }
}

.hint {
  position: fixed;
  top: 24vh;
  left: 0;
  right: 0;
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

.arrow {
  $height: 100px;
  $width: 50px;

  @mixin shadow($opacity, $blur) {
    box-shadow: 1px 2px $blur rgba(0, 191, 255, $opacity);
  }

  height: $height;
  width: $width;
  position: fixed;
  top: 100px;
  color: white;
  font-size: 1.3em;
  line-height: $height; // Makes arrow vertically centered
  text-align: center;
  user-select: none;
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

  @media screen and (orientation: portrait) {
    display: none;
  }

  &.right {
    border-radius: 0 14px 14px 0;
    transform-origin: top left;
    left: $side-width + $center-width;
    @media screen and (max-width: $max-width-before-fixed) {
      left: calc((100% - #{$center-min-width}) / 2 + #{$center-min-width});
    }
  }

  &.left {
    border-radius: 14px 0 0 14px;
    transform-origin: top right;
    left: calc(#{$side-width} - #{$width});
    @media screen and (max-width: $max-width-before-fixed) {
      left: calc((100% - #{$center-min-width}) / 2 - #{$width});
    }
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

.arrow {
  z-index: 0;
}

.footer {
  position: absolute;
  bottom: 24px;
  left: 0;
  right: 0;
  text-align: center;
  text-overflow: ellipsis;
  overflow: hidden;
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

.hint {
  &-enter {
    opacity: 0;

    &-to {
      opacity: 1;
    }

    &-active {
      transition: all 0.3s ease-out;
      transition-delay: 0.1s;
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
      transition-delay: 0.2s;
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
