// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Var) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "VAR")
	// Intentionally ignoring node.Location for fingerprinting
}
