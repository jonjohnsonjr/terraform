// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package terraform

import (
	"context"

	"github.com/hashicorp/terraform/internal/tfdiags"
)

// NodeEvalableProvider represents a provider during an "eval" walk.
// This special provider node type just initializes a provider and
// fetches its schema, without configuring it or otherwise interacting
// with it.
type NodeEvalableProvider struct {
	*NodeAbstractProvider
}

var _ GraphNodeExecutable = (*NodeEvalableProvider)(nil)

// GraphNodeExecutable
func (n *NodeEvalableProvider) Execute(_ context.Context, ectx EvalContext, op walkOperation) (diags tfdiags.Diagnostics) {
	_, err := ectx.InitProvider(n.Addr)
	return diags.Append(err)
}
