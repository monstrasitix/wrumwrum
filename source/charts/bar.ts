import { IRenderer } from "@/renderer/renderer";

export class BarChart {
  constructor(
    //
    public renderer: IRenderer
  ) {}

  public draw() {
    this.renderer.fill = "dodgerblue";
    this.renderer.rectangle(10, 10, 200, 200);

    this.renderer.fill = "orange";
    this.renderer.circle(100, 100, 50);
  }
}
