<template>
  <div class="container" ref="container">
    <canvas ref="canvas" width="1000" height="1000" :style="style"></canvas>
  </div>
</template>

<script>
import { loadImage } from "../utils.js";
import { zoom, zoomIdentity, zoomTransform } from "d3-zoom";
import { select, event, mouse } from "d3-selection";
import { interval } from "d3-timer";

// const url = "http://localhost:4000/";
const url = "http://78.47.194.107:4000/";

export default {
  data: function() {
    return {
      x: 0,
      y: 0,
      k: 1
    };
  },
  computed: {
    style: function() {
      const { x, y, k } = this;
      return `transform: translate(${x}px,${y}px) scale(${k})`;
    },
    canvas: function() {
      return this.$refs.canvas;
    },
    context: function() {
      return this.canvas.getContext("2d");
    },
    zoom: function() {
      return zoom()
        .scaleExtent([1, 30])
        .translateExtent([
          [0, 0],
          [1000, 1000]
        ])
        .duration(500)
        .on("zoom", this.zoomed);
    }
  },
  methods: {
    zoomRandom: function() {
      const x = parseInt(Math.random()*1000)
      const y = parseInt(Math.random()*1000)
      const zoom = 1+Math.random()*4
      this.container
        .transition()
        .duration(5000)
        // .call(this.zoom.transform, zoomIdentity.scale(1))
        .call(
          this.zoom.transform,
          zoomIdentity.translate(500,500).scale(zoom).translate(-x, -y)
        )
    },
    zoomed: function() {
      // console.log("zoomed", event.transform);
      const { x, y, k } = event.transform;
      this.x = x;
      this.y = y;
      this.k = k;
    },
    loadImage() {
      const imageUrl = url + "latest?" + Date.now();
      // const imageUrl = "canvas.png";
      loadImage(imageUrl).then(image => {
        // console.log(image);
        // this.context.clearRect(0, 0, 1000, 1000);
        this.context.drawImage(image, 0, 0, 1000, 1000);
      });
    }
  },
  mounted: function() {
    interval(this.loadImage, 4000);
    this.loadImage();
    this.container = select(this.$refs.container).call(this.zoom);
    // .on("click", this.onClick);
    // this.zoom.scaleTo(this.container, 1);

    interval(this.zoomRandom, 20000);
    this.zoomRandom()
  }
};
</script>

<style scoped>
.container {
  position: absolute;
  top: 0px;
  width: 100%;
  height: 100vh;
  overflow: hidden;
}
canvas {
  position: absolute;
  width: 1000px;
  height: 1000px;
  will-change: transform;
  transform-origin: 0 0;
  image-rendering: pixelated;
  /* border: 1px solid #000; */
}
</style>
