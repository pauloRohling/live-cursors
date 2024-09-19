import { ChangeDetectionStrategy, Component, computed, input } from "@angular/core";
import { ColorUtils } from "../../utils/color-utils";
import { PointerComponent } from "../pointer/pointer.component";

@Component({
  selector: "app-cursor",
  standalone: true,
  imports: [PointerComponent],
  templateUrl: "./cursor.component.html",
  styleUrl: "./cursor.component.scss",
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class CursorComponent {
  readonly name = input<string>("");
  readonly color = input.required<string>();

  protected readonly isLight = computed(() => ColorUtils.isLight(this.color()));
}
