//go:build !no_runtime_type_checking

package awsroute53

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCidrRoutingConfig_CreateParameters(props *CidrRoutingConfigProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateCidrRoutingConfig_WithDefaultLocationNameParameters(collectionId *string) error {
	if collectionId == nil {
		return fmt.Errorf("parameter collectionId is required, but nil was provided")
	}

	return nil
}

