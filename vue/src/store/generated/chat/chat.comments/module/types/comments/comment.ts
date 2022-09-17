/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chat.comments";

export interface Comment {
  id: number;
  content: string;
  replyTo: string;
  createdAt: number;
  extention: string;
  extensionType: string;
  creator: string;
}

const baseComment: object = {
  id: 0,
  content: "",
  replyTo: "",
  createdAt: 0,
  extention: "",
  extensionType: "",
  creator: "",
};

export const Comment = {
  encode(message: Comment, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.content !== "") {
      writer.uint32(18).string(message.content);
    }
    if (message.replyTo !== "") {
      writer.uint32(26).string(message.replyTo);
    }
    if (message.createdAt !== 0) {
      writer.uint32(32).int32(message.createdAt);
    }
    if (message.extention !== "") {
      writer.uint32(42).string(message.extention);
    }
    if (message.extensionType !== "") {
      writer.uint32(50).string(message.extensionType);
    }
    if (message.creator !== "") {
      writer.uint32(58).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Comment {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseComment } as Comment;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.content = reader.string();
          break;
        case 3:
          message.replyTo = reader.string();
          break;
        case 4:
          message.createdAt = reader.int32();
          break;
        case 5:
          message.extention = reader.string();
          break;
        case 6:
          message.extensionType = reader.string();
          break;
        case 7:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Comment {
    const message = { ...baseComment } as Comment;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = String(object.content);
    } else {
      message.content = "";
    }
    if (object.replyTo !== undefined && object.replyTo !== null) {
      message.replyTo = String(object.replyTo);
    } else {
      message.replyTo = "";
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = Number(object.createdAt);
    } else {
      message.createdAt = 0;
    }
    if (object.extention !== undefined && object.extention !== null) {
      message.extention = String(object.extention);
    } else {
      message.extention = "";
    }
    if (object.extensionType !== undefined && object.extensionType !== null) {
      message.extensionType = String(object.extensionType);
    } else {
      message.extensionType = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Comment): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.content !== undefined && (obj.content = message.content);
    message.replyTo !== undefined && (obj.replyTo = message.replyTo);
    message.createdAt !== undefined && (obj.createdAt = message.createdAt);
    message.extention !== undefined && (obj.extention = message.extention);
    message.extensionType !== undefined &&
      (obj.extensionType = message.extensionType);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Comment>): Comment {
    const message = { ...baseComment } as Comment;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = object.content;
    } else {
      message.content = "";
    }
    if (object.replyTo !== undefined && object.replyTo !== null) {
      message.replyTo = object.replyTo;
    } else {
      message.replyTo = "";
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = object.createdAt;
    } else {
      message.createdAt = 0;
    }
    if (object.extention !== undefined && object.extention !== null) {
      message.extention = object.extention;
    } else {
      message.extention = "";
    }
    if (object.extensionType !== undefined && object.extensionType !== null) {
      message.extensionType = object.extensionType;
    } else {
      message.extensionType = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
