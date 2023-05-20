//go:build !no_runtime_type_checking

package awscdkbatchalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateReason_CustomParameters(customReasonProps *CustomReason) error {
	if customReasonProps == nil {
		return fmt.Errorf("parameter customReasonProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(customReasonProps, func() string { return "parameter customReasonProps" }); err != nil {
		return err
	}

	return nil
}

