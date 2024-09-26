import { inject, Injectable } from "@angular/core";
import { WebSocketService } from "./websocket.service";
import { filter, map, scan, share } from "rxjs";
import { MessageType } from "../model/message-type";
import { User } from "../model/user";
import { Cursor } from "../model/cursor";
import { Position } from "../model/position";

type CursorState = {
  active: User | undefined;
  cursors: Map<string, Cursor>;
};

@Injectable({ providedIn: "root" })
export class CursorService {
  private readonly websocketService = inject(WebSocketService);

  private readonly state$ = this.websocketService.messages$.pipe(
    scan(
      (accumulator, current) => {
        if (current.type === MessageType.SELF) {
          return { ...accumulator, active: current.data as User };
        }

        if (current.type === MessageType.USER) {
          const cursor = { ...(current.data as User), x: 0, y: 0 };
          accumulator.cursors.set(current.data.id, cursor);
          return accumulator;
        }

        const cursor = accumulator.cursors.get(current.data.id);
        if (cursor) {
          const position = current.data as Position;
          accumulator.cursors.set(cursor.id, { ...cursor, x: position.x, y: position.y });
        }
        return accumulator;
      },
      <CursorState>{ active: undefined, cursors: new Map() },
    ),
    share(),
  );

  readonly active$ = this.state$.pipe(
    map((state) => state.active),
    filter(Boolean),
  );

  readonly cursors$ = this.state$.pipe(map((state) => Array.from(state.cursors.values())));
}
