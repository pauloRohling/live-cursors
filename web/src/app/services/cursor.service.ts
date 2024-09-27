import { inject, Injectable } from "@angular/core";
import { WebSocketService } from "./websocket.service";
import { filter, map, scan, share } from "rxjs";
import { MessageType } from "../model/message-type";
import { Client } from "../model/client";
import { Cursor } from "../model/cursor";
import { Position } from "../model/position";

type CursorState = {
  active: Client | undefined;
  cursors: Map<string, Cursor>;
};

@Injectable({ providedIn: "root" })
export class CursorService {
  private readonly websocketService = inject(WebSocketService);

  private readonly state$ = this.websocketService.messages$.pipe(
    scan(
      (accumulator, current) => {
        if (current.type === MessageType.POSITION) {
          const position = current.data as Position;
          const cursor = accumulator.cursors.get(position.id);
          if (cursor) {
            accumulator.cursors.set(cursor.id, { ...cursor, x: position.x, y: position.y });
          }
          return accumulator;
        }

        const client = current.data as Client;

        if (current.type === MessageType.SELF) {
          return { ...accumulator, active: client };
        }

        if (current.type === MessageType.REMOVE) {
          accumulator.cursors.delete(client.id);
          return accumulator;
        }

        const cursor = { ...client, x: -50, y: -50 };
        accumulator.cursors.set(client.id, cursor);
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
