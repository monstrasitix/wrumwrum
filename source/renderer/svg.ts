import type { IRenderer } from "./renderer";

export class SVGRenderer implements IRenderer {
  private svg: SVGElement;
  private namespace = "http://www.w3.org/2000/svg";

  public fill: string | null = null;
  public stroke: string | null = null;

  constructor() {
    this.svg = document.createElementNS(this.namespace, "svg") as SVGElement;
    this.svg.setAttribute("xmlns", this.namespace);
  }

  public mount(target: HTMLElement): void {
    target.appendChild(this.svg);
  }

  get size(): [number, number] {
    return [
      //
      this.svg.clientWidth,
      this.svg.clientHeight,
    ];
  }

  set size([w, h]: [number, number]) {
    this.svg.setAttribute("width", w.toString());
    this.svg.setAttribute("height", h.toString());
  }

  private applyColor(element: Element) {
    if (this.fill) {
      element.setAttribute("fill", this.fill);
    }

    if (this.stroke) {
      element.setAttribute("stroke", this.stroke);
    }
  }

  public rectangle(x: number, y: number, width: number, height: number) {
    const rect = document.createElementNS(this.namespace, "rect");
    rect.setAttribute("x", x.toString());
    rect.setAttribute("y", y.toString());
    rect.setAttribute("width", width.toString());
    rect.setAttribute("height", height.toString());

    this.applyColor(rect);
    this.svg.appendChild(rect);
  }

  public circle(x: number, y: number, radius: number) {
    const circle = document.createElementNS(this.namespace, "circle");
    circle.setAttribute("cx", x.toString());
    circle.setAttribute("cy", y.toString());
    circle.setAttribute("r", radius.toString());

    this.applyColor(circle);
    this.svg.appendChild(circle);
  }
}
