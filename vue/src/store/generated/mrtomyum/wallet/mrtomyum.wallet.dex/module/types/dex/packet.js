/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "mrtomyum.wallet.dex";
const baseDexPacketData = {};
export const DexPacketData = {
    encode(message, writer = Writer.create()) {
        if (message.noData !== undefined) {
            NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
        }
        if (message.sellOrderPacket !== undefined) {
            SellOrderPacketData.encode(message.sellOrderPacket, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseDexPacketData };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.noData = NoData.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.sellOrderPacket = SellOrderPacketData.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseDexPacketData };
        if (object.noData !== undefined && object.noData !== null) {
            message.noData = NoData.fromJSON(object.noData);
        }
        else {
            message.noData = undefined;
        }
        if (object.sellOrderPacket !== undefined &&
            object.sellOrderPacket !== null) {
            message.sellOrderPacket = SellOrderPacketData.fromJSON(object.sellOrderPacket);
        }
        else {
            message.sellOrderPacket = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.noData !== undefined &&
            (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
        message.sellOrderPacket !== undefined &&
            (obj.sellOrderPacket = message.sellOrderPacket
                ? SellOrderPacketData.toJSON(message.sellOrderPacket)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseDexPacketData };
        if (object.noData !== undefined && object.noData !== null) {
            message.noData = NoData.fromPartial(object.noData);
        }
        else {
            message.noData = undefined;
        }
        if (object.sellOrderPacket !== undefined &&
            object.sellOrderPacket !== null) {
            message.sellOrderPacket = SellOrderPacketData.fromPartial(object.sellOrderPacket);
        }
        else {
            message.sellOrderPacket = undefined;
        }
        return message;
    },
};
const baseNoData = {};
export const NoData = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseNoData };
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
    fromJSON(_) {
        const message = { ...baseNoData };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseNoData };
        return message;
    },
};
const baseSellOrderPacketData = {
    amountDenom: "",
    amount: 0,
    price: 0,
};
export const SellOrderPacketData = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSellOrderPacketData };
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
    fromJSON(object) {
        const message = { ...baseSellOrderPacketData };
        if (object.amountDenom !== undefined && object.amountDenom !== null) {
            message.amountDenom = String(object.amountDenom);
        }
        else {
            message.amountDenom = "";
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = Number(object.amount);
        }
        else {
            message.amount = 0;
        }
        if (object.price !== undefined && object.price !== null) {
            message.price = Number(object.price);
        }
        else {
            message.price = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.amountDenom !== undefined &&
            (obj.amountDenom = message.amountDenom);
        message.amount !== undefined && (obj.amount = message.amount);
        message.price !== undefined && (obj.price = message.price);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSellOrderPacketData };
        if (object.amountDenom !== undefined && object.amountDenom !== null) {
            message.amountDenom = object.amountDenom;
        }
        else {
            message.amountDenom = "";
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = 0;
        }
        if (object.price !== undefined && object.price !== null) {
            message.price = object.price;
        }
        else {
            message.price = 0;
        }
        return message;
    },
};
const baseSellOrderPacketAck = {};
export const SellOrderPacketAck = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSellOrderPacketAck };
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
    fromJSON(_) {
        const message = { ...baseSellOrderPacketAck };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseSellOrderPacketAck };
        return message;
    },
};
