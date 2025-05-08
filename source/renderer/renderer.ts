export interface IRenderer {
  mount(target: HTMLElement): void;

  get size(): [number, number];
  set size(dimensions: [number, number]);

  mount(target: HTMLElement): void;

  rectangle(x: number, y: number, width: number, height: number): void;
  circle(x: number, y: number, radius: number): void;
}
