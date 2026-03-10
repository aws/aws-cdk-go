//go:build !no_runtime_type_checking

package previewawsautoscalingevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEC2InstanceTerminateLifecycleAction_Ec2InstanceTerminateLifecycleActionPatternParameters(options *EC2InstanceTerminateLifecycleAction_EC2InstanceTerminateLifecycleActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

