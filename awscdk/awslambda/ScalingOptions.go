package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration options for scaling a capacity provider, including scaling mode and policies.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"))
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//
//   capacityProvider := lambda.NewCapacityProvider(this, jsii.String("MyCapacityProvider"), &CapacityProviderProps{
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	ScalingOptions: lambda.ScalingOptions_Manual([]TargetTrackingScalingPolicy{
//   		lambda.TargetTrackingScalingPolicy_CpuUtilization(jsii.Number(70)),
//   	}),
//   })
//
type ScalingOptions interface {
	// The scaling mode for the capacity provider.
	ScalingMode() *string
	// The target tracking scaling policies used when scaling mode is 'Manual'.
	ScalingPolicies() *[]TargetTrackingScalingPolicy
}

// The jsii proxy struct for ScalingOptions
type jsiiProxy_ScalingOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_ScalingOptions) ScalingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalingOptions) ScalingPolicies() *[]TargetTrackingScalingPolicy {
	var returns *[]TargetTrackingScalingPolicy
	_jsii_.Get(
		j,
		"scalingPolicies",
		&returns,
	)
	return returns
}


// Creates scaling options where the capacity provider manages scaling automatically.
func ScalingOptions_Auto() ScalingOptions {
	_init_.Initialize()

	var returns ScalingOptions

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.ScalingOptions",
		"auto",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creates manual scaling options with custom target tracking scaling policies.
//
// At least one policy is required.
func ScalingOptions_Manual(scalingPolicies *[]TargetTrackingScalingPolicy) ScalingOptions {
	_init_.Initialize()

	if err := validateScalingOptions_ManualParameters(scalingPolicies); err != nil {
		panic(err)
	}
	var returns ScalingOptions

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.ScalingOptions",
		"manual",
		[]interface{}{scalingPolicies},
		&returns,
	)

	return returns
}

