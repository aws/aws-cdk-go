package awsapprunner


// Describes the settings for the health check that AWS App Runner performs to monitor the health of a service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigurationProperty := &HealthCheckConfigurationProperty{
//   	HealthyThreshold: jsii.Number(123),
//   	Interval: jsii.Number(123),
//   	Path: jsii.String("path"),
//   	Protocol: jsii.String("protocol"),
//   	Timeout: jsii.Number(123),
//   	UnhealthyThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-healthcheckconfiguration.html
//
type CfnService_HealthCheckConfigurationProperty struct {
	// The number of consecutive checks that must succeed before App Runner decides that the service is healthy.
	//
	// Default: `1`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-healthcheckconfiguration.html#cfn-apprunner-service-healthcheckconfiguration-healthythreshold
	//
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time interval, in seconds, between health checks.
	//
	// Default: `5`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-healthcheckconfiguration.html#cfn-apprunner-service-healthcheckconfiguration-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The URL that health check requests are sent to.
	//
	// `Path` is only applicable when you set `Protocol` to `HTTP` .
	//
	// Default: `"/"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-healthcheckconfiguration.html#cfn-apprunner-service-healthcheckconfiguration-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The IP protocol that App Runner uses to perform health checks for your service.
	//
	// If you set `Protocol` to `HTTP` , App Runner sends health check requests to the HTTP path specified by `Path` .
	//
	// Default: `TCP`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-healthcheckconfiguration.html#cfn-apprunner-service-healthcheckconfiguration-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The time, in seconds, to wait for a health check response before deciding it failed.
	//
	// Default: `2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-healthcheckconfiguration.html#cfn-apprunner-service-healthcheckconfiguration-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive checks that must fail before App Runner decides that the service is unhealthy.
	//
	// Default: `5`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-healthcheckconfiguration.html#cfn-apprunner-service-healthcheckconfiguration-unhealthythreshold
	//
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

