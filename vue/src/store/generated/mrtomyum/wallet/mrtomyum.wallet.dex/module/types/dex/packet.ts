/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mrtomyum.wallet.dex";

export interface DexPacketData {
  noData: NoData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  sellOrderPacket: SellOrderPacketData | undefined;
}

export interface NoData {}

/**
 * this line is used by starport scaffolding # ibc/packet/proto/message
 * SellOrderPacketData defines a struct for the packet payload
 */
export interface SellOrderPacketData {
  amountDenom: string;
  amount: number;
  price: number;
}

/** SellOrderPacketAck defines a struct for the packet acknowledgment */
export interface SellOrderPacketAck {}

const baseDexPacketData: object = {};

export const DexPacketData = {
  encode(message: DexPacketData, writer: Writer = Writer.create()): Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
    }
    if (message.sellOrderPacket !== undefined) {
      SellOrderPacketData.encode(
        message.sellOrderPacket,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DexPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDexPacketData } as DexPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32());
          break;
        case 2:
          message.sellOrderPacket = SellOrderPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DexPacketData {
    const message = { ...baseDexPacketData } as DexPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromJSON(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.sellOrderPacket !== undefined &&
      object.sellOrderPacket !== null
    ) {
      message.sellOrderPacket = SellOrderPacketData.fromJSON(
        object.sellOrderPacket
      );
    } else {
      message.sellOrderPacket = undefined;
    }
    return message;
  },

  toJSON(message: DexPacketData): unknown {
    const obj: any = {};
    message.noData !== undefined &&
      (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
    message.sellOrderPacket !== undefined &&
      (obj.sellOrderPacket = message.sellOrderPacket
        ? SellOrderPacketData.toJSON(message.sellOrderPacket)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<DexPacketData>): DexPacketData {
    const message = { ...baseDexPacketData } as DexPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromPartial(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.sellOrderPacket !== undefined &&
      object.sellOrderPacket !== null
    ) {
      message.sellOrderPacket = SellOrderPacketData.fromPartial(
        object.sellOrderPacket
      );
    } else {
      message.sellOrderPacket = undefined;
    }
    return message;
  },
};

const baseNoData: object = {};

export const NoData = {
  encode(_: NoData, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNoData } as NoData;
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

  fromJSON(_: any): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },

  toJSON(_: NoData): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<NoData>): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },
};

const baseSellOrderPacketData: object = {
  amountDenom: "",
  amount: 0,
  price: 0,
};

export const SellOrderPacketData = {
  encode(
    message: SellOrderPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.amountDenom !== "") {
      writer.uint32(10).string(message.amountDenom);
    }
    if (message.amount !== 0) {
      writer.uint32(16).int32(message.amount);
    }
    if (message.price !== 0) {
      writer.uint32(24).int32(message.price);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SellOrderPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSellOrderPacketData } as SellOrderPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amountDenom = reader.string();
          break;
        case 2:
          message.amount = reader.int32();
          break;
        case 3:
          message.price = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SellOrderPacketData {
    const message = { ...baseSellOrderPacketData } as SellOrderPacketData;
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

  toJSON(message: SellOrderPacketData): unknown {
    const obj: any = {};
    message.amountDenom !== undefined &&
      (obj.amountDenom = message.amountDenom);
    message.amount !== undefined && (obj.amount = message.amount);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial(object: DeepPartial<SellOrderPacketData>): SellOrderPacketData {
    const message = { ...baseSellOrderPacketData } as SellOrderPacketData;
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

const baseSellOrderPacketAck: object = {};

export const SellOrderPacketAck = {
  encode(_: SellOrderPacketAck, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SellOrderPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSellOrderPacketAck } as SellOrderPacketAck;
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

  fromJSON(_: any): SellOrderPacketAck {
    const message = { ...baseSellOrderPacketAck } as SellOrderPacketAck;
    return message;
  },

  toJSON(_: SellOrderPacketAck): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<SellOrderPacketAck>): SellOrderPacketAck {
    const message = { ...baseSellOrderPacketAck } as SellOrderPacketAck;
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
