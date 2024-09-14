import { Injectable } from "@angular/core";
import { environment } from "../../environments/environment";
import { webSocket } from "rxjs/webSocket";

@Injectable({ providedIn: "root" })
export class WebSocketService {
  private readonly socket = webSocket({
    url: environment.ws,
  });

  readonly messages$ = this.socket.asObservable();

  send(content: any): void {
    if (!(content instanceof Object)) {
      return;
    }

    let payload: string;
    try {
      payload = JSON.stringify(content);
    } catch (e) {
      console.error(e);
      return;
    }

    this.socket.next(payload);
  }
}
