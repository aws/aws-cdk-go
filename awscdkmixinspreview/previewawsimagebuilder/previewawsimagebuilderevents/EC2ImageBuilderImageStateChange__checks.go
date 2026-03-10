//go:build !no_runtime_type_checking

package previewawsimagebuilderevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEC2ImageBuilderImageStateChange_Ec2ImageBuilderImageStateChangePatternParameters(options *EC2ImageBuilderImageStateChange_EC2ImageBuilderImageStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

