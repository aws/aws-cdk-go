//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateArn_ExtractResourceNameParameters(arn *string, resourceType *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	if resourceType == nil {
		return fmt.Errorf("parameter resourceType is required, but nil was provided")
	}

	return nil
}

func validateArn_FormatParameters(components *ArnComponents) error {
	if components == nil {
		return fmt.Errorf("parameter components is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(components, func() string { return "parameter components" }); err != nil {
		return err
	}

	return nil
}

func validateArn_SplitParameters(arn *string, arnFormat ArnFormat) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	if arnFormat == "" {
		return fmt.Errorf("parameter arnFormat is required, but nil was provided")
	}

	return nil
}

