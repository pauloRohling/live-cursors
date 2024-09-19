import { ChangeDetectionStrategy, Component, inject, OnInit } from "@angular/core";
import { AsyncPipe, JsonPipe } from "@angular/common";
import { CursorComponent } from "../../components/cursor/cursor.component";
import { WebSocketService } from "../../services/websocket.service";
import { fromEvent, map, Observable, scan, tap, throttleTime } from "rxjs";
import { Point } from "../../model/point";

@Component({
  selector: "app-canvas-page",
  standalone: true,
  imports: [AsyncPipe, CursorComponent, JsonPipe],
  templateUrl: "./canvas-page.component.html",
  styleUrl: "./canvas-page.component.scss",
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class CanvasPageComponent implements OnInit {
  private readonly websocketService = inject(WebSocketService);
  private readonly userUuid = crypto.randomUUID();

  protected readonly cursors: Observable<any> = this.websocketService.messages$.pipe(
    scan((cursors, payload: any) => {
      const cursor = JSON.parse(payload);
      const userUuid = cursor["userUuid"];
      if (userUuid === this.userUuid) {
        return cursors;
      }
      return { ...cursors, [userUuid]: cursor };
    }, {}),
    map((cursors) => Object.values(cursors)),
  );

  ngOnInit(): void {
    fromEvent(window, "mousemove")
      .pipe(
        throttleTime(60),
        map((event) => {
          const mouseEvent = event as MouseEvent;
          return { x: mouseEvent.x, y: mouseEvent.y };
        }),
        tap((point: Point) => this.websocketService.send({ userUuid: this.userUuid, point })),
      )
      .subscribe();

    this.websocketService.messages$.subscribe(console.log);
  }
}
