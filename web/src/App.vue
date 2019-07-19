<template>
  <div id="app">
    <div class="left"></div>
    <div class="center">
      <header>
        <input type="text" placeholder="Type here to search" v-model="query" @input="search" />
      </header>

      <div class="content">
        <transition appear name="card">
          <Card v-if="showCard" class="card" :card="card" :key="card.title" />
        </transition>

        <transition name="hint">
          <Hint v-if="!showCard" class="hint" />
        </transition>
      </div>
    </div>
    <div class="right"></div>
  </div>
</template>

<script>
import Card from "./components/Card.vue";
import Hint from "./components/Hint.vue";
import axios from "axios";
import { debounce, isEmpty } from "lodash";

export default {
  name: "app",
  components: { Card, Hint },
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
@font-face {
  font-family: Roboto;
  src: url("../assets/fonts/roboto.ttf");
}

@font-face {
  font-family: "Fira Code";
  src: url("../assets/fonts/firaCode.woff2");
}

html,
body,
#app {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
  background: #f8f6f6;
}

#app {
  display: flex;
  flex-direction: row;
  align-items: stretch;

  .left,
  .right {
    flex: 4;
  }

  .center {
    flex: 9;
  }
}

.center {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  .content {
    flex: 1;
  }
}

.center > header {
  display: flex;
  flex-direction: row;
  background: white;
  border-radius: 0 0 8px 8px;
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

.content {
  position: relative;
  .card,
  .hint {
    position: absolute;
  }
  .card {
    top: 0;
  }
  .hint {
    transform: rotate(-6deg);
    text-align: center;
    width: 100%;
    top: 20%;
    font-size: 1vw;
  }
}

.content {
  .card {
    z-index: 2;
  }
  .hint {
    z-index: 1;
  }
}

.card {
  &-enter {
    opacity: 0;
    transform: translateY(-20px);
    &-to {
      opacity: 1;
      transform: translateY(0);
    }
    &-active {
      transition: all 0.1s ease-in-out;
    }
  }
  &-leave {
    opacity: 1;
    transform: translateY(0);
    &-to {
      opacity: 0;
      transform: translateY(30px);
    }
    &-active {
      transition: all 0.1s ease-in-out;
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
      transition: all 0.1s ease-in-out;
    }
  }
  &-leave {
    opacity: 1;
    &-to {
      opacity: 0;
    }
    &-active {
      transition: all 0.1s ease-in-out;
    }
  }
}
</style>
