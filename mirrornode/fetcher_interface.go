package mirrornode

import "context"

type NftFetcherInterface interface {
	GetNftInfo(ctx context.Context, nftId string) (any, error)
}

type TokenFetcherInterface interface {
}

type TopicFetcherInterface interface {
}

type AccountFetcherInferface interface {
}

type TransactionFetcherInferface interface {
}

type SmartContractFetcherInterface interface {
}
