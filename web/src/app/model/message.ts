import { Position } from "./position";
import { User } from "./user";
import { MessageType } from "./message-type";

export interface Message {
  type: MessageType;
  data: Position | User;
  timestamp: number;
}
