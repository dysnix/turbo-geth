package stagedsync

import (
	"encoding/binary"

	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
	"github.com/ledgerwatch/turbo-geth/common/etl"
	"github.com/ledgerwatch/turbo-geth/core/rawdb"
	"github.com/ledgerwatch/turbo-geth/ethdb"
)

func extractHeaders(k []byte, v []byte, next etl.ExtractNextFunc) error {
	// We only want to extract entries composed by Block Number + Header Hash
	if len(k) != 40 {
		return nil
	}
	return next(k, common.CopyBytes(k[8:]), common.CopyBytes(k[:8]))
}

func SpawnBlockHashStage(s *StageState, stateDB ethdb.Database, quit <-chan struct{}) error {
	headHash := rawdb.ReadHeadHeaderHash(stateDB)
	headNumber := rawdb.ReadHeaderNumber(stateDB, headHash)
	if s.BlockNumber == *headNumber {
		s.Done()
		return nil
	}

	startKey := make([]byte, 8)
	binary.BigEndian.PutUint64(startKey, s.BlockNumber)
	endKey := dbutils.HeaderKey(*headNumber, headHash) // Make sure we stop at head

	if err := etl.Transform(
		stateDB,
		dbutils.HeaderPrefix,
		dbutils.HeaderNumberPrefix,
		".",
		extractHeaders,
		etl.IdentityLoadFunc,
		etl.TransformArgs{
			ExtractStartKey: startKey,
			ExtractEndKey:   endKey,
			Quit:            quit,
		},
	); err != nil {
		return err
	}
	return s.DoneAndUpdate(stateDB, *headNumber)
}
