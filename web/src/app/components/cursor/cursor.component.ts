import { ChangeDetectionStrategy, Component, computed, input } from "@angular/core";
import { ColorUtils } from "../../utils/color-utils";

@Component({
  selector: "app-cursor",
  standalone: true,
  imports: [],
  templateUrl: "./cursor.component.html",
  styleUrl: "./cursor.component.scss",
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class CursorComponent {
  readonly name = input.required<string>();
  readonly color = input.required<string>();

  protected readonly isLight = computed(() => ColorUtils.isLight(this.color()));
}
