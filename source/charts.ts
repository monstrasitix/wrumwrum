import { BarChart } from "@/charts/bar";
import { SVGRenderer } from "@/renderer/svg";

document.addEventListener("DOMContentLoaded", function (this: Document) {
  const target = this.getElementById("charts");
  if (target === null) {
    throw new Error("No target");
  }

  const renderer = new SVGRenderer();
  renderer.mount(target);
  renderer.size = [800, 600];

  const chart = new BarChart(renderer);
  chart.draw();
});
