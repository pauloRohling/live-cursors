import { Injectable } from "@angular/core";
import { environment } from "../../environments/environment";
import { webSocket } from "rxjs/webSocket";
import { MessageType } from "../model/message-type";
import { Message } from "../model/message";

@Injectable({ providedIn: "root" })
export class WebSocketService {
  private readonly socket = webSocket<Message>({ url: environment.ws });

  readonly messages$ = this.socket.asObservable();

  send(content: any): void {
    if (!(content instanceof Object)) {
      return;
    }

    const payload = {
      data: content,
      type: MessageType.POSITION,
      timestamp: new Date().getUTCMilliseconds(),
    };

    this.socket.next(payload);
  }
}
