<template>
  <div class="hint">
      <p class="title">Cody</p>
      <p class="content">Cody</p>
  </div>
</template>

<script>
export default {
  name: "hint"
};
</script>

<style lang="scss" scoped>
.hint {
  font-size: 1vw;
  text-align: center;
}

p {
  margin: 0;
  font-family: "Fira Code";
  color: #c7c7c7;
  &.title {
    font-size: 4em;
    font-weight: bolder;
  }
  &.content {
    font-size: 2em;
  }
}
</style>
