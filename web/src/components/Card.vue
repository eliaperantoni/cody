<template>
  <div class="card">
    <div class="wrapper">
      <header>
        <nav>
          <div v-for="(bread, i) in card.breads" :key="bread">
            <span>{{bread}}</span>
            <span v-if="i < card.breads.length - 1">&nbsp;&gt;&nbsp;</span>
          </div>
        </nav>

        <h1>{{card.title}}</h1>
      </header>

      <div class="markdown" v-html="renderedContent" />
    </div>
  </div>
</template>

<script>
import "../../assets/monokai-sublime.css";
import MarkdownIt from "markdown-it";

const md = new MarkdownIt({
  html: true,
  linkify: true
}).use(require('markdown-it-highlightjs'), {});

export default {
  name: "card",
  props: ["card"],
  computed: {
    renderedContent() {
      return md.render(this.card.content);
    }
  }
};
</script>

<style lang="scss" scoped>
.card {
  background: white;
  border-radius: 13px;
  box-shadow: 0 1px 37px rgba(0, 177, 255, 0.09);
  overflow-x: hidden;
  .wrapper {
    box-sizing: border-box;
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding: 2.6em;
  }
}

h1,
p {
  overflow: hidden;
  text-overflow: ellipsis;
}

header {
  > nav,
  > h1 {
    text-shadow: 0 0.4px 1px rgba(0, 0, 0, 0.1);
  }
}

nav {
  font-family: "Fira Code";
  color: #8e8e8e;
  margin-bottom: 0.3em;
  display: flex;
  flex-direction: row;
  font-size: 1.1em;

  > div > span:first-child:hover {
    text-decoration: underline;
    cursor: pointer;
  }
}

h1 {
  font-family: Roboto;
  font-size: 2.2em;
  font-weight: 400;
  color: #3d3d3d;
  margin: 0;
  margin-left: -0.04em;
}

.markdown::v-deep {
  width: 100%;
  word-wrap: break-word;
  $font-base-color: #3d3d3d;
  font-family: "Fira Code";
  color: $font-base-color;
  font-size: 1.2em;
  p code {
    background: rgb(198, 234, 255);
    color: rgb(2, 155, 216);
    border: 1px solid rgb(105, 195, 255);
    border-radius: 4px;
    padding: 1px 4px 1px 4px;
  }
  pre code {
    width: 100%;
  }
  * {
    &.danger,
    &.warning,
    &.info,
    &.success {
      padding: 18px;
      color: darken($font-base-color, 10);
    }
    &.danger {
      $base: rgba(255, 0, 98, 0.4);
      background: $base;
      border-left: 8px solid saturate($base, 40);
    }
    &.warning {
      $base: rgba(255, 208, 0, 0.4);
      background: $base;
      border-left: 8px solid saturate($base, 40);
    }
    &.info {
      $base: rgba(0, 183, 255, 0.4);
      background: $base;
      border-left: 8px solid saturate($base, 40);
    }
    &.success {
      $base: rgba(0, 255, 0, 0.4);
      background: $base;
      border-left: 8px solid saturate($base, 40);
    }
  }
}
</style>
