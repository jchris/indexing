// constants used by secondary indexing system.

package common

import (
	"errors"
)

// error codes

// ErrorEmptyN1QLExpression
var ErrorEmptyN1QLExpression = errors.New("secondary.emptyN1QLExpression")

// ErrorUnexpectedPayload
var ErrorUnexpectedPayload = errors.New("secondary.unexpectedPayload")

// ErrorClosed
var ErrorClosed = errors.New("genServer.closed")

// ErrorChannelFull
var ErrorChannelFull = errors.New("secondary.channelFull")

// ErrorNotMyVbucket
var ErrorNotMyVbucket = errors.New("secondary.notMyVbucket")

// ErrorInvalidRequest
var ErrorInvalidRequest = errors.New("secondary.invalidRequest")

// ErrorNotFound
var ErrorNotFound = errors.New("secondary.notFound")

// ProtobufDataPathMajorNum major version number for mutation data path.
var ProtobufDataPathMajorNum byte // = 0

// ProtobufDataPathMinorNum minor version number for mutation data path.
var ProtobufDataPathMinorNum byte = 1

// ErrScanTimedOut from indexer
var ErrScanTimedOut = errors.New("Index scan timed out")

// ErrClientCancel when query client cancels an ongoing scan request.
var ErrClientCancel = errors.New("Client requested cancel")
