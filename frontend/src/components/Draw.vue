<template>
  <div class="">
    <colors @color="changeColor" />
    <c3canvas :color="color" />
  </div>
</template>

<script>
import C3canvas from "./C3canvas.vue";
import Colors from "./Colors.vue";

export default {
  name: "draw",
  data: function() {
    return {
      color: [255, 0, 0]
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

<style scoped></style>
