import type { IRenderer } from "./renderer";

export class CanvasRenderer implements IRenderer {
  private ctx: CanvasRenderingContext2D;

  public fill: string | null = null;
  public stroke: string | null = null;

  constructor() {
    this.ctx = document.createElement("canvas").getContext("2d")!;
  }

  public mount(target: HTMLElement): void {
    target.appendChild(this.ctx.canvas);
  }

  get size(): [number, number] {
    return [
      //
      this.ctx.canvas.width,
      this.ctx.canvas.height,
    ];
  }

  set size([width, height]: [number, number]) {
    this.ctx.canvas.width = width;
    this.ctx.canvas.height = height;
  }

  private applyColor() {
    if (this.fill) {
      this.ctx.fillStyle = this.fill;
    }

    if (this.stroke) {
      this.ctx.strokeStyle = this.stroke;
    }
  }

  public rectangle(x: number, y: number, width: number, height: number): void {
    this.applyColor();
    this.ctx.fillRect(x, y, width, height);
  }

  public circle(x: number, y: number, radius: number): void {
    this.applyColor();
    this.ctx.arc(x, y, radius, 0, Math.PI * 2, false);
  }
}
