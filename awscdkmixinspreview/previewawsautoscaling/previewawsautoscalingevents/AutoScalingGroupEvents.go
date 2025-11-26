package previewawsautoscalingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsautoscaling"
)

// EventBridge event patterns for AutoScalingGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroupRef IAutoScalingGroupRef
//
//   autoScalingGroupEvents := awscdkmixinspreview.Events.AutoScalingGroupEvents_FromAutoScalingGroup(autoScalingGroupRef)
//
// Experimental.
type AutoScalingGroupEvents interface {
	// EventBridge event pattern for AutoScalingGroup EC2 Instance-launch Lifecycle Action.
	// Experimental.
	EC2InstanceLaunchLifecycleActionPattern(options *AutoScalingGroupEvents_EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance Launch Successful.
	// Experimental.
	EC2InstanceLaunchSuccessfulPattern(options *AutoScalingGroupEvents_EC2InstanceLaunchSuccessful_EC2InstanceLaunchSuccessfulProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance Launch Unsuccessful.
	// Experimental.
	EC2InstanceLaunchUnsuccessfulPattern(options *AutoScalingGroupEvents_EC2InstanceLaunchUnsuccessful_EC2InstanceLaunchUnsuccessfulProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance-terminate Lifecycle Action.
	// Experimental.
	EC2InstanceTerminateLifecycleActionPattern(options *AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction_EC2InstanceTerminateLifecycleActionProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance Terminate Successful.
	// Experimental.
	EC2InstanceTerminateSuccessfulPattern(options *AutoScalingGroupEvents_EC2InstanceTerminateSuccessful_EC2InstanceTerminateSuccessfulProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance Terminate Unsuccessful.
	// Experimental.
	EC2InstanceTerminateUnsuccessfulPattern(options *AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful_EC2InstanceTerminateUnsuccessfulProps) *awsevents.EventPattern
}

// The jsii proxy struct for AutoScalingGroupEvents
type jsiiProxy_AutoScalingGroupEvents struct {
	_ byte // padding
}

// Create AutoScalingGroupEvents from a AutoScalingGroup reference.
// Experimental.
func AutoScalingGroupEvents_FromAutoScalingGroup(autoScalingGroupRef interfacesawsautoscaling.IAutoScalingGroupRef) AutoScalingGroupEvents {
	_init_.Initialize()

	if err := validateAutoScalingGroupEvents_FromAutoScalingGroupParameters(autoScalingGroupRef); err != nil {
		panic(err)
	}
	var returns AutoScalingGroupEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents",
		"fromAutoScalingGroup",
		[]interface{}{autoScalingGroupRef},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) EC2InstanceLaunchLifecycleActionPattern(options *AutoScalingGroupEvents_EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps) *awsevents.EventPattern {
	if err := a.validateEC2InstanceLaunchLifecycleActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"eC2InstanceLaunchLifecycleActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) EC2InstanceLaunchSuccessfulPattern(options *AutoScalingGroupEvents_EC2InstanceLaunchSuccessful_EC2InstanceLaunchSuccessfulProps) *awsevents.EventPattern {
	if err := a.validateEC2InstanceLaunchSuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"eC2InstanceLaunchSuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) EC2InstanceLaunchUnsuccessfulPattern(options *AutoScalingGroupEvents_EC2InstanceLaunchUnsuccessful_EC2InstanceLaunchUnsuccessfulProps) *awsevents.EventPattern {
	if err := a.validateEC2InstanceLaunchUnsuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"eC2InstanceLaunchUnsuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) EC2InstanceTerminateLifecycleActionPattern(options *AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction_EC2InstanceTerminateLifecycleActionProps) *awsevents.EventPattern {
	if err := a.validateEC2InstanceTerminateLifecycleActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"eC2InstanceTerminateLifecycleActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) EC2InstanceTerminateSuccessfulPattern(options *AutoScalingGroupEvents_EC2InstanceTerminateSuccessful_EC2InstanceTerminateSuccessfulProps) *awsevents.EventPattern {
	if err := a.validateEC2InstanceTerminateSuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"eC2InstanceTerminateSuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) EC2InstanceTerminateUnsuccessfulPattern(options *AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful_EC2InstanceTerminateUnsuccessfulProps) *awsevents.EventPattern {
	if err := a.validateEC2InstanceTerminateUnsuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"eC2InstanceTerminateUnsuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

