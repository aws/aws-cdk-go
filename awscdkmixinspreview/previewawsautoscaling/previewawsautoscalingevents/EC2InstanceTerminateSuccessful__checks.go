//go:build !no_runtime_type_checking

package previewawsautoscalingevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEC2InstanceTerminateSuccessful_Ec2InstanceTerminateSuccessfulPatternParameters(options *EC2InstanceTerminateSuccessful_EC2InstanceTerminateSuccessfulProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

