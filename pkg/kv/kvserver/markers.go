// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package kvserver

import (
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/errors"
)

// NB: don't change the string here; this will cause cross-version issues
// since this singleton is used as a marker.
var errMarkSnapshotError = errors.New("snapshot failed")

// isSnapshotError returns true iff the error indicates that a snapshot failed.
func isSnapshotError(err error) bool {
	return errors.Is(err, errMarkSnapshotError)
}

// NB: don't change the string here; this will cause cross-version issues
// since this singleton is used as a marker.
var errMarkCanRetryReplicationChangeWithUpdatedDesc = errors.New("should retry with updated descriptor")

// IsRetriableReplicationChangeError detects whether an error (which is
// assumed to have been emitted by the KV layer during a replication change
// operation) is likely to succeed when retried with an updated descriptor.
func IsRetriableReplicationChangeError(err error) bool {
	return errors.Is(err, errMarkCanRetryReplicationChangeWithUpdatedDesc) || isSnapshotError(err)
}

const (
	descChangedRangeSubsumedErrorFmt = "descriptor changed: expected %s != [actual] nil (range subsumed)"
	descChangedErrorFmt              = "descriptor changed: [expected] %s != [actual] %s"
)

func newDescChangedError(desc, actualDesc *roachpb.RangeDescriptor) error {
	if actualDesc == nil {
		return errors.Mark(errors.Newf(descChangedRangeSubsumedErrorFmt, desc), errMarkCanRetryReplicationChangeWithUpdatedDesc)
	}
	return errors.Mark(errors.Newf(descChangedErrorFmt, desc, actualDesc), errMarkCanRetryReplicationChangeWithUpdatedDesc)
}

func wrapDescChangedError(err error, desc, actualDesc *roachpb.RangeDescriptor) error {
	if actualDesc == nil {
		return errors.Mark(errors.Wrapf(err, descChangedRangeSubsumedErrorFmt, desc), errMarkCanRetryReplicationChangeWithUpdatedDesc)
	}
	return errors.Mark(errors.Wrapf(err, descChangedErrorFmt, desc, actualDesc), errMarkCanRetryReplicationChangeWithUpdatedDesc)
}

// NB: don't change the string here; this will cause cross-version issues
// since this singleton is used as a marker.
var errMarkInvalidReplicationChange = errors.New("invalid replication change")

// IsIllegalReplicationChangeError detects whether an error (assumed to have been emitted
// by a replication change) indicates that the replication change is illegal, meaning
// that it's unlikely to be handled through a retry. Examples of this are attempts to add
// a store that is already a member of the supplied descriptor. A non-example is a change
// detected in the descriptor at the KV layer relative to that supplied as input to the
// replication change, which would likely benefit from a retry.
func IsIllegalReplicationChangeError(err error) bool {
	return errors.Is(err, errMarkInvalidReplicationChange)
}
