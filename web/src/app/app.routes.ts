import { Routes } from "@angular/router";

export const routes: Routes = [
  {
    path: "",
    pathMatch: "full",
    loadComponent: () => import("./pages/canvas-page/canvas-page.component").then((value) => value.CanvasPageComponent),
  },
];
