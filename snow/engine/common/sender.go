// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/ava-labs/gecko/ids"
)

// Sender defines how a consensus engine sends messages and requests to other
// validators
type Sender interface {
	FrontierSender
	AcceptedSender
	FetchSender
	QuerySender
}

// FrontierSender defines how a consensus engine sends frontier messages to
// other validators
type FrontierSender interface {
	// GetAcceptedFrontier requests that every validator in [validatorIDs] sends
	// an AcceptedFrontier message.
	GetAcceptedFrontier(validatorIDs ids.ShortSet, requestID uint32)

	// AcceptedFrontier responds to a AcceptedFrontier message with this
	// engine's current accepted frontier.
	AcceptedFrontier(validatorID ids.ShortID, requestID uint32, containerIDs ids.Set)
}

// AcceptedSender defines how a consensus engine sends messages pertaining to
// accepted containers
type AcceptedSender interface {
	// GetAccepted requests that every validator in [validatorIDs] sends an
	// Accepted message with all the IDs in [containerIDs] that the validator
	// thinks is accepted.
	GetAccepted(validatorIDs ids.ShortSet, requestID uint32, containerIDs ids.Set)

	// Accepted responds to a GetAccepted message with a set of IDs of
	// containers that are accepted.
	Accepted(validatorID ids.ShortID, requestID uint32, containerIDs ids.Set)
}

// FetchSender defines how a consensus engine sends retrieval messages to other
// validators
type FetchSender interface {
	// Request a container from a validator.
	// Request that the specified validator send the specified container
	// to this validator
	Get(validatorID ids.ShortID, requestID uint32, containerID ids.ID)

	// Tell the specified validator that the container whose ID is <containerID>
	// has body <container>
	Put(validatorID ids.ShortID, requestID uint32, containerID ids.ID, container []byte)
}

// QuerySender defines how a consensus engine sends query messages to other
// validators
type QuerySender interface {
	// Request from the specified validators their preferred frontier, given the
	// existence of the specified container.
	// This is the same as PullQuery, except that this message includes not only
	// the ID of the container but also its body.
	PushQuery(validatorIDs ids.ShortSet, requestID uint32, containerID ids.ID, container []byte)

	// Request from the specified validators their preferred frontier, given the
	// existence of the specified container.
	PullQuery(validatorIDs ids.ShortSet, requestID uint32, containerID ids.ID)

	// Chits sends chits to the specified validator
	Chits(validatorID ids.ShortID, requestID uint32, votes ids.Set)
}