<template>
  <div class="sidebar">
    <header>
      <img
        src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Google_%22G%22_Logo.svg/1024px-Google_%22G%22_Logo.svg.png"
        alt="cody logo"
      />
      <span>Cody</span>
    </header>
    <ul>
      <li
        v-for="link in links"
        :key="link.name"
        :class="{highlighted: link.name === highlightedLink}"
      >
        <span>{{link.displayName}}</span>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "sidebar",
  props: ["highlightedLink"],
  data() {
    return {
      links: [/*
        {
          displayName: "Home",
          name: "home"
        },
        {
          displayName: "Sauce",
          name: "source"
        },
        {
          displayName: "About",
          name: "about"
        }
      */]
    };
  }
};
</script>

<style lang="scss" scoped>
header {
    display: flex;
    flex-direction: row;
    align-items: center;
    img {
        height: 40px;
        width: auto;
    }
    span {
        font-family: "Fira Code";
        color: rgb(59, 59, 59);
        font-size: 1.8em;
        margin-left: 10px;
    }
}

ul {
  list-style: none;
  margin: 18px 0 0 32px;
  padding: 0;
}
li {
  cursor: pointer;
  span {
    font-family: "Fira Code";
    font-size: 1.2em;
    font-weight: 500;
    &:hover {
      text-decoration: underline;
    }
  }

  &:before {
    content: "►";
    margin-right: 6px;
    font-size: 0.8em;
    vertical-align: 0.2em;
  }

  span, &:before {
    color: rgb(92, 92, 92);
  }
}

.highlighted {
  opacity: 0.5;
}
</style>

