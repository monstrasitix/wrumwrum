import { SVGRenderer } from "@/renderer/svg";

document.addEventListener("DOMContentLoaded", function (this: Document) {
  const target = this.getElementById("charts");
  if (target === null) {
    throw new Error("No target");
  }

  const renderer = new SVGRenderer();
  renderer.mount(target);
  renderer.size = [800, 600];

  renderer.fill = "red";
  renderer.rectangle(10, 10, 200, 200);

  renderer.fill = "orange";
  renderer.circle(100, 100, 50);
});
