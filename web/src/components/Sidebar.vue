<template>
  <ul>
    <li v-for="link in links" :key="link.name" :class="{highlighted: link.name === highlightedLink}">
      <span>{{link.displayName}}</span>
    </li>
  </ul>
</template>

<script>
export default {
  name: "sidebar",
  props: ["highlightedLink"],
  data() {
      return {
          links: [
              {
                  displayName: "Home",
                  name: "home"
              },
              {
                  displayName: "About",
                  name: "about"
              },
          ]
      }
  },
};
</script>

<style lang="scss">
ul {
  list-style: none;
}
li {
  cursor: pointer;
  span {
    font-family: "Fira Code";
    font-size: 1.2em;
    font-weight: 500;
    color: rgb(59, 59, 59);
    &:hover {
      text-decoration: underline;
    }
  }

  &:before {
    content: "►";
    margin-right: 6px;
    font-size: 0.8rem;
    vertical-align: 0.2rem;
  }
}

.highlighted {
    opacity: 0.5;
}
</style>

