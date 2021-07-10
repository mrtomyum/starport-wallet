/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "mrtomyum.wallet.dex";

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCancelSellOrder {
  creator: string;
  amountDenom: string;
  orderID: number;
}

export interface MsgCancelSellOrderResponse {}

export interface MsgSendSellOrder {
  sender: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  amountDenom: string;
  amount: number;
  price: number;
}

export interface MsgSendSellOrderResponse {}

const baseMsgCancelSellOrder: object = {
  creator: "",
  amountDenom: "",
  orderID: 0,
};

export const MsgCancelSellOrder = {
  encode(
    message: MsgCancelSellOrder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amountDenom !== "") {
      writer.uint32(18).string(message.amountDenom);
    }
    if (message.orderID !== 0) {
      writer.uint32(24).int32(message.orderID);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelSellOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelSellOrder } as MsgCancelSellOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amountDenom = reader.string();
          break;
        case 3:
          message.orderID = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelSellOrder {
    const message = { ...baseMsgCancelSellOrder } as MsgCancelSellOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amountDenom !== undefined && object.amountDenom !== null) {
      message.amountDenom = String(object.amountDenom);
    } else {
      message.amountDenom = "";
    }
    if (object.orderID !== undefined && object.orderID !== null) {
      message.orderID = Number(object.orderID);
    } else {
      message.orderID = 0;
    }
    return message;
  },

  toJSON(message: MsgCancelSellOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amountDenom !== undefined &&
      (obj.amountDenom = message.amountDenom);
    message.orderID !== undefined && (obj.orderID = message.orderID);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCancelSellOrder>): MsgCancelSellOrder {
    const message = { ...baseMsgCancelSellOrder } as MsgCancelSellOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amountDenom !== undefined && object.amountDenom !== null) {
      message.amountDenom = object.amountDenom;
    } else {
      message.amountDenom = "";
    }
    if (object.orderID !== undefined && object.orderID !== null) {
      message.orderID = object.orderID;
    } else {
      message.orderID = 0;
    }
    return message;
  },
};

const baseMsgCancelSellOrderResponse: object = {};

export const MsgCancelSellOrderResponse = {
  encode(
    _: MsgCancelSellOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCancelSellOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCancelSellOrderResponse,
    } as MsgCancelSellOrderResponse;
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

  fromJSON(_: any): MsgCancelSellOrderResponse {
    const message = {
      ...baseMsgCancelSellOrderResponse,
    } as MsgCancelSellOrderResponse;
    return message;
  },

  toJSON(_: MsgCancelSellOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCancelSellOrderResponse>
  ): MsgCancelSellOrderResponse {
    const message = {
      ...baseMsgCancelSellOrderResponse,
    } as MsgCancelSellOrderResponse;
    return message;
  },
};

const baseMsgSendSellOrder: object = {
  sender: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  amountDenom: "",
  amount: 0,
  price: 0,
};

export const MsgSendSellOrder = {
  encode(message: MsgSendSellOrder, writer: Writer = Writer.create()): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.amountDenom !== "") {
      writer.uint32(42).string(message.amountDenom);
    }
    if (message.amount !== 0) {
      writer.uint32(48).int32(message.amount);
    }
    if (message.price !== 0) {
      writer.uint32(56).int32(message.price);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendSellOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendSellOrder } as MsgSendSellOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.amountDenom = reader.string();
          break;
        case 6:
          message.amount = reader.int32();
          break;
        case 7:
          message.price = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendSellOrder {
    const message = { ...baseMsgSendSellOrder } as MsgSendSellOrder;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.amountDenom !== undefined && object.amountDenom !== null) {
      message.amountDenom = String(object.amountDenom);
    } else {
      message.amountDenom = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = Number(object.price);
    } else {
      message.price = 0;
    }
    return message;
  },

  toJSON(message: MsgSendSellOrder): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.amountDenom !== undefined &&
      (obj.amountDenom = message.amountDenom);
    message.amount !== undefined && (obj.amount = message.amount);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendSellOrder>): MsgSendSellOrder {
    const message = { ...baseMsgSendSellOrder } as MsgSendSellOrder;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.amountDenom !== undefined && object.amountDenom !== null) {
      message.amountDenom = object.amountDenom;
    } else {
      message.amountDenom = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = 0;
    }
    return message;
  },
};

const baseMsgSendSellOrderResponse: object = {};

export const MsgSendSellOrderResponse = {
  encode(
    _: MsgSendSellOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendSellOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendSellOrderResponse,
    } as MsgSendSellOrderResponse;
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

  fromJSON(_: any): MsgSendSellOrderResponse {
    const message = {
      ...baseMsgSendSellOrderResponse,
    } as MsgSendSellOrderResponse;
    return message;
  },

  toJSON(_: MsgSendSellOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendSellOrderResponse>
  ): MsgSendSellOrderResponse {
    const message = {
      ...baseMsgSendSellOrderResponse,
    } as MsgSendSellOrderResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CancelSellOrder(
    request: MsgCancelSellOrder
  ): Promise<MsgCancelSellOrderResponse>;
  SendSellOrder(request: MsgSendSellOrder): Promise<MsgSendSellOrderResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CancelSellOrder(
    request: MsgCancelSellOrder
  ): Promise<MsgCancelSellOrderResponse> {
    const data = MsgCancelSellOrder.encode(request).finish();
    const promise = this.rpc.request(
      "mrtomyum.wallet.dex.Msg",
      "CancelSellOrder",
      data
    );
    return promise.then((data) =>
      MsgCancelSellOrderResponse.decode(new Reader(data))
    );
  }

  SendSellOrder(request: MsgSendSellOrder): Promise<MsgSendSellOrderResponse> {
    const data = MsgSendSellOrder.encode(request).finish();
    const promise = this.rpc.request(
      "mrtomyum.wallet.dex.Msg",
      "SendSellOrder",
      data
    );
    return promise.then((data) =>
      MsgSendSellOrderResponse.decode(new Reader(data))
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
