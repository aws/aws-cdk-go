package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Health check settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheck := awscdk.Aws_autoscaling.HealthCheck_Ec2(&Ec2HealthCheckOptions{
//   	Grace: cdk.Duration_Minutes(jsii.Number(30)),
//   })
//
type HealthCheck interface {
	GracePeriod() awscdk.Duration
	Type() *string
}

// The jsii proxy struct for HealthCheck
type jsiiProxy_HealthCheck struct {
	_ byte // padding
}

func (j *jsiiProxy_HealthCheck) GracePeriod() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Use EC2 for health checks.
func HealthCheck_Ec2(options *Ec2HealthCheckOptions) HealthCheck {
	_init_.Initialize()

	if err := validateHealthCheck_Ec2Parameters(options); err != nil {
		panic(err)
	}
	var returns HealthCheck

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.HealthCheck",
		"ec2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Use ELB for health checks.
//
// It considers the instance unhealthy if it fails either the EC2 status checks or the load balancer health checks.
func HealthCheck_Elb(options *ElbHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	if err := validateHealthCheck_ElbParameters(options); err != nil {
		panic(err)
	}
	var returns HealthCheck

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.HealthCheck",
		"elb",
		[]interface{}{options},
		&returns,
	)

	return returns
}

