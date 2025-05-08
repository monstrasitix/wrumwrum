await Bun.build({
  outdir: "./static/scripts",
  target: "browser",
  entrypoints: [
    //
    "./source/charts.ts",
  ],
});
