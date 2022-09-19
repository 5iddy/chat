/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "chat.profile";

export interface MsgCreateProfile {
  creator: string;
  name: string;
  bio: string;
  website: string;
}

export interface MsgCreateProfileResponse {}

export interface MsgUpdateProfile {
  creator: string;
  name: string;
  bio: string;
  website: string;
}

export interface MsgUpdateProfileResponse {}

export interface MsgDeleteProfile {
  creator: string;
  name: string;
}

export interface MsgDeleteProfileResponse {}

export interface MsgAddBioToProfile {
  creator: string;
  name: string;
  bio: string;
}

export interface MsgAddBioToProfileResponse {}

export interface MsgAddWebsiteToProfile {
  creator: string;
  name: string;
  website: string;
}

export interface MsgAddWebsiteToProfileResponse {}

const baseMsgCreateProfile: object = {
  creator: "",
  name: "",
  bio: "",
  website: "",
};

export const MsgCreateProfile = {
  encode(message: MsgCreateProfile, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.bio !== "") {
      writer.uint32(26).string(message.bio);
    }
    if (message.website !== "") {
      writer.uint32(34).string(message.website);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateProfile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateProfile } as MsgCreateProfile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.bio = reader.string();
          break;
        case 4:
          message.website = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateProfile {
    const message = { ...baseMsgCreateProfile } as MsgCreateProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.bio !== undefined && object.bio !== null) {
      message.bio = String(object.bio);
    } else {
      message.bio = "";
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = String(object.website);
    } else {
      message.website = "";
    }
    return message;
  },

  toJSON(message: MsgCreateProfile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.bio !== undefined && (obj.bio = message.bio);
    message.website !== undefined && (obj.website = message.website);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateProfile>): MsgCreateProfile {
    const message = { ...baseMsgCreateProfile } as MsgCreateProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.bio !== undefined && object.bio !== null) {
      message.bio = object.bio;
    } else {
      message.bio = "";
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = object.website;
    } else {
      message.website = "";
    }
    return message;
  },
};

const baseMsgCreateProfileResponse: object = {};

export const MsgCreateProfileResponse = {
  encode(
    _: MsgCreateProfileResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateProfileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateProfileResponse,
    } as MsgCreateProfileResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateProfileResponse {
    const message = {
      ...baseMsgCreateProfileResponse,
    } as MsgCreateProfileResponse;
    return message;
  },

  toJSON(_: MsgCreateProfileResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateProfileResponse>
  ): MsgCreateProfileResponse {
    const message = {
      ...baseMsgCreateProfileResponse,
    } as MsgCreateProfileResponse;
    return message;
  },
};

const baseMsgUpdateProfile: object = {
  creator: "",
  name: "",
  bio: "",
  website: "",
};

export const MsgUpdateProfile = {
  encode(message: MsgUpdateProfile, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.bio !== "") {
      writer.uint32(26).string(message.bio);
    }
    if (message.website !== "") {
      writer.uint32(34).string(message.website);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateProfile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateProfile } as MsgUpdateProfile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.bio = reader.string();
          break;
        case 4:
          message.website = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateProfile {
    const message = { ...baseMsgUpdateProfile } as MsgUpdateProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.bio !== undefined && object.bio !== null) {
      message.bio = String(object.bio);
    } else {
      message.bio = "";
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = String(object.website);
    } else {
      message.website = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateProfile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.bio !== undefined && (obj.bio = message.bio);
    message.website !== undefined && (obj.website = message.website);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateProfile>): MsgUpdateProfile {
    const message = { ...baseMsgUpdateProfile } as MsgUpdateProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.bio !== undefined && object.bio !== null) {
      message.bio = object.bio;
    } else {
      message.bio = "";
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = object.website;
    } else {
      message.website = "";
    }
    return message;
  },
};

const baseMsgUpdateProfileResponse: object = {};

export const MsgUpdateProfileResponse = {
  encode(
    _: MsgUpdateProfileResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateProfileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateProfileResponse,
    } as MsgUpdateProfileResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateProfileResponse {
    const message = {
      ...baseMsgUpdateProfileResponse,
    } as MsgUpdateProfileResponse;
    return message;
  },

  toJSON(_: MsgUpdateProfileResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateProfileResponse>
  ): MsgUpdateProfileResponse {
    const message = {
      ...baseMsgUpdateProfileResponse,
    } as MsgUpdateProfileResponse;
    return message;
  },
};

const baseMsgDeleteProfile: object = { creator: "", name: "" };

export const MsgDeleteProfile = {
  encode(message: MsgDeleteProfile, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteProfile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteProfile } as MsgDeleteProfile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteProfile {
    const message = { ...baseMsgDeleteProfile } as MsgDeleteProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteProfile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteProfile>): MsgDeleteProfile {
    const message = { ...baseMsgDeleteProfile } as MsgDeleteProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseMsgDeleteProfileResponse: object = {};

export const MsgDeleteProfileResponse = {
  encode(
    _: MsgDeleteProfileResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteProfileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteProfileResponse,
    } as MsgDeleteProfileResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteProfileResponse {
    const message = {
      ...baseMsgDeleteProfileResponse,
    } as MsgDeleteProfileResponse;
    return message;
  },

  toJSON(_: MsgDeleteProfileResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteProfileResponse>
  ): MsgDeleteProfileResponse {
    const message = {
      ...baseMsgDeleteProfileResponse,
    } as MsgDeleteProfileResponse;
    return message;
  },
};

const baseMsgAddBioToProfile: object = { creator: "", name: "", bio: "" };

export const MsgAddBioToProfile = {
  encode(
    message: MsgAddBioToProfile,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.bio !== "") {
      writer.uint32(26).string(message.bio);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddBioToProfile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddBioToProfile } as MsgAddBioToProfile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.bio = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddBioToProfile {
    const message = { ...baseMsgAddBioToProfile } as MsgAddBioToProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.bio !== undefined && object.bio !== null) {
      message.bio = String(object.bio);
    } else {
      message.bio = "";
    }
    return message;
  },

  toJSON(message: MsgAddBioToProfile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.bio !== undefined && (obj.bio = message.bio);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddBioToProfile>): MsgAddBioToProfile {
    const message = { ...baseMsgAddBioToProfile } as MsgAddBioToProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.bio !== undefined && object.bio !== null) {
      message.bio = object.bio;
    } else {
      message.bio = "";
    }
    return message;
  },
};

