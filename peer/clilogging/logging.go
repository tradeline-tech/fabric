/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package clilogging

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tradeline-tech/fabric/common/flogging"
	"github.com/tradeline-tech/fabric/peer/common"
)

const (
	loggingFuncName = "logging"
	loggingCmdDes   = "Logging configuration: getlevel|setlevel|getlogspec|setlogspec|revertlevels."
)

var logger = flogging.MustGetLogger("cli.logging")

// Cmd returns the cobra command for Logging
func Cmd(cf *LoggingCmdFactory) *cobra.Command {
	loggingCmd.AddCommand(getLevelCmd(cf))
	loggingCmd.AddCommand(setLevelCmd(cf))
	loggingCmd.AddCommand(revertLevelsCmd(cf))
	loggingCmd.AddCommand(getLogSpecCmd(cf))
	loggingCmd.AddCommand(setLogSpecCmd(cf))

	return loggingCmd
}

var loggingCmd = &cobra.Command{
	Use:              loggingFuncName,
	Short:            fmt.Sprint(loggingCmdDes),
	Long:             fmt.Sprint(loggingCmdDes),
	PersistentPreRun: common.InitCmd,
}
