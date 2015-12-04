// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateForeignTableStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATEFOREIGNTABLESTMT")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Servername != nil {
		io.WriteString(ctx.hash, *node.Servername)
	}
}
