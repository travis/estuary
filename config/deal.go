package config

import (
	"github.com/application-research/filclient"
	"github.com/libp2p/go-libp2p-core/protocol"
)

const (
	cliDealProtocolv110 = "v110"
	cliDealProtocolv120 = "v120"
)

var DealProtocolsVersionsMap = map[string]protocol.ID{
	cliDealProtocolv110: filclient.DealProtocolv110,
	cliDealProtocolv120: filclient.DealProtocolv120,
}

type Deal struct {
	FailOnTransferFailure        bool                 `json:"fail_on_transfer_failure"`
	Disable                      bool                 `json:"disable"`
	Verified                     bool                 `json:"verified"`
	EnabledDealProtocolsVersions map[protocol.ID]bool `json:"enabled_deal_protocol_versions"`
}
