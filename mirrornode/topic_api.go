package mirrornode

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	bns "github.com/mrbns/gokit/utility"
)

// TopicMessage
type TopicMessage struct {
	ChunkInfo          *ChunkInfo `json:"chunk_info"`
	ConsensusTimestamp string     `json:"consensus_timestamp"`
	Message            string     `json:"message"`
	PayerAccountID     *string    `json:"payer_account_id"`
	RunningHash        string     `json:"running_hash"`
	RunningHashVersion int64      `json:"running_hash_version"`
	SequenceNumber     int64      `json:"sequence_number"`
	TopicID            *string    `json:"topic_id"`
}

type ChunkInfo struct {
	InitialTransactionID *TransactionID `json:"initial_transaction_id,omitempty"`
	Number               *int64         `json:"number,omitempty"`
	Total                *int64         `json:"total,omitempty"`
	Scheduled            bool           `json:"scheduled"`
	Nonce                *int64         `json:"nonce"`
}

// TransactionId
type TransactionID struct {
	AccountID             *string `json:"account_id"`
	Nonce                 *int64  `json:"nonce"`
	Scheduled             *bool   `json:"scheduled"`
	TransactionValidStart *string `json:"transaction_valid_start,omitempty"`
}

type GetTopicMessagesByIdOptions struct {
	Encoding       string
	Limit          int64
	Sequencenumber int64

	// The order in which items are listed either  "asc" or "desc"
	Order string

	// The consensus timestamp as a Unix timestamp in seconds.nanoseconds format with an
	// Timestamp []string
	TopicId string
}

// TopicMessagesResponse
type GetTopicMessagesResponse struct {
	Links    *Links         `json:"links,omitempty"`
	Messages []TopicMessage `json:"messages,omitempty"`
}

func GetTopicMessagesById(options GetTopicMessagesByIdOptions) (*GetTopicMessagesResponse, error) {
	if options.TopicId == "" {
		return nil, fmt.Errorf("please provide a valid topic id")
	}

	searchParams := url.Values{}
	searchParams.Add("encoding", bns.Ternary(options.Encoding != "", options.Encoding, "base64"))
	searchParams.Add("limit", strconv.FormatInt(bns.Ternary(options.Limit > 0, options.Limit, 50), 10))

	if options.Order == "asc" || options.Order == "desc" {
		searchParams.Add("order", options.Order)
	}

	if options.Sequencenumber > 0 {
		searchParams.Add("sequencenumber", strconv.FormatInt(options.Sequencenumber, 10))
	}

	searchParamsString := searchParams.Encode()
	resp, err := http.Get(api_url(fmt.Sprintf("/api/v1/topics/%v/messages?%v", options.TopicId, searchParamsString)))
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	err = checkApiStatus(resp)
	if err != nil {
		return nil, err
	}

	topicMessages := new(GetTopicMessagesResponse)
	err = json.NewDecoder(resp.Body).Decode(topicMessages)
	if err != nil {
		return nil, err
	}

	return topicMessages, nil
}
