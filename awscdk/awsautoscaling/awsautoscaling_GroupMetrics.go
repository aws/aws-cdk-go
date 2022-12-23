package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A set of group metrics.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   // Enable monitoring of all group metrics
//   // Enable monitoring of all group metrics
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: machineImage,
//
//   	// ...
//
//   	groupMetrics: []groupMetrics{
//   		autoscaling.*groupMetrics.all(),
//   	},
//   })
//
//   // Enable monitoring for a subset of group metrics
//   // Enable monitoring for a subset of group metrics
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: machineImage,
//
//   	// ...
//
//   	groupMetrics: []*groupMetrics{
//   		autoscaling.NewGroupMetrics(autoscaling.groupMetric_MIN_SIZE(), autoscaling.*groupMetric_MAX_SIZE()),
//   	},
//   })
//
// Experimental.
type GroupMetrics interface {
}

// The jsii proxy struct for GroupMetrics
type jsiiProxy_GroupMetrics struct {
	_ byte // padding
}

// Experimental.
func NewGroupMetrics(metrics ...GroupMetric) GroupMetrics {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range metrics {
		args = append(args, a)
	}

	j := jsiiProxy_GroupMetrics{}

	_jsii_.Create(
		"monocdk.aws_autoscaling.GroupMetrics",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewGroupMetrics_Override(g GroupMetrics, metrics ...GroupMetric) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range metrics {
		args = append(args, a)
	}

	_jsii_.Create(
		"monocdk.aws_autoscaling.GroupMetrics",
		args,
		g,
	)
}

// Report all group metrics.
// Experimental.
func GroupMetrics_All() GroupMetrics {
	_init_.Initialize()

	var returns GroupMetrics

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.GroupMetrics",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

