package awsvpclattice


// The health check configuration of a target group.
//
// Health check configurations aren't used for `LAMBDA` and `ALB` target groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigProperty := &HealthCheckConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   	HealthCheckIntervalSeconds: jsii.Number(123),
//   	HealthCheckTimeoutSeconds: jsii.Number(123),
//   	HealthyThresholdCount: jsii.Number(123),
//   	Matcher: &MatcherProperty{
//   		HttpCode: jsii.String("httpCode"),
//   	},
//   	Path: jsii.String("path"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	ProtocolVersion: jsii.String("protocolVersion"),
//   	UnhealthyThresholdCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html
//
type CfnTargetGroup_HealthCheckConfigProperty struct {
	// Indicates whether health checking is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The approximate amount of time, in seconds, between health checks of an individual target.
	//
	// The range is 5–300 seconds. The default is 30 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-healthcheckintervalseconds
	//
	HealthCheckIntervalSeconds *float64 `field:"optional" json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// The amount of time, in seconds, to wait before reporting a target as unhealthy.
	//
	// The range is 1–120 seconds. The default is 5 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-healthchecktimeoutseconds
	//
	HealthCheckTimeoutSeconds *float64 `field:"optional" json:"healthCheckTimeoutSeconds" yaml:"healthCheckTimeoutSeconds"`
	// The number of consecutive successful health checks required before considering an unhealthy target healthy.
	//
	// The range is 2–10. The default is 5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-healthythresholdcount
	//
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// The codes to use when checking for a successful response from a target.
	//
	// These are called *Success codes* in the console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-matcher
	//
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// The destination for health checks on the targets.
	//
	// If the protocol version is `HTTP/1.1` or `HTTP/2` , specify a valid URI (for example, `/path?query` ). The default path is `/` . Health checks are not supported if the protocol version is `gRPC` , however, you can choose `HTTP/1.1` or `HTTP/2` and specify a valid URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port used when performing health checks on targets.
	//
	// The default setting is the port that a target receives traffic on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol used when performing health checks on targets.
	//
	// The possible protocols are `HTTP` and `HTTPS` . The default is `HTTP` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol version used when performing health checks on targets.
	//
	// The possible protocol versions are `HTTP1` and `HTTP2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-protocolversion
	//
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// The number of consecutive failed health checks required before considering a target unhealthy.
	//
	// The range is 2–10. The default is 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-healthcheckconfig.html#cfn-vpclattice-targetgroup-healthcheckconfig-unhealthythresholdcount
	//
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
}

