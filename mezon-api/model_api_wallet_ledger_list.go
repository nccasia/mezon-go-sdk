/*
 * Mezon API v2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Contact: hello@heroiclabs.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiWalletLedgerList struct {
	WalletLedger []ApiWalletLedger `json:"walletLedger,omitempty"`
	PrevCursor string `json:"prevCursor,omitempty"`
	NextCursor string `json:"nextCursor,omitempty"`
}