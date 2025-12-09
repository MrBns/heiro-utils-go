package mirrornode

type MirroNodeNetworksType string

const (
	Mainnet    MirroNodeNetworksType = "mainnet"
	TestNet    MirroNodeNetworksType = "testnet"
	PreviewNet MirroNodeNetworksType = "previewnet"
	LocalNet   MirroNodeNetworksType = "localnet"
)

type MirrorNodeFetcher struct {
	Network MirroNodeNetworksType
}

/*
Create Mirrornode Fetcher Instance.
*/
func NewMirrorNodeFetcher(network MirroNodeNetworksType) *MirrorNodeFetcher {

	return &MirrorNodeFetcher{
		Network: network,
	}
}
