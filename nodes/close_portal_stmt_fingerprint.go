// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ClosePortalStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CLOSEPORTALSTMT")

	if node.Portalname != nil {
		io.WriteString(ctx.hash, *node.Portalname)
	}
}
