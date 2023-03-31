package awsapprunner


// Describes the settings for the health check that AWS App Runner performs to monitor the health of a service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigurationProperty := &healthCheckConfigurationProperty{
//   	healthyThreshold: jsii.Number(123),
//   	interval: jsii.Number(123),
//   	path: jsii.String("path"),
//   	protocol: jsii.String("protocol"),
//   	timeout: jsii.Number(123),
//   	unhealthyThreshold: jsii.Number(123),
//   }
//
type CfnService_HealthCheckConfigurationProperty struct {
	// The number of consecutive checks that must succeed before App Runner decides that the service is healthy.
	//
	// Default: `1`.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time interval, in seconds, between health checks.
	//
	// Default: `5`.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The URL that health check requests are sent to.
	//
	// `Path` is only applicable when you set `Protocol` to `HTTP` .
	//
	// Default: `"/"`.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The IP protocol that App Runner uses to perform health checks for your service.
	//
	// If you set `Protocol` to `HTTP` , App Runner sends health check requests to the HTTP path specified by `Path` .
	//
	// Default: `TCP`.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The time, in seconds, to wait for a health check response before deciding it failed.
	//
	// Default: `2`.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive checks that must fail before App Runner decides that the service is unhealthy.
	//
	// Default: `5`.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

