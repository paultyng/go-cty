// Package ptypes provides messages and Go implementations for marshalling and 
// unmarshalling cty types and values.
//
// If the same type information is provided both at encoding and decoding time
// then values can be round-tripped without loss, except for capsule types
// which are not currently supported. Original type information is sent in each
// Value message, so it can also be inferred on the receiving end.
//
// See the Protocol Buffers documentation for importing definitions for use in 
// your own code: https://developers.google.com/protocol-buffers/docs/proto3#importing-definitions
package ptypes
