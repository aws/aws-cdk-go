package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How existing instances should be updated.
//
// Example:
//   var vpc Vpc
//   var instanceType InstanceType
//   var machineImage IMachineImage
//   var alarm Alarm
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	UpdatePolicy: autoscaling.UpdatePolicy_InstanceRefresh(&InstanceRefreshOptions{
//   		Strategy: autoscaling.InstanceRefreshStrategy_ROLLING,
//   		MinHealthyPercentage: jsii.Number(90),
//   		MaxHealthyPercentage: jsii.Number(100),
//   		InstanceWarmup: awscdk.Duration_Seconds(jsii.Number(300)),
//   		CheckpointPercentages: []*f64{
//   			jsii.Number(50),
//   			jsii.Number(100),
//   		},
//   		CheckpointDelay: awscdk.Duration_Minutes(jsii.Number(10)),
//   		Alarms: []IAlarmRef{
//   			alarm,
//   		},
//   	}),
//   })
//
type UpdatePolicy interface {
}

// The jsii proxy struct for UpdatePolicy
type jsiiProxy_UpdatePolicy struct {
	_ byte // padding
}

func NewUpdatePolicy_Override(u UpdatePolicy) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.UpdatePolicy",
		nil, // no parameters
		u,
	)
}

// Use instance refresh to update the instances in the AutoScalingGroup.
//
// When properties that trigger an instance refresh change (such as LaunchTemplate
// or MixedInstancesPolicy), CloudFormation starts an instance refresh to replace
// instances gradually while maintaining availability.
func UpdatePolicy_InstanceRefresh(options *InstanceRefreshOptions) UpdatePolicy {
	_init_.Initialize()

	if err := validateUpdatePolicy_InstanceRefreshParameters(options); err != nil {
		panic(err)
	}
	var returns UpdatePolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.UpdatePolicy",
		"instanceRefresh",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Create a new AutoScalingGroup and switch over to it.
func UpdatePolicy_ReplacingUpdate() UpdatePolicy {
	_init_.Initialize()

	var returns UpdatePolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.UpdatePolicy",
		"replacingUpdate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Replace the instances in the AutoScalingGroup one by one, or in batches.
func UpdatePolicy_RollingUpdate(options *RollingUpdateOptions) UpdatePolicy {
	_init_.Initialize()

	if err := validateUpdatePolicy_RollingUpdateParameters(options); err != nil {
		panic(err)
	}
	var returns UpdatePolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.UpdatePolicy",
		"rollingUpdate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

