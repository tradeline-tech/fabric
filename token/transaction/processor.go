/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package transaction

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/tradeline-tech/fabric/common/flogging"
	"github.com/tradeline-tech/fabric/core/ledger"
	"github.com/tradeline-tech/fabric/protos/common"
)

var logger = flogging.MustGetLogger("fabtoken-processor")

// Processor implements the interface 'github.com/hyperledger/fabric/core/ledger/customtx/Processor'
// for FabToken transactions
type Processor struct {
	TMSManager TMSManager
}

func (p *Processor) GenerateSimulationResults(txEnv *common.Envelope, simulator ledger.TxSimulator, initializingLedger bool) error {
	// Extract channel header and token transaction
	ch, ttx, ci, err := UnmarshalTokenTransaction(txEnv.Payload)
	if err != nil {
		return errors.WithMessage(err, "failed unmarshalling token transaction")
	}

	// Get a TMSTxProcessor that corresponds to the channel
	txProcessor, err := p.TMSManager.GetTxProcessor(ch.ChannelId)
	if err != nil {
		return errors.WithMessage(err, "failed getting committer")
	}

	// Extract the read dependencies and ledger updates associated to the transaction using simulator
	err = txProcessor.ProcessTx(ch.TxId, ci, ttx, simulator)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed committing transaction for channel %s", ch.ChannelId))
	}

	return err
}
