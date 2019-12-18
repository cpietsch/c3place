export function loadImage(url) {
  return new Promise((resolve, error) => {
    const image = new Image();
    image.onload = _ => resolve(image);
    image.onerror = _ => {
      console.error("could not load");
      resolve(null);
    };
    // image.crossOrigin = "";
    image.src = url;
  });
}
