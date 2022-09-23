//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awscodedeploy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateLambdaDeploymentConfig_ImportParameters(_scope constructs.Construct, _id *string, props *LambdaDeploymentConfigImportProps) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