const baseMsgAddBioToProfileResponse: object = {};

export const MsgAddBioToProfileResponse = {
  encode(
    _: MsgAddBioToProfileResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddBioToProfileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddBioToProfileResponse,
    } as MsgAddBioToProfileResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddBioToProfileResponse {
    const message = {
      ...baseMsgAddBioToProfileResponse,
    } as MsgAddBioToProfileResponse;
    return message;
  },

  toJSON(_: MsgAddBioToProfileResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddBioToProfileResponse>
  ): MsgAddBioToProfileResponse {
    const message = {
      ...baseMsgAddBioToProfileResponse,
    } as MsgAddBioToProfileResponse;
    return message;
  },
};

const baseMsgAddWebsiteToProfile: object = {
  creator: "",
  name: "",
  website: "",
};

export const MsgAddWebsiteToProfile = {
  encode(
    message: MsgAddWebsiteToProfile,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.website !== "") {
      writer.uint32(26).string(message.website);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddWebsiteToProfile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddWebsiteToProfile } as MsgAddWebsiteToProfile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.website = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddWebsiteToProfile {
    const message = { ...baseMsgAddWebsiteToProfile } as MsgAddWebsiteToProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = String(object.website);
    } else {
      message.website = "";
    }
    return message;
  },

  toJSON(message: MsgAddWebsiteToProfile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.website !== undefined && (obj.website = message.website);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAddWebsiteToProfile>
  ): MsgAddWebsiteToProfile {
    const message = { ...baseMsgAddWebsiteToProfile } as MsgAddWebsiteToProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = object.website;
    } else {
      message.website = "";
    }
    return message;
  },
};

const baseMsgAddWebsiteToProfileResponse: object = {};

export const MsgAddWebsiteToProfileResponse = {
  encode(
    _: MsgAddWebsiteToProfileResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddWebsiteToProfileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddWebsiteToProfileResponse,
    } as MsgAddWebsiteToProfileResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddWebsiteToProfileResponse {
    const message = {
      ...baseMsgAddWebsiteToProfileResponse,
    } as MsgAddWebsiteToProfileResponse;
    return message;
  },

  toJSON(_: MsgAddWebsiteToProfileResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddWebsiteToProfileResponse>
  ): MsgAddWebsiteToProfileResponse {
    const message = {
      ...baseMsgAddWebsiteToProfileResponse,
    } as MsgAddWebsiteToProfileResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateProfile(request: MsgCreateProfile): Promise<MsgCreateProfileResponse>;
  UpdateProfile(request: MsgUpdateProfile): Promise<MsgUpdateProfileResponse>;
  DeleteProfile(request: MsgDeleteProfile): Promise<MsgDeleteProfileResponse>;
  AddBioToProfile(
    request: MsgAddBioToProfile
  ): Promise<MsgAddBioToProfileResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddWebsiteToProfile(
    request: MsgAddWebsiteToProfile
  ): Promise<MsgAddWebsiteToProfileResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateProfile(request: MsgCreateProfile): Promise<MsgCreateProfileResponse> {
    const data = MsgCreateProfile.encode(request).finish();
    const promise = this.rpc.request("chat.profile.Msg", "CreateProfile", data);
    return promise.then((data) =>
      MsgCreateProfileResponse.decode(new Reader(data))
    );
  }

  UpdateProfile(request: MsgUpdateProfile): Promise<MsgUpdateProfileResponse> {
    const data = MsgUpdateProfile.encode(request).finish();
    const promise = this.rpc.request("chat.profile.Msg", "UpdateProfile", data);
    return promise.then((data) =>
      MsgUpdateProfileResponse.decode(new Reader(data))
    );
  }

  DeleteProfile(request: MsgDeleteProfile): Promise<MsgDeleteProfileResponse> {
    const data = MsgDeleteProfile.encode(request).finish();
    const promise = this.rpc.request("chat.profile.Msg", "DeleteProfile", data);
    return promise.then((data) =>
      MsgDeleteProfileResponse.decode(new Reader(data))
    );
  }

  AddBioToProfile(
    request: MsgAddBioToProfile
  ): Promise<MsgAddBioToProfileResponse> {
    const data = MsgAddBioToProfile.encode(request).finish();
    const promise = this.rpc.request(
      "chat.profile.Msg",
      "AddBioToProfile",
      data
    );
    return promise.then((data) =>
      MsgAddBioToProfileResponse.decode(new Reader(data))
    );
  }

  AddWebsiteToProfile(
    request: MsgAddWebsiteToProfile
  ): Promise<MsgAddWebsiteToProfileResponse> {
    const data = MsgAddWebsiteToProfile.encode(request).finish();
    const promise = this.rpc.request(
      "chat.profile.Msg",
      "AddWebsiteToProfile",
      data
    );
    return promise.then((data) =>
      MsgAddWebsiteToProfileResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
