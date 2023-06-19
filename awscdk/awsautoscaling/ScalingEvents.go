package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of ScalingEvents, you can use one of the predefined lists, such as ScalingEvents.ERRORS or create a custom group by instantiating a `NotificationTypes` object, e.g: `new NotificationTypes(`NotificationType.INSTANCE_LAUNCH`)`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingEvents := awscdk.Aws_autoscaling.ScalingEvents_ALL()
//
// Experimental.
type ScalingEvents interface {
}

// The jsii proxy struct for ScalingEvents
type jsiiProxy_ScalingEvents struct {
	_ byte // padding
}

// Experimental.
func NewScalingEvents(types ...ScalingEvent) ScalingEvents {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range types {
		args = append(args, a)
	}

	j := jsiiProxy_ScalingEvents{}

	_jsii_.Create(
		"monocdk.aws_autoscaling.ScalingEvents",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewScalingEvents_Override(s ScalingEvents, types ...ScalingEvent) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range types {
		args = append(args, a)
	}

	_jsii_.Create(
		"monocdk.aws_autoscaling.ScalingEvents",
		args,
		s,
	)
}

func ScalingEvents_ALL() ScalingEvents {
	_init_.Initialize()
	var returns ScalingEvents
	_jsii_.StaticGet(
		"monocdk.aws_autoscaling.ScalingEvents",
		"ALL",
		&returns,
	)
	return returns
}

func ScalingEvents_ERRORS() ScalingEvents {
	_init_.Initialize()
	var returns ScalingEvents
	_jsii_.StaticGet(
		"monocdk.aws_autoscaling.ScalingEvents",
		"ERRORS",
		&returns,
	)
	return returns
}

func ScalingEvents_LAUNCH_EVENTS() ScalingEvents {
	_init_.Initialize()
	var returns ScalingEvents
	_jsii_.StaticGet(
		"monocdk.aws_autoscaling.ScalingEvents",
		"LAUNCH_EVENTS",
		&returns,
	)
	return returns
}

func ScalingEvents_TERMINATION_EVENTS() ScalingEvents {
	_init_.Initialize()
	var returns ScalingEvents
	_jsii_.StaticGet(
		"monocdk.aws_autoscaling.ScalingEvents",
		"TERMINATION_EVENTS",
		&returns,
	)
	return returns
}

