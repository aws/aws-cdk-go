//go:build !no_runtime_type_checking

package awsmediapackagev2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateInputConfiguration_CmafParameters(props *CmafInputProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

