(function(t){function n(n){for(var o,c,i=n[0],s=n[1],l=n[2],f=0,p=[];f<i.length;f++)c=i[f],Object.prototype.hasOwnProperty.call(r,c)&&r[c]&&p.push(r[c][0]),r[c]=0;for(o in s)Object.prototype.hasOwnProperty.call(s,o)&&(t[o]=s[o]);u&&u(n);while(p.length)p.shift()();return a.push.apply(a,l||[]),e()}function e(){for(var t,n=0;n<a.length;n++){for(var e=a[n],o=!0,i=1;i<e.length;i++){var s=e[i];0!==r[s]&&(o=!1)}o&&(a.splice(n--,1),t=c(c.s=e[0]))}return t}var o={},r={app:0},a=[];function c(n){if(o[n])return o[n].exports;var e=o[n]={i:n,l:!1,exports:{}};return t[n].call(e.exports,e,e.exports,c),e.l=!0,e.exports}c.m=t,c.c=o,c.d=function(t,n,e){c.o(t,n)||Object.defineProperty(t,n,{enumerable:!0,get:e})},c.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},c.t=function(t,n){if(1&n&&(t=c(t)),8&n)return t;if(4&n&&"object"===typeof t&&t&&t.__esModule)return t;var e=Object.create(null);if(c.r(e),Object.defineProperty(e,"default",{enumerable:!0,value:t}),2&n&&"string"!=typeof t)for(var o in t)c.d(e,o,function(n){return t[n]}.bind(null,o));return e},c.n=function(t){var n=t&&t.__esModule?function(){return t["default"]}:function(){return t};return c.d(n,"a",n),n},c.o=function(t,n){return Object.prototype.hasOwnProperty.call(t,n)},c.p="/c3place/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],s=i.push.bind(i);i.push=n,i=i.slice();for(var l=0;l<i.length;l++)n(i[l]);var u=s;a.push([0,"chunk-vendors"]),e()})({0:function(t,n,e){t.exports=e("56d7")},"034f":function(t,n,e){"use strict";var o=e("85ec"),r=e.n(o);r.a},"56d7":function(t,n,e){"use strict";e.r(n);e("e260"),e("e6cf"),e("cca6"),e("a79d");var o=e("2b0e"),r=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("div",{attrs:{id:"app"}},[e("c3canvas")],1)},a=[],c=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("div",{ref:"container",staticClass:"container"},[e("canvas",{ref:"canvas",style:t.style,attrs:{width:"1000",height:"1000"},on:{click:t.onClick}})])},i=[];e("99af"),e("d3b7");function s(t){return new Promise((function(n,e){var o=new Image;o.onload=function(t){return n(o)},o.onerror=function(t){console.error("could not load"),n(null)},o.src=t}))}var l=e("d934"),u=e("0165"),f=e("00a5"),p="http://localhost:4000/",d={data:function(){return{x:0,y:0,k:1}},computed:{style:function(){var t=this.x,n=this.y,e=this.k;return"transform: translate(".concat(t,"px,").concat(n,"px) scale(").concat(e,")")},canvas:function(){return this.$refs.canvas},context:function(){return this.canvas.getContext("2d")},zoom:function(){return Object(l["a"])().scaleExtent([1,20]).translateExtent([[0,0],[1e3,1e3]]).duration(500).on("zoom",this.zoomed)}},methods:{zoomed:function(){var t=u["c"].transform,n=t.x,e=t.y,o=t.k;this.x=n,this.y=e,this.k=o},onClick:function(t){var n=this,e=this.x,o=this.y,r=this.k,a=l["b"].translate(e,o).scale(r).invert([t.layerX,t.layerY]),c={x:Math.floor(a[0]),y:Math.floor(a[1]),r:255,g:0,b:0};fetch(p+"pixel",{method:"POST",headers:{Accept:"application/json","Content-Type":"application/json"},body:JSON.stringify(c)}).then((function(t){t.ok||console.log(t.statusText),console.log("fetch",t),n.loadImage()})),console.log("sending",c)},loadImage:function(){var t=this,n=p+"latest";s(n).then((function(n){console.log(n),t.context.clearRect(0,0,1e3,1e3),t.context.drawImage(n,0,0,1e3,1e3)}))}},mounted:function(){this.loadImage(),this.container=Object(f["a"])(this.$refs.container).call(this.zoom),this.zoom.scaleTo(this.container,1)}},h=d,v=(e("84d2"),e("2877")),y=Object(v["a"])(h,c,i,!1,null,"3928f2b0",null),m=y.exports,b={name:"app",components:{C3canvas:m}},g=b,x=(e("034f"),Object(v["a"])(g,r,a,!1,null,null,null)),O=x.exports;o["a"].config.productionTip=!1,new o["a"]({render:function(t){return t(O)}}).$mount("#app")},"84d2":function(t,n,e){"use strict";var o=e("d790"),r=e.n(o);r.a},"85ec":function(t,n,e){},d790:function(t,n,e){}});
//# sourceMappingURL=app.14ab1ab1.js.map