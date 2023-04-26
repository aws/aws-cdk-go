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
//   		UnhealthyThresholdCount: jsii.Number(123),
//   	},
//   	ProtocolVersion: jsii.String("protocolVersion"),
//   }
//
type CfnTargetGroup_TargetGroupConfigProperty struct {
	// The port on which the targets are listening.
	//
	// For HTTP, the default is `80` . For HTTPS, the default is `443`
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The protocol to use for routing traffic to the targets.
	//
	// Default is the protocol of a target group.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The ID of the VPC.
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// The health check configuration.
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The protocol version.
	//
	// Default value is `HTTP1` .
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
}

