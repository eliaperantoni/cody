<template>
  <div class="header">
    <input
      type="text"
      placeholder="How do I ..."
      :value="value"
      @input="$emit('input', $event.target.value)"
    />
  </div>
</template>

<script>
export default {
  props: ["value"],
  name: "header"
};
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  flex-direction: row;
  background: white;
  border-radius: 0 0 8px 8px;
  box-shadow: 0 1px 37px rgba(0, 177, 255, 0.16);

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
</style>

