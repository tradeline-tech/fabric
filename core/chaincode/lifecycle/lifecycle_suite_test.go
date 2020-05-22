/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package lifecycle_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tradeline-tech/fabric/core/chaincode/lifecycle"
	"github.com/tradeline-tech/fabric/core/chaincode/shim"
)

//go:generate counterfeiter -o mock/chaincode_stub.go --fake-name ChaincodeStub . chaincodeStub
type chaincodeStub interface {
	shim.ChaincodeStubInterface
}

//go:generate counterfeiter -o mock/chaincode_store.go --fake-name ChaincodeStore . chaincodeStore
type chaincodeStore interface {
	lifecycle.ChaincodeStore
}

//go:generate counterfeiter -o mock/package_parser.go --fake-name PackageParser . packageParser
type packageParser interface {
	lifecycle.PackageParser
}

//go:generate counterfeiter -o mock/scc_functions.go --fake-name SCCFunctions . sccFunctions
type sccFunctions interface {
	lifecycle.SCCFunctions
}

func TestLifecycle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lifecycle Suite")
}
