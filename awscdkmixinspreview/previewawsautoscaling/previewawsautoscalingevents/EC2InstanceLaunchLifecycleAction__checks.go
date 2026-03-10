//go:build !no_runtime_type_checking

package previewawsautoscalingevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEC2InstanceLaunchLifecycleAction_Ec2InstanceLaunchLifecycleActionPatternParameters(options *EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

