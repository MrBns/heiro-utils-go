package mirrornode

import (
	"encoding/json"
	"fmt"
	"net/http"

	hiero "github.com/hiero-ledger/hiero-sdk-go/v2/sdk"
)

// Nft
type MirrorNodeNft struct {
	AccountID         *string `json:"account_id"`
	CreatedTimestamp  *string `json:"created_timestamp"`
	DelegatingSpender *string `json:"delegating_spender"`
	Deleted           bool    `json:"deleted,omitempty"`
	Metadata          *string `json:"metadata,omitempty"`
	ModifiedTimestamp *string `json:"modified_timestamp"`
	SerialNumber      int64   `json:"serial_number,omitempty"`
	Spender           *string `json:"spender"`
	TokenID           string  `json:"token_id"`
}

func GetNftDetails(nftIdStr string) (*MirrorNodeNft, error) {
	nftId, err := hiero.NftIDFromString(nftIdStr)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(api_url(fmt.Sprintf("/api/v1/tokens/%v/nfts/%v", nftId.TokenID, nftId.SerialNumber)))
	if err != nil {
		return nil, err
	}
	err = checkApiStatus(resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var nft MirrorNodeNft
	err = json.NewDecoder(resp.Body).Decode(&nft)
	if err != nil {
		return nil, err
	}

	return &nft, nil
}
