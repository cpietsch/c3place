(function(n){function t(t){for(var o,c,i=t[0],s=t[1],u=t[2],f=0,p=[];f<i.length;f++)c=i[f],Object.prototype.hasOwnProperty.call(r,c)&&r[c]&&p.push(r[c][0]),r[c]=0;for(o in s)Object.prototype.hasOwnProperty.call(s,o)&&(n[o]=s[o]);l&&l(t);while(p.length)p.shift()();return a.push.apply(a,u||[]),e()}function e(){for(var n,t=0;t<a.length;t++){for(var e=a[t],o=!0,i=1;i<e.length;i++){var s=e[i];0!==r[s]&&(o=!1)}o&&(a.splice(t--,1),n=c(c.s=e[0]))}return n}var o={},r={app:0},a=[];function c(t){if(o[t])return o[t].exports;var e=o[t]={i:t,l:!1,exports:{}};return n[t].call(e.exports,e,e.exports,c),e.l=!0,e.exports}c.m=n,c.c=o,c.d=function(n,t,e){c.o(n,t)||Object.defineProperty(n,t,{enumerable:!0,get:e})},c.r=function(n){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(n,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(n,"__esModule",{value:!0})},c.t=function(n,t){if(1&t&&(n=c(n)),8&t)return n;if(4&t&&"object"===typeof n&&n&&n.__esModule)return n;var e=Object.create(null);if(c.r(e),Object.defineProperty(e,"default",{enumerable:!0,value:n}),2&t&&"string"!=typeof n)for(var o in n)c.d(e,o,function(t){return n[t]}.bind(null,o));return e},c.n=function(n){var t=n&&n.__esModule?function(){return n["default"]}:function(){return n};return c.d(t,"a",t),t},c.o=function(n,t){return Object.prototype.hasOwnProperty.call(n,t)},c.p="/c3place/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],s=i.push.bind(i);i.push=t,i=i.slice();for(var u=0;u<i.length;u++)t(i[u]);var l=s;a.push([0,"chunk-vendors"]),e()})({0:function(n,t,e){n.exports=e("56d7")},"034f":function(n,t,e){"use strict";var o=e("85ec"),r=e.n(o);r.a},4872:function(n,t,e){"use strict";var o=e("5f2e"),r=e.n(o);r.a},"56d7":function(n,t,e){"use strict";e.r(t);e("e260"),e("e6cf"),e("cca6"),e("a79d");var o=e("2b0e"),r=function(){var n=this,t=n.$createElement,e=n._self._c||t;return e("div",{attrs:{id:"app"}},[e("c3canvas")],1)},a=[],c=function(){var n=this,t=n.$createElement,e=n._self._c||t;return e("div",{ref:"container",staticClass:"container"},[e("canvas",{ref:"canvas",style:n.style,attrs:{width:"1000",height:"1000"},on:{click:n.onClick}})])},i=[];e("99af"),e("d3b7");function s(n){return new Promise((function(t,e){var o=new Image;o.onload=function(n){return t(o)},o.onerror=function(n){console.error("could not load"),t(null)},o.src=n}))}var u=e("d934"),l=e("0165"),f=e("00a5"),p="http://localhost:4000/",d={data:function(){return{x:0,y:0,k:1}},computed:{style:function(){var n=this.x,t=this.y,e=this.k;return"transform: translate(".concat(n,"px,").concat(t,"px) scale(").concat(e,")")},canvas:function(){return this.$refs.canvas},context:function(){return this.canvas.getContext("2d")},zoom:function(){return Object(u["a"])().scaleExtent([1,20]).translateExtent([[0,0],[1e3,1e3]]).duration(500).on("zoom",this.zoomed)}},methods:{zoomed:function(){var n=l["c"].transform,t=n.x,e=n.y,o=n.k;this.x=t,this.y=e,this.k=o},onClick:function(n){var t=this,e={x:n.layerX,y:n.layerY,r:255,g:0,b:0};fetch(p+"pixel",{method:"POST",mode:"no-cors",headers:{Accept:"application/json","Content-Type":"application/json"},body:JSON.stringify(e)}).then((function(n){n.ok||console.log(n.statusText),console.log("fetch",n),t.loadImage()})),console.log("sending",e)},loadImage:function(){var n=this,t="canvas.png";s(t).then((function(t){n.context.drawImage(t,0,0,1e3,1e3)}))}},mounted:function(){this.loadImage(),this.container=Object(f["a"])(this.$refs.container).call(this.zoom),this.zoom.scaleTo(this.container,1)}},h=d,v=(e("4872"),e("2877")),m=Object(v["a"])(h,c,i,!1,null,"30535915",null),y=m.exports,b={name:"app",components:{C3canvas:y}},g=b,x=(e("034f"),Object(v["a"])(g,r,a,!1,null,null,null)),O=x.exports;o["a"].config.productionTip=!1,new o["a"]({render:function(n){return n(O)}}).$mount("#app")},"5f2e":function(n,t,e){},"85ec":function(n,t,e){}});
//# sourceMappingURL=app.9577c812.js.map