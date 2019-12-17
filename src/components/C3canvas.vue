<template>
  <div class="canvas">
    <canvas ref="canvas" @click="onClick"></canvas>
  </div>
</template>

<script>
import { loadImage } from "../utils.js";
const url = "http://localhost:4000/";

export default {
  computed: {
    canvas: function() {
      return this.$refs.canvas;
    },
    context: function() {
      return this.canvas.getContext("2d");
    }
  },
  methods: {
    onClick: function(el) {
      const payload = {
        x: el.layerX,
        y: el.layerY,
        r: 255,
        g: 0,
        b: 0
      };

      fetch(url + "pixel", {
        method: "POST",
        mode: "no-cors",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
      })
        // .then(d => d.json())
        .then(response => {
          if (!response.ok) {
            console.log(response.statusText);
          }
          console.log("fetch", response);
          this.loadImage();
        });

      console.log("sending", payload);
    },
    loadImage() {
      loadImage(url + "latest").then(image => {
        // console.log(image);
        this.context.drawImage(image, 0, 0, 1000, 1000);
      });
    }
  },
  mounted: function() {
    this.loadImage();
  }
};
</script>

<style scoped>
.canvas {
  position: relative;
}
canvas {
  position: absolute;
  width: 1000px;
  height: 1000px;
}
</style>
