/*
 * NFD Management Service
 *
 * Service for querying and managing NFDs
 *
 * API version: 1.0
 * Contact: feedback@txnlab.dev
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SendToVaultRequestBody struct {
	// Base amount (in base units of specified asset - so decimals must be considered) of asset to send.   If multiple assets specified, amount is ignored and ALL of each are sent
	Amount int64 `json:"amount"`
	// Algorand ASA IDs to transfer (and opt-in inside vault if necessary) - use asset 0 to send ALGO.  Specifying multiple assets means ALL of each are sent and amount is ignored. 13 is max assets that can be specified if they're being sent (2 for MBR payments, 2 for opt-in txns (8+4 asset opt-ins), 12 asset transfers).  If opt-in only then 64 is maximum (1 MBR per 8 assets, 8 assets per txn * 8 txns)
	Assets []int64 `json:"assets"`
	// Optional note to include in asset send transaction
	Note string `json:"note,omitempty"`
	// Whether to only opt-in to the asset, instead of including asset transfer txn
	OptInOnly bool   `json:"optInOnly"`
	Sender    string `json:"sender"`
}
