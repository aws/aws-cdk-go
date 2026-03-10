//go:build !no_runtime_type_checking

package previewawsimagebuilderevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEC2ImageBuilderCVEDetected_Ec2ImageBuilderCVEDetectedPatternParameters(options *EC2ImageBuilderCVEDetected_EC2ImageBuilderCVEDetectedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

