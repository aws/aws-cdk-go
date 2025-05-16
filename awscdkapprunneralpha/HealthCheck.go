package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner"
)

// Contains static factory methods for creating health checks for different protocols.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	HealthCheck: apprunner.HealthCheck_Http(&HttpHealthCheckOptions{
//   		HealthyThreshold: jsii.Number(5),
//   		Interval: awscdk.Duration_Seconds(jsii.Number(10)),
//   		Path: jsii.String("/"),
//   		Timeout: awscdk.Duration_*Seconds(jsii.Number(10)),
//   		UnhealthyThreshold: jsii.Number(10),
//   	}),
//   })
//
// Experimental.
type HealthCheck interface {
	// Experimental.
	HealthCheckProtocolType() HealthCheckProtocolType
	// Experimental.
	HealthyThreshold() *float64
	// Experimental.
	Interval() awscdk.Duration
	// Experimental.
	Path() *string
	// Experimental.
	Timeout() awscdk.Duration
	// Experimental.
	UnhealthyThreshold() *float64
	// Experimental.
	Bind() *awsapprunner.CfnService_HealthCheckConfigurationProperty
}

// The jsii proxy struct for HealthCheck
type jsiiProxy_HealthCheck struct {
	_ byte // padding
}

func (j *jsiiProxy_HealthCheck) HealthCheckProtocolType() HealthCheckProtocolType {
	var returns HealthCheckProtocolType
	_jsii_.Get(
		j,
		"healthCheckProtocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Interval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}


// Construct a HTTP health check.
// Experimental.
func HealthCheck_Http(options *HttpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	if err := validateHealthCheck_HttpParameters(options); err != nil {
		panic(err)
	}
	var returns HealthCheck

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.HealthCheck",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a TCP health check.
// Experimental.
func HealthCheck_Tcp(options *TcpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	if err := validateHealthCheck_TcpParameters(options); err != nil {
		panic(err)
	}
	var returns HealthCheck

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.HealthCheck",
		"tcp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthCheck) Bind() *awsapprunner.CfnService_HealthCheckConfigurationProperty {
	var returns *awsapprunner.CfnService_HealthCheckConfigurationProperty

	_jsii_.Invoke(
		h,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

