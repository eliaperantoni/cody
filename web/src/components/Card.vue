<template>
  <article class="card">
    <header>
      <nav>
        <div v-for="(bread, i) in card.breads" :key="bread">
          <span>{{bread}}</span>
          <span v-if="i < card.breads.length - 1">&nbsp;&gt;&nbsp;</span>
        </div>
      </nav>

      <h1>{{card.title}}</h1>
    </header>
  </article>
</template>

<script>
export default {
  name: "card",
  props: ["card"]
};
</script>

<style lang="scss" scoped>
article {
  background: white;
  border-radius: 13px;
  box-shadow: 0 1px 37px rgba(0, 177, 255, 0.09);
  padding: 2.6em;

  header {
    > nav,
    > h1 {
      text-shadow: 0 0.4px 1px rgba(0, 0, 0, 0.1);
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
  }
}
</style>
