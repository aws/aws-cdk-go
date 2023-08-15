package awsvpclattice


// Describes the configuration of a target group.
//
// Lambda functions don't support target group configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupConfigProperty := &TargetGroupConfigProperty{
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   	// the properties below are optional
//   	HealthCheck: &HealthCheckConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   		HealthCheckIntervalSeconds: jsii.Number(123),
//   		HealthCheckTimeoutSeconds: jsii.Number(123),
//   		HealthyThresholdCount: jsii.Number(123),
//   		Matcher: &MatcherProperty{
//   			HttpCode: jsii.String("httpCode"),
//   		},
//   		Path: jsii.String("path"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		ProtocolVersion: jsii.String("protocolVersion"),
//   		UnhealthyThresholdCount: jsii.Number(123),
//   	},
//   	IpAddressType: jsii.String("ipAddressType"),
//   	ProtocolVersion: jsii.String("protocolVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html
//
type CfnTargetGroup_TargetGroupConfigProperty struct {
	// The port on which the targets are listening.
	//
	// For HTTP, the default is 80. For HTTPS, the default is 443.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The protocol to use for routing traffic to the targets.
	//
	// Default is the protocol of a target group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-vpcidentifier
	//
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// The health check configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The type of IP address used for the target group.
	//
	// The possible values are `ipv4` and `ipv6` . This is an optional parameter. If not specified, the default is `ipv4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-ipaddresstype
	//
	// Default: - "IPV4".
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The protocol version.
	//
	// Default value is `HTTP1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-protocolversion
	//
	// Default: - "HTTP1".
	//
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
}

