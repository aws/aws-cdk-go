//go:build !no_runtime_type_checking

package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewOrcOutputFormatParameters(props *OrcOutputFormatProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

