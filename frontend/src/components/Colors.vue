<template>
  <div class="colors">
    <div v-for="(color, i) in colors" :key="i" @click="$emit('color', color)" :style="style(color)" />
  </div>
</template>

<script>
const palette = [
  [0, 0, 0], // blasck
  [255, 255, 255], // white
  [228, 228, 228], // light grey
  [136, 136, 136], // grey
  [34, 34, 34], //black

  [255, 167, 209], // pink
  [229, 0, 9], // red
  [229, 149, 0], // orange

  [160, 106, 66], // brown

  [229, 217, 0], // yellow
  [148, 224, 68], // light green
  [2, 190, 1], // green
  [0, 211, 221], // cyan

  [0, 131, 199], // medium blue
  [0, 0, 234], // dark blue
  [207, 110, 228], // light purple
  [130, 0, 128] // dark Purple
];

export default {
  computed: {
    colors: () => palette
  },
  methods: {
    style([r, g, b]) {
      return {
        background: `rgb(${r}, ${g}, ${b})`
      };
    }
  }
};
</script>

<style scoped>
.colors {
  width: 40px;
}
.colors > div {
  width: 40px;
  height: 40px;
  height: 40px;
}
</style>
