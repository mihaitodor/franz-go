// Package kmsg contains Kafka request and response types and autogenerated
// serialization and deserialization functions.
//
// This package reserves the right to add new fields to struct types as Kafka
// adds new fields over time without bumping the major API version.
package kmsg

import (
	"github.com/twmb/kgo/kbin"
)

// Request represents a type that can be requested to Kafka.
type Request interface {
	// Key returns the protocol key for this message kind.
	Key() int16
	// MaxVersion returns the maximum protocol version this message
	// supports.
	//
	// This function allows one to implement a client that chooses message
	// versions based off of the max of a message's max version in the
	// client and the broker's max supported version.
	MaxVersion() int16
	// MinVersion returns the minimum protocol version this message supports.
	MinVersion() int16
	// SetVersion sets the version to use for this request and response.
	SetVersion(int16)
	// GetVersion returns the version currently set to use for the request
	// and response.
	GetVersion() int16
	// AppendTo appends this message in wire protocol form to a slice and
	// returns the slice.
	AppendTo([]byte) []byte
	// ResponseKind returns an empty Response that is expected for
	// this message request.
	ResponseKind() Response
}

// AdminRequest represents a request that must be issued to Kafka controllers.
type AdminRequest interface {
	// IsAdminRequest returns true if the request should be issued to a
	// Kafka controller.
	IsAdminRequest() bool
	Request
}

// Response represents a type that Kafka responds with.
type Response interface {
	// ReadFrom parses all of the input slice into the response type.
	//
	// This should return an error if too much or too little data is input.
	ReadFrom([]byte) error
}

// AppendRequest appends a full message request to dst, returning the updated
// slice. This message is the full body that needs to be written to issue a
// Kafka request.
//
// clientID is optional; nil means to not send, whereas empty means the client
// id is the empty string.
func AppendRequest(
	dst []byte,
	r Request,
	correlationID int32,
	clientID *string,
) []byte {
	dst = append(dst, 0, 0, 0, 0) // reserve length
	dst = kbin.AppendInt16(dst, r.Key())
	dst = kbin.AppendInt16(dst, r.GetVersion())
	dst = kbin.AppendInt32(dst, correlationID)
	dst = kbin.AppendNullableString(dst, clientID)
	dst = r.AppendTo(dst)
	kbin.AppendInt32(dst[:0], int32(len(dst[4:])))
	return dst
}

// StringPtr is a helper to return a pointer to a string.
func StringPtr(in string) *string {
	return &in
}