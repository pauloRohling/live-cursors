import { ChangeDetectionStrategy, Component } from "@angular/core";

@Component({
  selector: "app-pointer",
  standalone: true,
  imports: [],
  templateUrl: "./pointer.component.html",
  styleUrl: "./pointer.component.scss",
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PointerComponent {}
