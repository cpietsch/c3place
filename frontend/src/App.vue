<template>
  <div id="app">
    <colors @color="changeColor" />
    <c3canvas :color="color" />
  </div>
</template>

<script>
import C3canvas from "./components/C3canvas.vue";
import Colors from "./components/Colors.vue";

export default {
  name: "app",
  data: function() {
    return {
      color: [0, 0, 0]
    };
  },
  methods: {
    changeColor: function(color) {
      console.log(color);
      this.color = color;
    }
  },
  components: {
    C3canvas,
    Colors
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
