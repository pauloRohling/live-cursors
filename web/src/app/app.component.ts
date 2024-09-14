import { ChangeDetectionStrategy, Component, inject, OnInit } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { CursorComponent } from "./components/cursor/cursor.component";
import { fromEvent, map, Observable, scan, tap, throttleTime } from "rxjs";
import { WebSocketService } from "./services/websocket.service";
import { Point } from "./model/point";
import { AsyncPipe, JsonPipe } from "@angular/common";

@Component({
  selector: "app-root",
  standalone: true,
  imports: [RouterOutlet, CursorComponent, AsyncPipe, JsonPipe],
  templateUrl: "./app.component.html",
  styleUrl: "./app.component.scss",
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppComponent implements OnInit {
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
