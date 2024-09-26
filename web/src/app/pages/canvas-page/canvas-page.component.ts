import { ChangeDetectionStrategy, Component, inject, OnInit } from "@angular/core";
import { AsyncPipe, JsonPipe } from "@angular/common";
import { CursorComponent } from "../../components/cursor/cursor.component";
import { fromEvent, map, tap, throttleTime, withLatestFrom } from "rxjs";
import { CursorService } from "../../services/cursor.service";
import { WebSocketService } from "../../services/websocket.service";

@Component({
  selector: "app-canvas-page",
  standalone: true,
  imports: [AsyncPipe, CursorComponent, JsonPipe],
  templateUrl: "./canvas-page.component.html",
  styleUrl: "./canvas-page.component.scss",
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class CanvasPageComponent implements OnInit {
  private readonly cursorService = inject(CursorService);
  private readonly websocketService = inject(WebSocketService);

  protected readonly cursors$ = this.cursorService.cursors$;

  ngOnInit(): void {
    fromEvent(window, "mousemove")
      .pipe(
        throttleTime(60),
        map((event) => {
          const mouseEvent = event as MouseEvent;
          return { x: mouseEvent.x, y: mouseEvent.y };
        }),
        withLatestFrom(this.cursorService.active$),
        tap(([point, user]) => this.websocketService.send({ id: user.id, ...point })),
      )
      .subscribe();
  }
}
