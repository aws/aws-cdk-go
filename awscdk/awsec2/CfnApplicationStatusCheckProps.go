package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplicationStatusCheck`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationStatusCheckProps := &CfnApplicationStatusCheckProps{
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	Aggregation: jsii.String("aggregation"),
//   	DeviceIndex: jsii.Number(123),
//   	FailureThreshold: jsii.Number(123),
//   	HealthCheckPaths: []interface{}{
//   		&HealthCheckPathProperty{
//   			Destinations: []interface{}{
//   				&HealthCheckPathDestinationProperty{
//   					SecurityGroupId: jsii.String("securityGroupId"),
//   					SubnetId: jsii.String("subnetId"),
//   				},
//   			},
//   			Source: &HealthCheckPathSourceProperty{
//   				SecurityGroupId: jsii.String("securityGroupId"),
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   	},
//   	InitializationGracePeriodSeconds: jsii.Number(123),
//   	Interval: jsii.Number(123),
//   	IpScope: jsii.String("ipScope"),
//   	IpVersion: jsii.String("ipVersion"),
//   	Path: jsii.String("path"),
//   	StatusCodeMatcher: jsii.String("statusCodeMatcher"),
//   	SuccessThreshold: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html
//
type CfnApplicationStatusCheckProps struct {
	// The port used for the health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network protocol used for the health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Whether this check is included in the rolled-up application status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-aggregation
	//
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// The network interface device index used for the health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-deviceindex
	//
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// The number of consecutive failed probes required to mark the instance unhealthy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-failurethreshold
	//
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// The source/destination network paths used for the health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-healthcheckpaths
	//
	HealthCheckPaths interface{} `field:"optional" json:"healthCheckPaths" yaml:"healthCheckPaths"`
	// Seconds to wait after instance launch before beginning health checks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-initializationgraceperiodseconds
	//
	InitializationGracePeriodSeconds *float64 `field:"optional" json:"initializationGracePeriodSeconds" yaml:"initializationGracePeriodSeconds"`
	// The interval, in seconds, between health check probes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The IP scope used for the health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-ipscope
	//
	IpScope *string `field:"optional" json:"ipScope" yaml:"ipScope"`
	// The IP version used for the health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-ipversion
	//
	IpVersion *string `field:"optional" json:"ipVersion" yaml:"ipVersion"`
	// The HTTP path used for the health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP status codes considered successful (e.g., "200-299").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-statuscodematcher
	//
	StatusCodeMatcher *string `field:"optional" json:"statusCodeMatcher" yaml:"statusCodeMatcher"`
	// The number of consecutive successful probes required to mark the instance healthy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-successthreshold
	//
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// Tags to apply to the application status check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The timeout, in seconds, for each health check probe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html#cfn-ec2-applicationstatuscheck-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

