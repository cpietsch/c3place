<template>
  <div id="app">
    <c3canvas />
  </div>
</template>

<script>
import C3canvas from "./components/C3canvas.vue";

export default {
  name: "app",
  components: {
    C3canvas
  }
};
</script>

<style>
body,
html {
  margin: 0;
  padding: 0;
}
#app {
}
</style>
