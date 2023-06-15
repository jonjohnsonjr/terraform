// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefile

import (
	"bytes"

	"github.com/hashicorp/terraform/internal/states"
)

// StatesMarshalEqual returns true if and only if the two given states have
// an identical (byte-for-byte) statefile representation.
//
// This function compares only the portions of the state that are persisted
// in state files, so for example it will not return false if the only
// differences between the two states are local values or descendent module
// outputs.
func StatesMarshalEqual(a, b *states.State) bool {
	var aBuf bytes.Buffer

	// nil states are not valid states, and so they can never martial equal.
	if a == nil || b == nil {
		return false
	}

	// We write here some temporary files that have no header information
	// populated, thus ensuring that we're only comparing the state itself
	// and not any metadata.
	err := WriteStateV4(&File{State: a}, &aBuf)
	if err != nil {
		// Should never happen, because we're writing to an in-memory buffer
		panic(err)
	}

	buf := make([]byte, 0, aBuf.Cap())
	bBuf := bytes.NewBuffer(buf)
	err = WriteStateV4(&File{State: b}, bBuf)
	if err != nil {
		// Should never happen, because we're writing to an in-memory buffer
		panic(err)
	}

	return bytes.Equal(aBuf.Bytes(), bBuf.Bytes())
}
