package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties used to define HTTP Based healthchecks.
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
type HttpHealthCheckOptions struct {
	// The number of consecutive checks that must succeed before App Runner decides that the service is healthy.
	// Default: 1.
	//
	// Experimental.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time interval, in seconds, between health checks.
	// Default: Duration.seconds(5)
	//
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The URL that health check requests are sent to.
	// Default: /.
	//
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The time, in seconds, to wait for a health check response before deciding it failed.
	// Default: Duration.seconds(2)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive checks that must fail before App Runner decides that the service is unhealthy.
	// Default: 5.
	//
	// Experimental.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

