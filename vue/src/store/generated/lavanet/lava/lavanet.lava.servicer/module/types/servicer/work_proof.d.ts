import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "lavanet.lava.servicer";
export interface WorkProof {
    data: string;
}
export declare const WorkProof: {
    encode(message: WorkProof, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): WorkProof;
    fromJSON(object: any): WorkProof;
    toJSON(message: WorkProof): unknown;
    fromPartial(object: DeepPartial<WorkProof>): WorkProof;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};