/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chat.chat";

export interface LetterBox {
  count: number;
  messages: string;
}

const baseLetterBox: object = { count: 0, messages: "" };

export const LetterBox = {
  encode(message: LetterBox, writer: Writer = Writer.create()): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    if (message.messages !== "") {
      writer.uint32(18).string(message.messages);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LetterBox {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLetterBox } as LetterBox;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.messages = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LetterBox {
    const message = { ...baseLetterBox } as LetterBox;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.messages !== undefined && object.messages !== null) {
      message.messages = String(object.messages);
    } else {
      message.messages = "";
    }
    return message;
  },

  toJSON(message: LetterBox): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.messages !== undefined && (obj.messages = message.messages);
    return obj;
  },

  fromPartial(object: DeepPartial<LetterBox>): LetterBox {
    const message = { ...baseLetterBox } as LetterBox;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.messages !== undefined && object.messages !== null) {
      message.messages = object.messages;
    } else {
      message.messages = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
