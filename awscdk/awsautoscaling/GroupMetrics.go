package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// ...
//
//   	GroupMetrics: []groupMetrics{
//   		autoscaling.*groupMetrics_All(),
//   	},
//   })
//
//   // Enable monitoring for a subset of group metrics
//   // Enable monitoring for a subset of group metrics
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// ...
//
//   	GroupMetrics: []*groupMetrics{
//   		autoscaling.NewGroupMetrics(autoscaling.GroupMetric_MIN_SIZE(), autoscaling.GroupMetric_MAX_SIZE()),
//   	},
//   })
//
type GroupMetrics interface {
}

// The jsii proxy struct for GroupMetrics
type jsiiProxy_GroupMetrics struct {
	_ byte // padding
}

func NewGroupMetrics(metrics ...GroupMetric) GroupMetrics {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range metrics {
		args = append(args, a)
	}

	j := jsiiProxy_GroupMetrics{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.GroupMetrics",
		args,
		&j,
	)

	return &j
}

func NewGroupMetrics_Override(g GroupMetrics, metrics ...GroupMetric) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range metrics {
		args = append(args, a)
	}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.GroupMetrics",
		args,
		g,
	)
}

// Report all group metrics.
func GroupMetrics_All() GroupMetrics {
	_init_.Initialize()

	var returns GroupMetrics

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.GroupMetrics",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

