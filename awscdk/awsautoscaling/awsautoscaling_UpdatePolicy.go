package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How existing instances should be updated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updatePolicy := awscdk.Aws_autoscaling.updatePolicy.replacingUpdate()
//
// Experimental.
type UpdatePolicy interface {
}

// The jsii proxy struct for UpdatePolicy
type jsiiProxy_UpdatePolicy struct {
	_ byte // padding
}

// Experimental.
func NewUpdatePolicy_Override(u UpdatePolicy) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling.UpdatePolicy",
		nil, // no parameters
		u,
	)
}

// Create a new AutoScalingGroup and switch over to it.
// Experimental.
func UpdatePolicy_ReplacingUpdate() UpdatePolicy {
	_init_.Initialize()

	var returns UpdatePolicy

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.UpdatePolicy",
		"replacingUpdate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Replace the instances in the AutoScalingGroup one by one, or in batches.
// Experimental.
func UpdatePolicy_RollingUpdate(options *RollingUpdateOptions) UpdatePolicy {
	_init_.Initialize()

	if err := validateUpdatePolicy_RollingUpdateParameters(options); err != nil {
		panic(err)
	}
	var returns UpdatePolicy

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.UpdatePolicy",
		"rollingUpdate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

