//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validatePolicyStatement_FromCedarParameters(cedarStatement *string) error {
	if cedarStatement == nil {
		return fmt.Errorf("parameter cedarStatement is required, but nil was provided")
	}

	return nil
}

func validateNewPolicyStatementParameters(props *PolicyStatementProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

