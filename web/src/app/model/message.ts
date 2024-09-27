import { Position } from "./position";
import { Client } from "./client";
import { MessageType } from "./message-type";

export interface Message {
  type: MessageType;
  data: Position | Client;
  timestamp: number;
}
