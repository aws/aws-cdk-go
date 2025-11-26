//go:build !no_runtime_type_checking

package previewawsautoscalingevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsautoscaling"
)

func (a *jsiiProxy_AutoScalingGroupEvents) validateEC2InstanceLaunchLifecycleActionPatternParameters(options *AutoScalingGroupEvents_EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AutoScalingGroupEvents) validateEC2InstanceLaunchSuccessfulPatternParameters(options *AutoScalingGroupEvents_EC2InstanceLaunchSuccessful_EC2InstanceLaunchSuccessfulProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AutoScalingGroupEvents) validateEC2InstanceLaunchUnsuccessfulPatternParameters(options *AutoScalingGroupEvents_EC2InstanceLaunchUnsuccessful_EC2InstanceLaunchUnsuccessfulProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AutoScalingGroupEvents) validateEC2InstanceTerminateLifecycleActionPatternParameters(options *AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction_EC2InstanceTerminateLifecycleActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AutoScalingGroupEvents) validateEC2InstanceTerminateSuccessfulPatternParameters(options *AutoScalingGroupEvents_EC2InstanceTerminateSuccessful_EC2InstanceTerminateSuccessfulProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AutoScalingGroupEvents) validateEC2InstanceTerminateUnsuccessfulPatternParameters(options *AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful_EC2InstanceTerminateUnsuccessfulProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAutoScalingGroupEvents_FromAutoScalingGroupParameters(autoScalingGroupRef interfacesawsautoscaling.IAutoScalingGroupRef) error {
	if autoScalingGroupRef == nil {
		return fmt.Errorf("parameter autoScalingGroupRef is required, but nil was provided")
	}

	return nil
}

