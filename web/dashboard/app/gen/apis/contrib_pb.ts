// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file apis/contrib.proto (syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message Review
 */
export class Review extends Message<Review> {
  /**
   * @generated from field: string user = 2;
   */
  user = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 3;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string state = 4;
   */
  state = "";

  constructor(data?: PartialMessage<Review>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "Review";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 2, name: "user", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "created_at", kind: "message", T: Timestamp },
    { no: 4, name: "state", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Review {
    return new Review().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Review {
    return new Review().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Review {
    return new Review().fromJsonString(jsonString, options);
  }

  static equals(a: Review | PlainMessage<Review> | undefined, b: Review | PlainMessage<Review> | undefined): boolean {
    return proto3.util.equals(Review, a, b);
  }
}

/**
 * @generated from message Pull
 */
export class Pull extends Message<Pull> {
  /**
   * @generated from field: int32 number = 1;
   */
  number = 0;

  /**
   * @generated from field: string user = 2;
   */
  user = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 3;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string state = 4;
   */
  state = "";

  /**
   * @generated from field: repeated Review reviews = 5;
   */
  reviews: Review[] = [];

  constructor(data?: PartialMessage<Pull>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "Pull";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "number", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "user", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "created_at", kind: "message", T: Timestamp },
    { no: 4, name: "state", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "reviews", kind: "message", T: Review, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pull {
    return new Pull().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pull {
    return new Pull().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pull {
    return new Pull().fromJsonString(jsonString, options);
  }

  static equals(a: Pull | PlainMessage<Pull> | undefined, b: Pull | PlainMessage<Pull> | undefined): boolean {
    return proto3.util.equals(Pull, a, b);
  }
}

/**
 * @generated from message Repository
 */
export class Repository extends Message<Repository> {
  /**
   * @generated from field: string org = 1;
   */
  org = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: repeated Pull pulls = 3;
   */
  pulls: Pull[] = [];

  constructor(data?: PartialMessage<Repository>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "Repository";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "org", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "pulls", kind: "message", T: Pull, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Repository {
    return new Repository().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Repository {
    return new Repository().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Repository {
    return new Repository().fromJsonString(jsonString, options);
  }

  static equals(a: Repository | PlainMessage<Repository> | undefined, b: Repository | PlainMessage<Repository> | undefined): boolean {
    return proto3.util.equals(Repository, a, b);
  }
}

