/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package metadata_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	common "github.com/tradeline-tech/fabric/common/metadata"
	"github.com/tradeline-tech/fabric/orderer/common/metadata"
)

func TestGetVersionInfo(t *testing.T) {
	expected := fmt.Sprintf(
		"%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n OS/Arch: %s\n",
		metadata.ProgramName, common.Version,
		common.CommitSHA,
		runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	)
	assert.Equal(t, expected, metadata.GetVersionInfo())
}
