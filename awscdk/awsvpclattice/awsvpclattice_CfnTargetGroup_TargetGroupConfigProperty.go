package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupConfigProperty := &targetGroupConfigProperty{
//   	port: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   	vpcIdentifier: jsii.String("vpcIdentifier"),
//
//   	// the properties below are optional
//   	healthCheck: &healthCheckConfigProperty{
//   		enabled: jsii.Boolean(false),
//   		healthCheckIntervalSeconds: jsii.Number(123),
//   		healthCheckTimeoutSeconds: jsii.Number(123),
//   		healthyThresholdCount: jsii.Number(123),
//   		matcher: &matcherProperty{
//   			httpCode: jsii.String("httpCode"),
//   		},
//   		path: jsii.String("path"),
//   		port: jsii.Number(123),
//   		protocol: jsii.String("protocol"),
//   		unhealthyThresholdCount: jsii.Number(123),
//   	},
//   	protocolVersion: jsii.String("protocolVersion"),
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

