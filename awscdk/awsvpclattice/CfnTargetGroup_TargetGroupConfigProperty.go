package awsvpclattice


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
	// `CfnTargetGroup.TargetGroupConfigProperty.Port`.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// `CfnTargetGroup.TargetGroupConfigProperty.Protocol`.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// `CfnTargetGroup.TargetGroupConfigProperty.VpcIdentifier`.
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// `CfnTargetGroup.TargetGroupConfigProperty.HealthCheck`.
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// `CfnTargetGroup.TargetGroupConfigProperty.ProtocolVersion`.
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
}

