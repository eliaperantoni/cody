<template>
  <div class="hint">
      <p class="title">Cody</p>
      <p class="content">{{randomMessage}}</p>
  </div>
</template>

<script>
import {sample} from "lodash";

export default {
  name: "hint",
  data() {
    return {
      randomMessage: "",
      randomMessages: [
        "Beep Boop!",
        "import antigravity",
        "let foo, bar",
        "Hello World!",
      ]
    }
  },
  created() {
    this.randomMessage = sample(this.randomMessages);
  }
};
</script>

<style lang="scss" scoped>
.hint {
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
}

p {
  margin: 0;
  font-family: "Fira Code";
  color: #c7c7c7;
  text-align: center;
  &.title {
    font-size: 4em;
    font-weight: bolder;
  }
  &.content {
    font-size: 1.5em;
    max-width: 30ch;
    overflow-wrap: break-word;
  }
}
</style>
