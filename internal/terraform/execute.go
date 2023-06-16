// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package terraform

import (
	"context"

	"github.com/hashicorp/terraform/internal/tfdiags"
)

// GraphNodeExecutable is the interface that graph nodes must implement to
// enable execution.
type GraphNodeExecutable interface {
	Execute(context.Context, EvalContext, walkOperation) tfdiags.Diagnostics
}
