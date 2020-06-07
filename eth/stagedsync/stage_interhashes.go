package stagedsync

import (
	"fmt"

	"github.com/ledgerwatch/turbo-geth/eth/stagedsync/stages"
	"github.com/ledgerwatch/turbo-geth/ethdb"
	"github.com/ledgerwatch/turbo-geth/log"
)

//nolint:interfacer
func SpawnIntermediateHashesStage(s *StageState, stateDB ethdb.Database, _ string, _ chan struct{}) error {
	lastProcessedBlockNumber, err := stages.GetStageProgress(stateDB, stages.IntermediateHashes)
	if err != nil {
		return fmt.Errorf("IntermediateHashes: get stage progress: %w", err)
	}
	var hashedStateBlockNumber uint64
	hashedStateBlockNumber, err = stages.GetStageProgress(stateDB, stages.IntermediateHashes)
	if err != nil {
		return fmt.Errorf("IntermediateHashes: get hashed state progress: %w", err)
	}
	log.Info("Generating intermediate hashes (currently no-op)", "from", lastProcessedBlockNumber, "to", hashedStateBlockNumber)
	// TODO: Actual work goes here
	return s.DoneAndUpdate(stateDB, lastProcessedBlockNumber)
}

//nolint:interfacer
func unwindIntermediateHashesStage(u *UnwindState, s *StageState, stateDB ethdb.Database, _ string, _ chan struct{}) error {
	if err := u.Done(stateDB); err != nil {
		return fmt.Errorf("unwind IntermediateHashes: reset: %w", err)
	}
	return nil
}
