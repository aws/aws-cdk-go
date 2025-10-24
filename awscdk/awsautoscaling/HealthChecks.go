package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Health check settings for multiple types.
//
// Example:
//   var vpc Vpc
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
//   	HealthChecks: autoscaling.HealthChecks_Ec2(&Ec2HealthChecksOptions{
//   		GracePeriod: awscdk.Duration_Seconds(jsii.Number(100)),
//   	}),
//   })
//
type HealthChecks interface {
	GracePeriod() awscdk.Duration
	Types() *[]*string
}

// The jsii proxy struct for HealthChecks
type jsiiProxy_HealthChecks struct {
	_ byte // padding
}

func (j *jsiiProxy_HealthChecks) GracePeriod() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthChecks) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}


// Use EC2 only for health checks.
func HealthChecks_Ec2(options *Ec2HealthChecksOptions) HealthChecks {
	_init_.Initialize()

	if err := validateHealthChecks_Ec2Parameters(options); err != nil {
		panic(err)
	}
	var returns HealthChecks

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.HealthChecks",
		"ec2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Use additional health checks other than EC2.
//
// Specify types other than EC2, as EC2 is always enabled.
// It considers the instance unhealthy if it fails either the EC2 status checks or the additional health checks.
func HealthChecks_WithAdditionalChecks(options *AdditionalHealthChecksOptions) HealthChecks {
	_init_.Initialize()

	if err := validateHealthChecks_WithAdditionalChecksParameters(options); err != nil {
		panic(err)
	}
	var returns HealthChecks

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.HealthChecks",
		"withAdditionalChecks",
		[]interface{}{options},
		&returns,
	)

	return returns
}

