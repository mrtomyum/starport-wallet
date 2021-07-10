import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "mrtomyum.wallet.dex";
export interface DexPacketData {
    noData: NoData | undefined;
    /** this line is used by starport scaffolding # ibc/packet/proto/field */
    sellOrderPacket: SellOrderPacketData | undefined;
}
export interface NoData {
}
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
export interface SellOrderPacketAck {
}
export declare const DexPacketData: {
    encode(message: DexPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DexPacketData;
    fromJSON(object: any): DexPacketData;
    toJSON(message: DexPacketData): unknown;
    fromPartial(object: DeepPartial<DexPacketData>): DexPacketData;
};
export declare const NoData: {
    encode(_: NoData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NoData;
    fromJSON(_: any): NoData;
    toJSON(_: NoData): unknown;
    fromPartial(_: DeepPartial<NoData>): NoData;
};
export declare const SellOrderPacketData: {
    encode(message: SellOrderPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SellOrderPacketData;
    fromJSON(object: any): SellOrderPacketData;
    toJSON(message: SellOrderPacketData): unknown;
    fromPartial(object: DeepPartial<SellOrderPacketData>): SellOrderPacketData;
};
export declare const SellOrderPacketAck: {
    encode(_: SellOrderPacketAck, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SellOrderPacketAck;
    fromJSON(_: any): SellOrderPacketAck;
    toJSON(_: SellOrderPacketAck): unknown;
    fromPartial(_: DeepPartial<SellOrderPacketAck>): SellOrderPacketAck;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
