package previewawsvpclatticemixins


// Describes the configuration of a target group.
//
// For more information, see [Target groups](https://docs.aws.amazon.com/vpc-lattice/latest/ug/target-groups.html) in the *Amazon VPC Lattice User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetGroupConfigProperty := &TargetGroupConfigProperty{
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
//   	LambdaEventStructureVersion: jsii.String("lambdaEventStructureVersion"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	ProtocolVersion: jsii.String("protocolVersion"),
//   	VpcIdentifier: jsii.String("vpcIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html
//
type CfnTargetGroupPropsMixin_TargetGroupConfigProperty struct {
	// The health check configuration.
	//
	// Not supported if the target group type is `LAMBDA` or `ALB` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The type of IP address used for the target group.
	//
	// Supported only if the target group type is `IP` . The default is `IPV4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-ipaddresstype
	//
	// Default: - "IPV4".
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The version of the event structure that your Lambda function receives.
	//
	// Supported only if the target group type is `LAMBDA` . The default is `V1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-lambdaeventstructureversion
	//
	LambdaEventStructureVersion *string `field:"optional" json:"lambdaEventStructureVersion" yaml:"lambdaEventStructureVersion"`
	// The port on which the targets are listening.
	//
	// For HTTP, the default is 80. For HTTPS, the default is 443. Not supported if the target group type is `LAMBDA` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use for routing traffic to the targets.
	//
	// The default is the protocol of the target group. Not supported if the target group type is `LAMBDA` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol version.
	//
	// The default is `HTTP1` . Not supported if the target group type is `LAMBDA` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-protocolversion
	//
	// Default: - "HTTP1".
	//
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// The ID of the VPC.
	//
	// Not supported if the target group type is `LAMBDA` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-targetgroupconfig.html#cfn-vpclattice-targetgroup-targetgroupconfig-vpcidentifier
	//
	VpcIdentifier *string `field:"optional" json:"vpcIdentifier" yaml:"vpcIdentifier"`
}

