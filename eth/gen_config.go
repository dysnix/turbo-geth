// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eth

import (
	"math/big"
	"time"

	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/consensus/ethash"
	"github.com/ledgerwatch/turbo-geth/core"
	"github.com/ledgerwatch/turbo-geth/eth/downloader"
	"github.com/ledgerwatch/turbo-geth/eth/gasprice"
	"github.com/ledgerwatch/turbo-geth/miner"
	"github.com/ledgerwatch/turbo-geth/params"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkID               uint64
		SyncMode                downloader.SyncMode
		DiscoveryURLs           []string
		Pruning               bool
		NoPrefetch              bool
		TxLookupLimit           uint64                 `toml:",omitempty"`
		Whitelist               map[uint64]common.Hash `toml:"-"`
		LightIngress            int                    `toml:",omitempty"`
		LightEgress             int                    `toml:",omitempty"`
		StorageMode             string
		ArchiveSyncInterval     int
		LightServ               int `toml:",omitempty"`
		LightPeers              int `toml:",omitempty"`
		OnlyAnnounce            bool
		SkipBcVersionCheck      bool `toml:"-"`
		DatabaseHandles         int  `toml:"-"`
		DatabaseCache           int
		DatabaseFreezer         string
		TrieCleanCache          int
		TrieDirtyCache          int
		TrieTimeout             time.Duration
		Miner                   miner.Config
		Ethash                  ethash.Config
		TxPool                  core.TxPoolConfig
		GPO                     gasprice.Config
		EnablePreimageRecording bool
		DocRoot                 string `toml:"-"`
		EWASMInterpreter        string
		EVMInterpreter          string
		RPCGasCap               *big.Int                       `toml:",omitempty"`
		Checkpoint              *params.TrustedCheckpoint      `toml:",omitempty"`
		CheckpointOracle        *params.CheckpointOracleConfig `toml:",omitempty"`
		OverrideIstanbul        *big.Int                       `toml:",omitempty"`
		OverrideMuirGlacier     *big.Int                       `toml:",omitempty"`
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkID = c.NetworkID
	enc.SyncMode = c.SyncMode
	enc.DiscoveryURLs = c.DiscoveryURLs
	enc.Pruning = c.Pruning
	enc.NoPrefetch = c.NoPrefetch
	enc.TxLookupLimit = c.TxLookupLimit
	enc.Whitelist = c.Whitelist
	enc.StorageMode = c.StorageMode.ToString()
	enc.ArchiveSyncInterval = c.ArchiveSyncInterval
	enc.LightServ = c.LightServ
	enc.LightIngress = c.LightIngress
	enc.LightEgress = c.LightEgress
	enc.LightPeers = c.LightPeers
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.DatabaseFreezer = c.DatabaseFreezer
	enc.TrieCleanCache = c.TrieCleanCache
	enc.TrieDirtyCache = c.TrieDirtyCache
	enc.TrieTimeout = c.TrieTimeout
	enc.Miner = c.Miner
	enc.Ethash = c.Ethash
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.DocRoot = c.DocRoot
	enc.EWASMInterpreter = c.EWASMInterpreter
	enc.EVMInterpreter = c.EVMInterpreter
	enc.RPCGasCap = c.RPCGasCap
	enc.Checkpoint = c.Checkpoint
	enc.CheckpointOracle = c.CheckpointOracle
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkID               *uint64
		SyncMode                *downloader.SyncMode
		DiscoveryURLs           []string
		Pruning               *bool
		NoPrefetch              *bool
		TxLookupLimit           *uint64                `toml:",omitempty"`
		Whitelist               map[uint64]common.Hash `toml:"-"`
		LightIngress            *int                   `toml:",omitempty"`
		LightEgress             *int                   `toml:",omitempty"`
		Mode                    *string
		ArchiveSyncInterval     *int
		LightServ               *int `toml:",omitempty"`
		LightPeers              *int `toml:",omitempty"`
		OnlyAnnounce            *bool
		SkipBcVersionCheck      *bool `toml:"-"`
		DatabaseHandles         *int  `toml:"-"`
		DatabaseCache           *int
		DatabaseFreezer         *string
		TrieCleanCache          *int
		TrieDirtyCache          *int
		TrieTimeout             *time.Duration
		Miner                   *miner.Config
		Ethash                  *ethash.Config
		TxPool                  *core.TxPoolConfig
		GPO                     *gasprice.Config
		EnablePreimageRecording *bool
		DocRoot                 *string `toml:"-"`
		EWASMInterpreter        *string
		EVMInterpreter          *string
		RPCGasCap               *big.Int                       `toml:",omitempty"`
		Checkpoint              *params.TrustedCheckpoint      `toml:",omitempty"`
		CheckpointOracle        *params.CheckpointOracleConfig `toml:",omitempty"`
		OverrideIstanbul        *big.Int                       `toml:",omitempty"`
		OverrideMuirGlacier     *big.Int                       `toml:",omitempty"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkID != nil {
		c.NetworkID = *dec.NetworkID
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.DiscoveryURLs != nil {
		c.DiscoveryURLs = dec.DiscoveryURLs
	}
	if dec.Pruning != nil {
		c.Pruning = *dec.Pruning
	}
	if dec.NoPrefetch != nil {
		c.NoPrefetch = *dec.NoPrefetch
	}
	if dec.TxLookupLimit != nil {
		c.TxLookupLimit = *dec.TxLookupLimit
	}
	if dec.Whitelist != nil {
		c.Whitelist = dec.Whitelist
	}
	if dec.Mode != nil {
		mode, err := StorageModeFromString(*dec.Mode)
		if err != nil {
			return err
		}
		c.StorageMode = mode
	}
	if dec.ArchiveSyncInterval != nil {
		c.ArchiveSyncInterval = *dec.ArchiveSyncInterval
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightIngress != nil {
		c.LightIngress = *dec.LightIngress
	}
	if dec.LightEgress != nil {
		c.LightEgress = *dec.LightEgress
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.DatabaseFreezer != nil {
		c.DatabaseFreezer = *dec.DatabaseFreezer
	}
	if dec.TrieCleanCache != nil {
		c.TrieCleanCache = *dec.TrieCleanCache
	}
	if dec.TrieDirtyCache != nil {
		c.TrieDirtyCache = *dec.TrieDirtyCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.Ethash != nil {
		c.Ethash = *dec.Ethash
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.EWASMInterpreter != nil {
		c.EWASMInterpreter = *dec.EWASMInterpreter
	}
	if dec.EVMInterpreter != nil {
		c.EVMInterpreter = *dec.EVMInterpreter
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = dec.RPCGasCap
	}
	if dec.Checkpoint != nil {
		c.Checkpoint = dec.Checkpoint
	}
	if dec.CheckpointOracle != nil {
		c.CheckpointOracle = dec.CheckpointOracle
	}
	return nil
}
