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
	Ec2InstanceLaunchLifecycleActionPattern(options *EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance Launch Successful.
	// Experimental.
	Ec2InstanceLaunchSuccessfulPattern(options *EC2InstanceLaunchSuccessful_EC2InstanceLaunchSuccessfulProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance Launch Unsuccessful.
	// Experimental.
	Ec2InstanceLaunchUnsuccessfulPattern(options *EC2InstanceLaunchUnsuccessful_EC2InstanceLaunchUnsuccessfulProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance-terminate Lifecycle Action.
	// Experimental.
	Ec2InstanceTerminateLifecycleActionPattern(options *EC2InstanceTerminateLifecycleAction_EC2InstanceTerminateLifecycleActionProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance Terminate Successful.
	// Experimental.
	Ec2InstanceTerminateSuccessfulPattern(options *EC2InstanceTerminateSuccessful_EC2InstanceTerminateSuccessfulProps) *awsevents.EventPattern
	// EventBridge event pattern for AutoScalingGroup EC2 Instance Terminate Unsuccessful.
	// Experimental.
	Ec2InstanceTerminateUnsuccessfulPattern(options *EC2InstanceTerminateUnsuccessful_EC2InstanceTerminateUnsuccessfulProps) *awsevents.EventPattern
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

func (a *jsiiProxy_AutoScalingGroupEvents) Ec2InstanceLaunchLifecycleActionPattern(options *EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps) *awsevents.EventPattern {
	if err := a.validateEc2InstanceLaunchLifecycleActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"ec2InstanceLaunchLifecycleActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) Ec2InstanceLaunchSuccessfulPattern(options *EC2InstanceLaunchSuccessful_EC2InstanceLaunchSuccessfulProps) *awsevents.EventPattern {
	if err := a.validateEc2InstanceLaunchSuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"ec2InstanceLaunchSuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) Ec2InstanceLaunchUnsuccessfulPattern(options *EC2InstanceLaunchUnsuccessful_EC2InstanceLaunchUnsuccessfulProps) *awsevents.EventPattern {
	if err := a.validateEc2InstanceLaunchUnsuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"ec2InstanceLaunchUnsuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) Ec2InstanceTerminateLifecycleActionPattern(options *EC2InstanceTerminateLifecycleAction_EC2InstanceTerminateLifecycleActionProps) *awsevents.EventPattern {
	if err := a.validateEc2InstanceTerminateLifecycleActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"ec2InstanceTerminateLifecycleActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) Ec2InstanceTerminateSuccessfulPattern(options *EC2InstanceTerminateSuccessful_EC2InstanceTerminateSuccessfulProps) *awsevents.EventPattern {
	if err := a.validateEc2InstanceTerminateSuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"ec2InstanceTerminateSuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupEvents) Ec2InstanceTerminateUnsuccessfulPattern(options *EC2InstanceTerminateUnsuccessful_EC2InstanceTerminateUnsuccessfulProps) *awsevents.EventPattern {
	if err := a.validateEc2InstanceTerminateUnsuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"ec2InstanceTerminateUnsuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

