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
#app {
}
</style>
