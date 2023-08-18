// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file apis/contrib.proto (syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message Contrib
 */
export class Contrib extends Message<Contrib> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string user = 2;
   */
  user = "";

  /**
   * @generated from field: google.protobuf.Timestamp date = 3;
   */
  date?: Timestamp;

  /**
   * @generated from field: string repository = 4;
   */
  repository = "";

  /**
   * @generated from oneof Contrib.contrib_type
   */
  contribType: {
    /**
     * @generated from field: string approve = 5;
     */
    value: string;
    case: "approve";
  } | {
    /**
     * @generated from field: string comment = 6;
     */
    value: string;
    case: "comment";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Contrib>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "Contrib";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "date", kind: "message", T: Timestamp },
    { no: 4, name: "repository", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "approve", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "contrib_type" },
    { no: 6, name: "comment", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "contrib_type" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Contrib {
    return new Contrib().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Contrib {
    return new Contrib().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Contrib {
    return new Contrib().fromJsonString(jsonString, options);
  }

  static equals(a: Contrib | PlainMessage<Contrib> | undefined, b: Contrib | PlainMessage<Contrib> | undefined): boolean {
    return proto3.util.equals(Contrib, a, b);
  }
}

/**
 * @generated from message ListContribsRequest
 */
export class ListContribsRequest extends Message<ListContribsRequest> {
  /**
   * @generated from field: string repository = 1;
   */
  repository = "";

  /**
   * @generated from field: int32 page_size = 2;
   */
  pageSize = 0;

  /**
   * @generated from field: string page_token = 3;
   */
  pageToken = "";

  constructor(data?: PartialMessage<ListContribsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ListContribsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "repository", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListContribsRequest {
    return new ListContribsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListContribsRequest {
    return new ListContribsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListContribsRequest {
    return new ListContribsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListContribsRequest | PlainMessage<ListContribsRequest> | undefined, b: ListContribsRequest | PlainMessage<ListContribsRequest> | undefined): boolean {
    return proto3.util.equals(ListContribsRequest, a, b);
  }
}

/**
 * @generated from message ListContribsResponse
 */
export class ListContribsResponse extends Message<ListContribsResponse> {
  /**
   * @generated from field: repeated Contrib contribs = 1;
   */
  contribs: Contrib[] = [];

  /**
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken = "";

  constructor(data?: PartialMessage<ListContribsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ListContribsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "contribs", kind: "message", T: Contrib, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListContribsResponse {
    return new ListContribsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListContribsResponse {
    return new ListContribsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListContribsResponse {
    return new ListContribsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListContribsResponse | PlainMessage<ListContribsResponse> | undefined, b: ListContribsResponse | PlainMessage<ListContribsResponse> | undefined): boolean {
    return proto3.util.equals(ListContribsResponse, a, b);
  }
}

