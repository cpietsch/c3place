<template>
  <div class="container" ref="container">
    <canvas ref="canvas" @click="onClick" width="1000" height="1000" :style="style"></canvas>
  </div>
</template>

<script>
import { loadImage } from "../utils.js";
import { zoom, zoomIdentity, zoomTransform } from "d3-zoom";
import { select, event, mouse } from "d3-selection";
import { interval } from "d3-timer";

// const url = "http://localhost:4000/";
const url = "http://bd0a681b.ngrok.io/";

export default {
  props: ["color"],
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
        .scaleExtent([1, 20])
        .translateExtent([
          [0, 0],
          [1000, 1000]
        ])
        .duration(500)
        .on("zoom", this.zoomed);
    }
  },
  methods: {
    zoomed: function() {
      // console.log("zoomed", event.transform);
      const { x, y, k } = event.transform;
      this.x = x;
      this.y = y;
      this.k = k;
    },
    onClick: function(el) {
      console.log("onClick", el);
      const { x, y, k } = this;
      const [r, g, b] = this.color;
      const paddingLeft = 40; // todo
      const inverted = zoomIdentity
        .translate(x, y)
        .scale(k)
        .invert([el.x - paddingLeft, el.y]);

      const payload = {
        x: Math.floor(inverted[0]),
        y: Math.floor(inverted[1]),
        r,
        g,
        b
      };

      this.context.fillStyle = "rgba(" + payload.r + "," + payload.g + "," + payload.b + ",1)";
      this.context.fillRect(payload.x, payload.y, 1, 1);

      fetch(url + "pixel", {
        method: "POST",
        // mode: "no-cors",
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
          // console.log("fetch", response);
          // this.loadImage();
        });

      console.log("sending", payload);
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
    interval(this.loadImage, 5000);
    this.loadImage();
    this.container = select(this.$refs.container).call(this.zoom);
    // .on("click", this.onClick);
    this.zoom.scaleTo(this.container, 1);

    // this.container
    //   .call(
    //     this.zoom.transform,
    //     zoomIdentity.translate(-1000 / 2, -1000 / 2).scale(3)
    //   )
    //   .transition()
    //   .duration(1000)
    //   .call(this.zoom.transform, zoomIdentity.scale(1));
  }
};
</script>

<style scoped>
.container {
  position: absolute;
  left: 40px;
  top: 0px;
  width: calc(100% - 40px);
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
