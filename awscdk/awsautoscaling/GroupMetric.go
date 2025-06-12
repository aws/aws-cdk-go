package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Group metrics that an Auto Scaling group sends to Amazon CloudWatch.
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
type GroupMetric interface {
	// The name of the group metric.
	Name() *string
}

// The jsii proxy struct for GroupMetric
type jsiiProxy_GroupMetric struct {
	_ byte // padding
}

func (j *jsiiProxy_GroupMetric) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewGroupMetric(name *string) GroupMetric {
	_init_.Initialize()

	if err := validateNewGroupMetricParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupMetric{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		[]interface{}{name},
		&j,
	)

	return &j
}

func NewGroupMetric_Override(g GroupMetric, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		[]interface{}{name},
		g,
	)
}

func GroupMetric_DESIRED_CAPACITY() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"DESIRED_CAPACITY",
		&returns,
	)
	return returns
}

func GroupMetric_IN_SERVICE_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"IN_SERVICE_INSTANCES",
		&returns,
	)
	return returns
}

func GroupMetric_MAX_SIZE() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"MAX_SIZE",
		&returns,
	)
	return returns
}

func GroupMetric_MIN_SIZE() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"MIN_SIZE",
		&returns,
	)
	return returns
}

func GroupMetric_PENDING_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"PENDING_INSTANCES",
		&returns,
	)
	return returns
}

func GroupMetric_STANDBY_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"STANDBY_INSTANCES",
		&returns,
	)
	return returns
}

func GroupMetric_TERMINATING_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"TERMINATING_INSTANCES",
		&returns,
	)
	return returns
}

func GroupMetric_TOTAL_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"TOTAL_INSTANCES",
		&returns,
	)
	return returns
}

