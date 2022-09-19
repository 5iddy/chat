/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chat.chat";

export interface Letter {
  from: string;
  to: string;
  content: string;
  contentType: string;
}

const baseLetter: object = { from: "", to: "", content: "", contentType: "" };

export const Letter = {
  encode(message: Letter, writer: Writer = Writer.create()): Writer {
    if (message.from !== "") {
      writer.uint32(10).string(message.from);
    }
    if (message.to !== "") {
      writer.uint32(18).string(message.to);
    }
    if (message.content !== "") {
      writer.uint32(26).string(message.content);
    }
    if (message.contentType !== "") {
      writer.uint32(34).string(message.contentType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Letter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLetter } as Letter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.from = reader.string();
          break;
        case 2:
          message.to = reader.string();
          break;
        case 3:
          message.content = reader.string();
          break;
        case 4:
          message.contentType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Letter {
    const message = { ...baseLetter } as Letter;
    if (object.from !== undefined && object.from !== null) {
      message.from = String(object.from);
    } else {
      message.from = "";
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = String(object.to);
    } else {
      message.to = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = String(object.content);
    } else {
      message.content = "";
    }
    if (object.contentType !== undefined && object.contentType !== null) {
      message.contentType = String(object.contentType);
    } else {
      message.contentType = "";
    }
    return message;
  },

  toJSON(message: Letter): unknown {
    const obj: any = {};
    message.from !== undefined && (obj.from = message.from);
    message.to !== undefined && (obj.to = message.to);
    message.content !== undefined && (obj.content = message.content);
    message.contentType !== undefined &&
      (obj.contentType = message.contentType);
    return obj;
  },

  fromPartial(object: DeepPartial<Letter>): Letter {
    const message = { ...baseLetter } as Letter;
    if (object.from !== undefined && object.from !== null) {
      message.from = object.from;
    } else {
      message.from = "";
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = object.to;
    } else {
      message.to = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = object.content;
    } else {
      message.content = "";
    }
    if (object.contentType !== undefined && object.contentType !== null) {
      message.contentType = object.contentType;
    } else {
      message.contentType = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
