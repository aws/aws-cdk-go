//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateInitUser_FromNameParameters(userName *string, options *InitUserOptions) error {
	if userName == nil {
		return fmt.Errorf("parameter userName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewInitUserParameters(userName *string, userOptions *InitUserOptions) error {
	if userName == nil {
		return fmt.Errorf("parameter userName is required, but nil was provided")
	}

	if userOptions == nil {
		return fmt.Errorf("parameter userOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(userOptions, func() string { return "parameter userOptions" }); err != nil {
		return err
	}

	return nil
}

