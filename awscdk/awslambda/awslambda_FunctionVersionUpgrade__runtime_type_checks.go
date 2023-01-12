//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FunctionVersionUpgrade) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateNewFunctionVersionUpgradeParameters(featureFlag *string) error {
	if featureFlag == nil {
		return fmt.Errorf("parameter featureFlag is required, but nil was provided")
	}

	return nil
}

